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
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	oberror "github.com/oceanbase/obkv-table-client-go/error"
	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/util"
)

const (
	connReaderBufferSize = 4 * 1024
	connWriterBufferSize = 4 * 1024
	minBatchWriteSize    = 1024

	connSystemReadBufferSize  = 128 * 1024
	connSystemWriteBufferSize = 64 * 1024
)

var bufferPool = NewLimitedPool(256, 8192)

var ezHeaderPool = sync.Pool{New: func() any {
	return make([]byte, protocol.EzHeaderLength)
}}

type ConnectionOption struct {
	ip             string
	port           int
	connectTimeout time.Duration
	loginTimeout   time.Duration

	tenantName   string
	databaseName string
	userName     string
	password     string
}

func NewConnectionOption(ip string, port int, connectTimeout time.Duration, loginTimeout time.Duration,
	tenantName string, databaseName string, userName string, password string) *ConnectionOption {
	return &ConnectionOption{
		ip:             ip,
		port:           port,
		connectTimeout: connectTimeout,
		loginTimeout:   loginTimeout,
		tenantName:     tenantName,
		databaseName:   databaseName,
		userName:       userName,
		password:       password,
	}
}

type Connection struct {
	option *ConnectionOption

	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer

	packetChannel      chan *packet
	packetChannelClose chan struct{}
	closeOnce          sync.Once

	mutex sync.Mutex
	seq   atomic.Uint32 // as channel id in ez header

	pending map[uint32]*call
	active  atomic.Bool

	uniqueId   uint64        // as trace0 in rpc header
	sequence   atomic.Uint64 // as trace1 in rpc header
	credential []byte
	tenantId   uint64

	ezHeaderLength  int
	rpcHeaderLength int
}

type packet struct {
	seq  uint32
	data []byte
}

// call represents an active RPC.
type call struct {
	err     error
	signal  chan *call // Strobes when call is complete.
	content []byte
}

const (
	defaultConnPendingSize   = 1024
	defaultPacketChannelSize = 1024
)

func NewConnection(option *ConnectionOption) *Connection {
	return &Connection{option: option, pending: make(map[uint32]*call, defaultConnPendingSize),
		packetChannel: make(chan *packet, defaultPacketChannelSize), packetChannelClose: make(chan struct{})}
}

func (c *Connection) Connect(ctx context.Context) error {
	address := fmt.Sprintf("%s:%s", c.option.ip, strconv.Itoa(c.option.port))
	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return errors.WithMessagef(err, "net dial, uniqueId: %d remote addr: %s", c.uniqueId, address)
	}
	c.conn = conn
	c.conn.(*net.TCPConn).SetNoDelay(true)
	c.conn.(*net.TCPConn).SetReadBuffer(connSystemReadBufferSize)
	c.conn.(*net.TCPConn).SetWriteBuffer(connSystemWriteBufferSize)
	c.reader = bufio.NewReaderSize(c.conn, connReaderBufferSize)
	c.writer = bufio.NewWriterSize(c.conn, connWriterBufferSize)

	// ez header length rpc header length
	c.ezHeaderLength = protocol.EzHeaderLength
	c.rpcHeaderLength = protocol.RpcHeaderEncodeSize
	if util.ObVersion() >= 4 {
		c.rpcHeaderLength = protocol.RpcHeaderEncodeSizeV4
	}

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
	go c.sendPacket()
	return nil
}

func (c *Connection) Login(ctx context.Context) error {
	loginRequest := protocol.NewObLoginRequest(c.option.tenantName, c.option.databaseName, c.option.userName, c.option.password)
	loginResponse := protocol.NewObLoginResponse()
	err := c.Execute(ctx, loginRequest, loginResponse)
	if err != nil {
		c.Close()
		return errors.WithMessagef(err, "execute login, uniqueId: %d remote addr: %s tenantname: %s databasename: %s",
			c.uniqueId, c.conn.RemoteAddr().String(), c.option.tenantName, c.option.databaseName)
	}

	c.credential = loginResponse.Credential()
	c.tenantId = loginResponse.TenantId()

	c.active.Store(true)
	return nil
}

func (c *Connection) Execute(ctx context.Context, request protocol.ObPayload, response protocol.ObPayload) error {
	seq := c.seq.Add(1)

	totalBuf := c.encodePacket(seq, request)

	call := &call{
		err:     nil,
		signal:  make(chan *call, 1),
		content: nil, // call back to user goroutine content
	}

	packet := &packet{
		seq:  seq,
		data: totalBuf,
	}

	c.mutex.Lock()
	c.pending[seq] = call
	c.mutex.Unlock()

	select {
	case c.packetChannel <- packet: // write packet channel success, but not equal to connection write success.
	case <-ctx.Done():
		// timeout
		c.mutex.Lock()
		delete(c.pending, seq)
		c.mutex.Unlock()
		return errors.WithMessage(ctx.Err(), "wait send packet to channel")
	}

	// wait call back
	select {
	case call = <-call.signal:
		if call.err != nil { // transport failed
			return errors.WithMessage(call.err, "receive packet")
		}
	case <-ctx.Done():
		// timeout
		c.mutex.Lock()
		delete(c.pending, seq)
		c.mutex.Unlock()
		return errors.WithMessage(ctx.Err(), "wait transport packet")
	}

	// transport success
	err := c.decodePacket(call.content, response)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) receivePacket() {
	defer c.Close()
	for {
		ezHeaderBuf := ezHeaderPool.Get().([]byte)

		_, err := io.ReadFull(c.reader, ezHeaderBuf)
		if err != nil {
			log.Warn("failed to connection read header", zap.Error(err), zap.Uint64("uniqueId", c.uniqueId))
			return
		}

		ezHeader := protocol.NewEzHeader()
		err = ezHeader.Decode(ezHeaderBuf)
		if err != nil {
			log.Warn("failed to decode ezHeader", zap.Error(err), zap.Uint64("uniqueId", c.uniqueId))
			return
		}

		ezHeaderPool.Put(ezHeaderBuf)

		var call *call

		contentLen := ezHeader.ContentLen()
		channelId := ezHeader.ChannelId()

		contentBuf := *bufferPool.Get(int(contentLen)) // reuse

		_, err = io.ReadFull(c.reader, contentBuf)
		if err != nil {
			// read failed
			c.mutex.Lock()
			call = c.pending[channelId]
			delete(c.pending, channelId)
			c.mutex.Unlock()
			if call == nil {
				log.Warn("failed to not found table packet", zap.Uint64("uniqueId", c.uniqueId), zap.Uint32("seq", channelId))
			} else {
				call.err = err
				call.done()
			}
			log.Warn("failed to connection read content", zap.Error(err), zap.Uint64("uniqueId", c.uniqueId))
			return
		}

		// read success
		c.mutex.Lock()
		call = c.pending[channelId]
		delete(c.pending, channelId)
		c.mutex.Unlock()

		// call already deleted
		if call == nil {
			log.Warn("failed to not found table packet", zap.Uint64("uniqueId", c.uniqueId), zap.Uint32("seq", channelId))
			continue
		}
		call.content = contentBuf
		call.done()
	}
}

func (c *Connection) sendPacket() {
	var packet *packet
	var gosched = false
	for {
		select {
		case packet = <-c.packetChannel:
		case <-c.packetChannelClose:
			// clear packet channel
			for {
				select {
				case packet = <-c.packetChannel:
					var call *call
					c.mutex.Lock()
					call = c.pending[packet.seq]
					delete(c.pending, packet.seq)
					c.mutex.Unlock()
					if call != nil {
						call.err = errors.New("send packet channel is close")
						call.done()
					}
				default:
					log.Warn("send packet channel closed", zap.Uint64("uniqueId", c.uniqueId))
					return
				}
			}
		}

		gosched = true
		c.writerWrite(packet)

	hasPacket:
		for { // write all packages at once reduce syscall.
			select {
			case packet = <-c.packetChannel:
				c.writerWrite(packet)
			default:
				break hasPacket
			}
		}

		if gosched { // only once true
			gosched = false
			if c.writer.Buffered() < minBatchWriteSize {
				runtime.Gosched()
				goto hasPacket
			}
		}

		c.writer.Flush()
	}
}

func (c *Connection) writerWrite(packet *packet) {
	_, err := c.writer.Write(packet.data)
	if err != nil {
		// write failed
		var call *call
		c.mutex.Lock()
		call = c.pending[packet.seq]
		delete(c.pending, packet.seq)
		c.mutex.Unlock()
		if call != nil {
			call.err = err
			call.done()
		}
		c.Close()
	}
	bufferPool.Put(&packet.data)
}

func (c *Connection) Close() {
	c.active.Store(false)
	c.closeOnce.Do(func() {
		close(c.packetChannelClose) // close packet channel
		c.conn.Close()
		// clear pending call
		c.mutex.Lock()
		for seq, call := range c.pending {
			delete(c.pending, seq)
			call.err = errors.New("connection close")
			call.done()
		}
		c.mutex.Unlock()
	})
}

func (c *Connection) encodePacket(seq uint32, request protocol.ObPayload) []byte {
	// encode request
	request.SetUniqueId(c.uniqueId)
	request.SetSequence(c.sequence.Add(1))
	request.SetTenantId(c.tenantId)
	request.SetCredential(c.credential)
	payloadLen := request.PayloadLen()

	totalLength := payloadLen + c.rpcHeaderLength + c.ezHeaderLength // total length
	totalBuf := *bufferPool.Get(totalLength)                         // only once get buf

	payloadBuf := totalBuf[c.ezHeaderLength+c.rpcHeaderLength:] // payload buf
	payloadBuffer := bytes.NewBuffer(payloadBuf)
	request.Encode(payloadBuffer)

	// encode rpc header
	rpcHeaderBuf := totalBuf[c.ezHeaderLength : c.ezHeaderLength+c.rpcHeaderLength] // rpc header buf
	rpcHeader := protocol.NewObRpcHeader()
	rpcHeader.SetPCode(request.PCode().Value())
	rpcHeader.SetFlag(request.Flag())
	rpcHeader.SetTenantId(request.TenantId())
	rpcHeader.SetSessionId(request.SessionId())
	rpcHeader.SetTimeout(request.Timeout())
	rpcHeader.SetTraceId0(request.UniqueId())
	rpcHeader.SetTraceId1(request.Sequence())
	rpcHeader.SetChecksum(util.Calculate(0, payloadBuf))
	rpcHeaderBuffer := bytes.NewBuffer(rpcHeaderBuf)
	rpcHeader.SetHLen(uint8(c.rpcHeaderLength))
	rpcHeader.Encode(rpcHeaderBuffer)

	// encode ez header
	ezHeaderBuf := totalBuf[:c.ezHeaderLength] // ez header buf
	ezHeader := protocol.NewEzHeader()
	ezHeader.SetChannelId(seq)
	ezHeader.SetContentLen(uint32(c.rpcHeaderLength + payloadLen))
	ezHeader.Encode(ezHeaderBuf)

	return totalBuf
}

func (c *Connection) decodePacket(contentBuf []byte, response protocol.ObPayload) error {
	contentBuffer := bytes.NewBuffer(contentBuf)

	// decode rpc header
	rpcHeader := protocol.NewObRpcHeader()
	rpcHeader.Decode(contentBuffer)

	// decode rpc response code
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

	// decode response
	response.SetUniqueId(rpcHeader.TraceId0())
	response.SetSequence(rpcHeader.TraceId1())
	response.Decode(contentBuffer)

	bufferPool.Put(&contentBuf) // reuse
	return nil
}

func (call *call) done() {
	select {
	case call.signal <- call:
		// ok
	default:
		log.Warn("rpc: discarding call reply due to insufficient signal chan capacity")
	}
}
