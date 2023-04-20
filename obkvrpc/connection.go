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

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/protocol"
	"github.com/oceanbase/obkv-table-client-go/util"
)

type Option struct {
	ip             string
	port           int
	connectTimeout time.Duration

	tenantName   string
	databaseName string
	userName     string
	password     string
}

func NewOption(ip string, port int, connectTimeout time.Duration, tenantName string, databaseName string, userName string, password string) *Option {
	return &Option{ip: ip, port: port, connectTimeout: connectTimeout, tenantName: tenantName, databaseName: databaseName, userName: userName, password: password}
}

type Connection struct {
	option *Option

	conn    net.Conn
	mutex   sync.Mutex
	seq     atomic.Uint32
	pending map[uint32]*Call
	active  atomic.Bool

	uuid           uuid.UUID
	traceIdCounter atomic.Uint32

	credential []byte
	tenantId   uint64
}

// Call represents an active RPC.
type Call struct {
	Error   error
	Done    chan *Call // Strobes when call is complete.
	Content []byte
}

func NewConnection(option *Option, uuid uuid.UUID) *Connection {
	return &Connection{option: option, uuid: uuid, pending: make(map[uint32]*Call)}
}

func (c *Connection) Connect() error {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", c.option.ip, strconv.Itoa(c.option.port)), c.option.connectTimeout)
	if err != nil {
		return errors.Wrap(err, "tcp connect failed")
	}
	c.conn = conn

	go c.receivePacket()
	return nil
}

func (c *Connection) Login() {
	loginRequest := protocol.NewLoginRequest(c.option.tenantName, c.option.databaseName, c.option.userName, c.option.password)
	loginResponse := protocol.NewLoginResponse()
	err := c.Execute(context.TODO(), loginRequest, loginResponse)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// TODO active = true
}

func (c *Connection) Execute(ctx context.Context, request protocol.Payload, response protocol.Payload) error {
	seq := c.seq.Add(1)

	request.SetTenantId(c.tenantId)
	request.SetCredential(c.credential)

	payloadBuf := c.encodePayload(request)

	rpcHeaderBuf := c.encodeRpcHeader(request, payloadBuf)

	done := make(chan *Call, 1)
	call := new(Call)
	call.Done = done

	c.sendPacket(call, seq, rpcHeaderBuf, payloadBuf)

	// wait call back
	select {
	case <-ctx.Done():
		// timeout
		c.mutex.Lock()
		delete(c.pending, seq)
		c.mutex.Unlock()
		return errors.Wrap(ctx.Err(), "send request and receive response")
	case call = <-call.Done:
		if call.Error != nil { // transport failed
			return errors.Wrap(call.Error, "send request and receive response")
		}
	}

	// transport success
	contentBuf := call.Content
	contentBuffer := bytes.NewBuffer(contentBuf)

	c.decodeRpcHeader(contentBuffer)

	payloadBuf = contentBuffer.Bytes()

	// TODO rpcResponseCode
	rpcResponseCode := protocol.NewRpcResponseCode()

	rpcResponseCode.Decode(contentBuffer)

	if rpcResponseCode.Code() != protocol.ObSuccess {
		fmt.Printf("failed to rpc response code not success, code: %d\n", rpcResponseCode.Code())
		return errors.Errorf("rpc response code not success,code : %d", rpcResponseCode.Code())
	}

	c.decodePayload(response, contentBuffer)

	return nil
}

func (c *Connection) receivePacket() {
	defer c.Close()
	for {
		ezHeaderBuf := make([]byte, protocol.EzHeaderLength)
		_, err := io.ReadFull(c.conn, ezHeaderBuf)
		if err != nil {
			fmt.Printf("failed to tcp connection read ezHeader, connection uuid: %d error: %s\n", c.uuid, err.Error())
			return
		}

		ezHeader := protocol.NewEzHeader()
		ezHeaderBuffer := bytes.NewBuffer(ezHeaderBuf)
		err = ezHeader.Decode(ezHeaderBuffer)
		if err != nil {
			fmt.Printf("failed to decode ezHeader, connection uuid: %d error: %s\n", c.uuid, err.Error())
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

			fmt.Printf("failed to tcp connection read content, connection uuid: %d error: %s\n", c.uuid, err.Error())
			return
		}

		// read success
		c.mutex.Lock()
		call = c.pending[channelId]
		delete(c.pending, channelId)
		c.mutex.Unlock()

		// call already deleted
		if call == nil {
			fmt.Printf("failed to not found table packet, connection uuid: %d seq: %d\n", c.uuid, channelId)
			continue
		}
		call.Content = contentBuf
		call.done()
	}
}

func (c *Connection) sendPacket(call *Call, seq uint32, rpcHeaderBuf []byte, payloadBuf []byte) {
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
		call.Error = err
		call.done()
		c.Close() // TODO
	}
	// write success
}

func (c *Connection) Close() {
	c.active.Store(false)
	c.conn.Close()
}

func (c *Connection) encodePayload(payload protocol.Payload) []byte {
	payloadLen := payload.PayloadLen()
	payloadBuf := make([]byte, payloadLen)
	payloadBuffer := bytes.NewBuffer(payloadBuf)
	payload.Encode(payloadBuffer)
	return payloadBuf
}

func (c *Connection) encodeRpcHeader(payload protocol.Payload, payloadBuf []byte) []byte {
	rpcHeader := protocol.NewRpcHeader()
	rpcHeader.SetPCode(payload.PCode().Value())
	rpcHeader.SetTimeout(payload.Timeout())
	rpcHeader.SetTenantId(payload.TenantId())
	rpcHeader.SetSessionId(payload.SessionId())
	rpcHeader.SetFlag(payload.Flag())
	rpcHeader.SetTraceId0(uint64(c.uuid.ID()))
	rpcHeader.SetTraceId1(uint64(c.traceIdCounter.Add(1)))
	// TODO To be added
	// rpcHeader.SetPriority(0)
	rpcHeader.SetChecksum(util.Calculate(0, payloadBuf))

	rpcHeaderBuf := rpcHeader.Encode()
	return rpcHeaderBuf
}

func (c *Connection) decodeRpcHeader(contentBuffer *bytes.Buffer) *protocol.RpcHeader {
	rpcHeader := protocol.NewRpcHeader()
	rpcHeader.Decode(contentBuffer)
	return rpcHeader
}

func (c *Connection) decodePayload(payload protocol.Payload, contentBuffer *bytes.Buffer) {
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
