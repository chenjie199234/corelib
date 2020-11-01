// Code generated by protoc-gen-go-mrpc. DO NOT EDIT.

package test

import (
	//std
	"context"

	//third
	"github.com/chenjie199234/Corelib/rpc/mrpc"
	"google.golang.org/protobuf/proto"
)

var PathMrpcTestHello = "/Test/Hello"

type MrpcTestClient struct {
	c *mrpc.MrpcClient
}

func (c *MrpcTestClient) Hello(ctx context.Context, req *HelloReq) (*HelloResp, error) {
	reqd, _ := proto.Marshal(req)
	respd, err := c.c.Call(ctx, "/Test/Hello", reqd)
	if err != nil {
		return nil, err
	}
	if respd == nil {
		return &HelloResp{}, nil
	}
	resp := new(HelloResp)
	if e := proto.Unmarshal(respd, resp); e != nil {
		//this is impossible
		return nil, mrpc.ERR[mrpc.ERRRESPONSE]
	}
	return resp, nil
}
func NewMrpcTestClient(c *mrpc.MrpcClient) *MrpcTestClient {
	return &MrpcTestClient{c: c}
}

type MrpcTestService struct {
	Midware func() map[string][]mrpc.OutsideHandler
	Hello   func(context.Context, *HelloReq) (*HelloResp, error)
}

func (s *MrpcTestService) hello(ctx context.Context, in []byte) ([]byte, error) {
	req := &HelloReq{}
	if e := proto.Unmarshal(in, req); e != nil {
		return nil, mrpc.ERR[mrpc.ERRREQUEST]
	}
	resp, err := s.Hello(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		d, _ := proto.Marshal(resp)
		return d, err
	}
	return nil, nil
}
func RegisterMrpcTestService(engine *mrpc.MrpcServer, instance *MrpcTestService) {
	var allmids map[string][]mrpc.OutsideHandler
	if instance.Midware != nil {
		allmids = instance.Midware()
	}
	//Hello
	if instance.Hello != nil {
		if mids, ok := allmids[PathMrpcTestHello]; ok && len(mids) != 0 {
			engine.RegisterHandler(PathMrpcTestHello, 200, append(mids, instance.hello)...)
		} else {
			engine.RegisterHandler(PathMrpcTestHello, 200, instance.hello)
		}
	}
}
