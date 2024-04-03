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
	"io"
	"runtime/debug"
	"sync"
	"time"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
)

// CodecServer implement interfaces to read/decode request
// and write/encode response
type CodecServer interface {
	ReadRequest(*Request) error
	WriteResponse(*Response) error
	Call(*Request, *Response) error
	Close()
}

// Server implement server frame
type Server struct {
	MethodRunPool *ants.Pool
	ReqObjPool    *sync.Pool
	RespObjPool   *sync.Pool
	CloseChan     *chan struct{} // to close and clean Server
}

// Request is generated by a decoder
type Request struct {
	Method string // use for mapping
	Args   [][]byte
	ID     string
}

type Response struct {
	ID         string
	RspContent []byte
}

// NewServer init a server
func NewServer(routinePoolSize int, expiredDuration time.Duration, ch *chan struct{}) (*Server, error) {
	var err error
	s := &Server{CloseChan: ch}
	s.MethodRunPool, err = ants.NewPool(routinePoolSize,
		ants.WithExpiryDuration(expiredDuration),
		ants.WithPanicHandler(func(p interface{}) {
			if err := recover(); err != nil {
				log.Error("RPCServer", nil, "catch panic", log.Any("error", err), log.String("stack", string(debug.Stack())))
			}
		}),
	)
	if err != nil {
		log.Warn("RPCServer", nil, "create goroutine pool failed", zap.Error(err))
		return s, err
	}
	s.ReqObjPool = &sync.Pool{
		New: func() interface{} {
			return new(Request)
		},
	}
	s.RespObjPool = &sync.Pool{
		New: func() interface{} {
			return new(Response)
		},
	}
	return s, nil
}

// PutRequest clear req and put back to req pool
func (s *Server) PutRequest(req *Request) {
	req.Method = ""
	req.ID = ""
	for i := range req.Args {
		req.Args[i] = req.Args[i][:0]
	}
	req.Args = req.Args[:0]
	// req.Err = nil
	s.ReqObjPool.Put(req)
}

// PutRequest clear req and put back to req pool
func (s *Server) PutResponse(resp *Response) {
	resp.RspContent = resp.RspContent[:0]
	resp.ID = ""
	s.RespObjPool.Put(resp)
}

// Close release resources
func (s *Server) Close() {
	s.MethodRunPool.Release()
}

// ReadRequestWrapper wrap ReadRequest for goroutine pool
type ReadRequestWrapper func()

// ReadRequest read requests until an error occurs
func (s *Server) ReadRequest(wg *sync.WaitGroup, reqChan chan<- *Request, cServer CodecServer) ReadRequestWrapper {
	return func() {
		defer wg.Done()
		for {
			isStop := false
			select {
			case <-*s.CloseChan:
				cServer.Close()
				s.Close()
				return
			default:
				// ReadRequest may include read and encode, depend on the cServer
				req := s.ReqObjPool.Get().(*Request)
				err := cServer.ReadRequest(req)
				if err != nil {
					if err == io.EOF {
						log.Info("RPCServer", req.ID, "connection closed", zap.Error(err))
					} else {
						log.Warn("RPCServer", req.ID, "fail to read command", zap.Error(err))
					}
					s.PutRequest(req)
					close(reqChan)
					isStop = true
				} else {
					reqChan <- req
				}
			}
			if isStop {
				break
			}
		}
	}
}

// RunWorkerWrapper wrap RunWorker for goroutine pool
type RunWorkerWrapper func()

// RunWorker keep processing requests from reqChan until channel reqChan is closed
func (s *Server) RunWorker(wg *sync.WaitGroup, reqChan <-chan *Request, cServer CodecServer) RunWorkerWrapper {
	return func() {
		defer func() {
			wg.Done()
			if err := recover(); err != nil {
				log.Error("RPCServer", nil, "RunWorker panic", log.Any("error", err), log.String("stack", string(debug.Stack())))
			}
		}()
		for req := range reqChan {
			resp := s.RespObjPool.Get().(*Response)
			err := cServer.Call(req, resp)
			if err != nil {
				log.Warn("RPCServer", req.ID, "fail to call method", zap.Error(err))
			}
			s.PutRequest(req)
			err = cServer.WriteResponse(resp)
			s.PutResponse(resp)
			if err != nil {
				log.Warn("RPCServer", nil, "fail to write response", zap.Error(err))
				break
			}
		}
	}
}

// ServeCodec serve a codec
func (s *Server) ServeCodec(cServer CodecServer) {
	wg := new(sync.WaitGroup)
	defer func() {
		wg.Wait()
		cServer.Close()
		if err := recover(); err != nil {
			log.Error("RPCServer", nil, "ServeCodec panic", log.Any("error", err), log.String("stack", string(debug.Stack())))
		}
	}()
	reqChan := make(chan *Request, 100)
	var err error

	// exec(request) -> Response
	wg.Add(1)
	if err = s.MethodRunPool.Submit(s.RunWorker(wg, reqChan, cServer)); err != nil {
		wg.Done()
		return
	}

	for {
		isStop := false
		select {
		case <-*s.CloseChan:
			cServer.Close()
			s.Close()
			return
		default:
			// ReadRequest may include read and encode, depend on the cServer
			req := s.ReqObjPool.Get().(*Request)
			err := cServer.ReadRequest(req)
			if err != nil {
				if err == io.EOF {
					log.Info("RPCServer", req.ID, "connection closed", zap.Error(err))
				} else {
					log.Warn("RPCServer", req.ID, "fail to read command", zap.Error(err))
				}
				s.PutRequest(req)
				close(reqChan)
				isStop = true
			} else {
				reqChan <- req
			}
		}
		if isStop {
			break
		}
	}
}
