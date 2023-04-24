package obkvrpc

import (
	"strconv"
)

type ObRpcClientOption struct {
	connPoolSize int
}

func (o *ObRpcClientOption) String() string {
	return "ObRpcClientOption{" +
		"connPoolSize:" + strconv.Itoa(o.connPoolSize) +
		"}"
}

func NewObRpcClientOption(connPoolSize int) *ObRpcClientOption {
	return &ObRpcClientOption{connPoolSize}
}

type ObRpcClient struct {
	opt *ObRpcClientOption
	// connPool
}

func (c *ObRpcClient) Init() error {
	// todo:impl
	return nil
}

func (c *ObRpcClient) Execute(request interface{}, result interface{}) error {
	// todo:impl
	return nil
}

func (c *ObRpcClient) Close() {
	// todo:impl
}

func (c *ObRpcClient) String() string {
	return "ObRpcClient{" +
		"opt:" + c.opt.String() +
		"}"
}

func NewObRpcClient(opt *ObRpcClientOption) *ObRpcClient {
	return &ObRpcClient{opt}
}
