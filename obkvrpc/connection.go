/*-
 * #%L
 * OBKV Table Client Framework
 * %%
 * Copyright (C) 2021 OceanBase
 * %%
 * OBKV Table Client Framework is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * #L%
 */

package obkvrpc

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"

	oberror "github.com/oceanbase/obkv-table-client-go/error"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type ConnectionOption struct {
	ip             string
	port           int
	connectTimeout time.Duration

	tenantName   string
	databaseName string
	userName     string
	password     string
}

func NewConnectionOption(ip string, port int, connectTimeout time.Duration, tenantName string, databaseName string, userName string, password string) *ConnectionOption {
	return &ConnectionOption{ip: ip, port: port, connectTimeout: connectTimeout, tenantName: tenantName, databaseName: databaseName, userName: userName, password: password}
}

type Connection struct {
	option *ConnectionOption

	conn    net.Conn
	mutex   sync.Mutex
	seq     atomic.Uint32 // as channel id in ez header
	pending map[uint32]*Call
	active  atomic.Bool

	uniqueId uint64        // as trace0 in rpc header
	sequence atomic.Uint64 // as trace1 in rpc header

	credential []byte
	tenantId   uint64
}

// Call represents an active RPC.
type Call struct {
	Error   error
	Done    chan *Call // Strobes when call is complete.
	Content []byte
}

func NewConnection(option *ConnectionOption) *Connection {
	return &Connection{option: option, pending: make(map[uint32]*Call)}
}

func (c *Connection) Connect() error {
	address := fmt.Sprintf("%s:%s", c.option.ip, strconv.Itoa(c.option.port))
	conn, err := net.DialTimeout("tcp", address, c.option.connectTimeout)
	if err != nil {
		return errors.WithMessagef(err, "net dial, uniqueId: %d remote addr: %s", c.uniqueId, address)
	}
	c.conn = conn

	/* layout of uniqueId(64 bytes)
	 * ip_: 32
	 * port_: 16;
	 * is_user_request_: 1;
	 * is_ipv6_:1;
	 * reserved_: 14;
	 */
	ip := int64(util.ConvertIpToUint32(c.conn.LocalAddr().(*net.TCPAddr).IP))
	port := int64(c.conn.LocalAddr().(*net.TCPAddr).Port << 32)
	var isUserRequest int64 = 1 << (32 + 16)
	var reserved int64 = 0
	c.uniqueId = uint64(ip | port | isUserRequest | reserved)

	go c.receivePacket()
	return nil
}

func (c *Connection) Login() error {
	loginRequest := protocol.NewObLoginRequest(c.option.tenantName, c.option.databaseName, c.option.userName, c.option.password)
	loginResponse := protocol.NewObLoginResponse()
	err := c.Execute(context.TODO(), loginRequest, loginResponse)
	if err != nil {
		c.Close()
		return errors.WithMessagef(err, "execute login, uniqueId: %d remote addr: %s tenantname: %s databasename: %s",
			c.uniqueId, c.conn.RemoteAddr().String(), c.option.tenantName, c.option.databaseName)
	}

	c.credential = loginResponse.Credential()
	c.tenantId = loginResponse.TenantId()

	// TODO active = true
	c.active.Store(true)
	return nil
}

func (c *Connection) Execute(ctx context.Context, request protocol.ObPayload, response protocol.ObPayload) error {
	seq := c.seq.Add(1)

	request.SetUniqueId(c.uniqueId)
	request.SetSequence(c.sequence.Add(1))

	request.SetTenantId(c.tenantId)
	request.SetCredential(c.credential)

	payloadBuf := c.encodePayload(request)

	rpcHeaderBuf := c.encodeRpcHeader(request, payloadBuf)

	done := make(chan *Call, 1)
	call := new(Call)
	call.Done = done

	err := c.sendPacket(call, seq, rpcHeaderBuf, payloadBuf)
	if err != nil {
		return errors.WithMessage(err, "send packet")
	}

	ctx, _ = context.WithTimeout(ctx, 10*time.Second) // todo temporary use

	// wait call back
	select {
	case <-ctx.Done():
		// timeout
		c.mutex.Lock()
		delete(c.pending, seq)
		c.mutex.Unlock()
		return errors.WithMessage(ctx.Err(), "wait transport packet")
	case call = <-call.Done:
		if call.Error != nil { // transport failed
			return errors.WithMessage(call.Error, "receive packet")
		}
	}

	// transport success
	contentBuf := call.Content
	contentBuffer := bytes.NewBuffer(contentBuf)

	rpcHeader := c.decodeRpcHeader(contentBuffer)

	rpcResponseCode := protocol.NewObRpcResponseCode()

	rpcResponseCode.Decode(contentBuffer)

	if rpcResponseCode.Code() != oberror.ObSuccess {
		return oberror.NewProtocolError(
			c.option.ip,
			c.option.port,
			rpcResponseCode.Code(),
			rpcHeader.TraceId1(),
			rpcHeader.TraceId0(),
			"",
		)
	}

	response.SetUniqueId(rpcHeader.TraceId0())
	response.SetSequence(rpcHeader.TraceId1())

	c.decodePayload(response, contentBuffer)

	return nil
}

func (c *Connection) receivePacket() {
	defer c.Close()
	for {
		ezHeaderBuf := make([]byte, protocol.EzHeaderLength)
		_, err := io.ReadFull(c.conn, ezHeaderBuf)
		if err != nil {
			fmt.Printf("failed to connection read ezHeader, connection uniqueId: %d error: %s\n", c.uniqueId, err.Error())
			return
		}

		ezHeader := protocol.NewEzHeader()
		ezHeaderBuffer := bytes.NewBuffer(ezHeaderBuf)
		err = ezHeader.Decode(ezHeaderBuffer)
		if err != nil {
			fmt.Printf("failed to decode ezHeader, connection uniqueId: %d error: %s\n", c.uniqueId, err.Error())
			return
		}

		var call *Call

		contentLen := ezHeader.ContentLen()
		channelId := ezHeader.ChannelId()

		// TODO Use buf pool optimization
		contentBuf := make([]byte, contentLen)
		_, err = io.ReadFull(c.conn, contentBuf)
		if err != nil {
			// read failed
			c.mutex.Lock()
			call = c.pending[channelId]
			delete(c.pending, channelId)
			c.mutex.Unlock()
			call.Error = err
			call.done()

			fmt.Printf("failed to connection read content, connection uniqueId: %d error: %s\n", c.uniqueId, err.Error())
			return
		}

		// read success
		c.mutex.Lock()
		call = c.pending[channelId]
		delete(c.pending, channelId)
		c.mutex.Unlock()

		// call already deleted
		if call == nil {
			fmt.Printf("failed to not found table packet, connection uniqueId: %d seq: %d\n", c.uniqueId, channelId)
			continue
		}
		call.Content = contentBuf
		call.done()
	}
}

func (c *Connection) sendPacket(call *Call, seq uint32, rpcHeaderBuf []byte, payloadBuf []byte) error {
	rpcHeaderLen := len(rpcHeaderBuf)
	payloadLen := len(payloadBuf)

	ezHeader := protocol.NewEzHeader()
	ezHeader.SetChannelId(seq)
	ezHeader.SetContentLen(uint32(rpcHeaderLen + payloadLen))

	ezHeaderBuf := ezHeader.Encode()

	ezHeaderLen := len(ezHeaderBuf)
	totalLen := ezHeaderLen + rpcHeaderLen + payloadLen

	// TODO Use buf pool optimization
	totalBuf := make([]byte, totalLen)
	copy(totalBuf[:ezHeaderLen], ezHeaderBuf)
	copy(totalBuf[ezHeaderLen:ezHeaderLen+rpcHeaderLen], rpcHeaderBuf)
	copy(totalBuf[ezHeaderLen+rpcHeaderLen:], payloadBuf)

	c.mutex.Lock()
	c.pending[seq] = call
	c.mutex.Unlock()

	_, err := c.conn.Write(totalBuf)
	if err != nil {
		// write failed
		c.mutex.Lock()
		delete(c.pending, seq)
		c.mutex.Unlock()
		c.Close()
		return errors.WithMessage(err, "conn write")
	}
	// write success
	return nil
}

func (c *Connection) Close() {
	c.active.Store(false)
	c.conn.Close()
}

func (c *Connection) encodePayload(payload protocol.ObPayload) []byte {
	payloadLen := payload.PayloadLen()
	payloadBuf := make([]byte, payloadLen)
	payloadBuffer := bytes.NewBuffer(payloadBuf)
	payload.Encode(payloadBuffer)
	return payloadBuf
}

func (c *Connection) encodeRpcHeader(payload protocol.ObPayload, payloadBuf []byte) []byte {
	rpcHeader := protocol.NewObRpcHeader()
	rpcHeader.SetPCode(payload.PCode().Value())
	rpcHeader.SetFlag(payload.Flag())
	rpcHeader.SetTenantId(payload.TenantId())
	rpcHeader.SetSessionId(payload.SessionId())
	rpcHeader.SetTimeout(payload.Timeout())
	rpcHeader.SetTraceId0(payload.UniqueId())
	rpcHeader.SetTraceId1(payload.Sequence())
	// TODO To be added
	// rpcHeader.SetPriority(0)
	rpcHeader.SetChecksum(util.Calculate(0, payloadBuf))

	rpcHeaderBuf := rpcHeader.Encode()
	return rpcHeaderBuf
}

func (c *Connection) decodeRpcHeader(contentBuffer *bytes.Buffer) *protocol.ObRpcHeader {
	rpcHeader := protocol.NewObRpcHeader()
	rpcHeader.Decode(contentBuffer)
	return rpcHeader
}

func (c *Connection) decodePayload(payload protocol.ObPayload, contentBuffer *bytes.Buffer) {
	payload.Decode(contentBuffer)
}

func (call *Call) done() {
	select {
	case call.Done <- call:
		// ok
	default:
		fmt.Printf("rpc: discarding Call reply due to insufficient Done chan capacity\n")
	}
}
