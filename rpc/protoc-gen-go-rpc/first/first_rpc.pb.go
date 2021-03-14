// Code generated by protoc-gen-go-rpc. DO NOT EDIT.

package first

import (
	second "./second"
	context "context"
	rpc "github.com/chenjie199234/Corelib/rpc"
	error1 "github.com/chenjie199234/Corelib/util/error"
	metadata "github.com/chenjie199234/Corelib/util/metadata"
	proto "github.com/golang/protobuf/proto"
	strconv "strconv"
	time "time"
)

var RpcPathFirstHello = "/first.First/Hello"
var RpcPathFirstWorld = "/first.First/World"

// FirstRpcClient is the client API for First service.
type FirstRpcClient interface {
	Hello(context.Context, *Helloreq) (*Helloresp, error)
	World(context.Context, *Worldreq) (*Worldresp, error)
}

type firstRpcClient struct {
	cc *rpc.RpcClient
}

//has race,will only return the first's call's client,the config will use the first call's config
func NewFirstRpcClient(timeout, conntimeout, hearttimeout, heartprobe time.Duration, selfgroup, selfname string, verifydata []byte, picker rpc.PickHandler, discover rpc.DiscoveryHandler) (FirstRpcClient, error) {
	c := &rpc.Config{
		Timeout:                time.Duration(timeout),
		ConnTimeout:            time.Duration(conntimeout),
		HeartTimeout:           time.Duration(hearttimeout),
		HeartPorbe:             time.Duration(heartprobe),
		GroupNum:               1,
		SocketRBuf:             1024,
		SocketWBuf:             1024,
		MaxMsgLen:              65535,
		MaxBufferedWriteMsgNum: 1024,
	}
	cc, e := rpc.NewRpcClient(c, selfgroup, selfname, verifydata, Group, Name, picker, discover)
	if e != nil {
		return nil, e
	}
	return &firstRpcClient{cc: cc}, nil
}

func (c *firstRpcClient) Hello(ctx context.Context, req *Helloreq) (*Helloresp, error) {
	if req == nil {
		return nil, rpc.ERRREQUEST
	}
	//gt check
	if float64(req.I32) <= 6 {
		return nil, rpc.ERRREQUEST
	}
	//not in check
	if vv := strconv.FormatInt(int64(req.I32), 10); vv == "8" ||
		vv == "9" {
		return nil, rpc.ERRREQUEST
	}
	//empty check
	if len(req.Ri32) == 0 {
		return nil, rpc.ERRREQUEST
	}
	//in check
	for _, v := range req.Ri32 {
		if vv := strconv.FormatInt(int64(v), 10); vv != "1" &&
			vv != "2" &&
			vv != "3" {
			return nil, rpc.ERRREQUEST
		}
	}
	//gt check
	if float64(req.Ui32) <= 6 {
		return nil, rpc.ERRREQUEST
	}
	//not in check
	if vv := strconv.FormatUint(uint64(req.Ui32), 10); vv == "8" ||
		vv == "9" {
		return nil, rpc.ERRREQUEST
	}
	//empty check
	if len(req.Rui32) == 0 {
		return nil, rpc.ERRREQUEST
	}
	//in check
	for _, v := range req.Rui32 {
		if vv := strconv.FormatUint(uint64(v), 10); vv != "1" &&
			vv != "2" &&
			vv != "3" {
			return nil, rpc.ERRREQUEST
		}
	}
	//empty check
	if len(req.Bs) == 0 {
		return nil, rpc.ERRREQUEST
	}
	//empty check
	if len(req.Rbs) == 0 {
		return nil, rpc.ERRREQUEST
	}
	//empty check
	if len(req.Ss) == 0 {
		return nil, rpc.ERRREQUEST
	}
	//empty check
	if len(req.Rss) == 0 {
		return nil, rpc.ERRREQUEST
	}
	//enum check
	if _, ok := second.ABC_name[int32(req.E)]; !ok {
		return nil, rpc.ERRREQUEST
	}
	//enum check
	for _, v := range req.Re {
		if _, ok := second.ABC_name[int32(v)]; !ok {
			return nil, rpc.ERRREQUEST
		}
	}
	if req.A != nil {
		//gt check
		if float64(req.A.Uid) <= 0 {
			return nil, rpc.ERRREQUEST
		}
	}
	for i := range req.Ra {
		if req.Ra[i] != nil {
			//gt check
			if float64(req.Ra[i].Uid) <= 0 {
				return nil, rpc.ERRREQUEST
			}
		}
	}
	//enum check
	for _, v := range req.Me {
		if _, ok := second.ABC_name[int32(v)]; !ok {
			return nil, rpc.ERRREQUEST
		}
	}
	//empty check
	if len(req.Me) == 0 {
		return nil, rpc.ERRREQUEST
	}
	//empty check
	if len(req.Mb) == 0 {
		return nil, rpc.ERRREQUEST
	}
	reqd, _ := proto.Marshal(req)
	callback, e := c.cc.Call(ctx, 200000000, RpcPathFirstHello, reqd, metadata.GetAllMetadata(ctx))
	if e.(*error1.Error) != nil {
		return nil, e
	}
	resp := new(Helloresp)
	if e := proto.Unmarshal(callback, resp); e != nil {
		return nil, rpc.ERRRESPONSE
	}
	return resp, nil
}
func (c *firstRpcClient) World(ctx context.Context, req *Worldreq) (*Worldresp, error) {
	if req == nil {
		return nil, rpc.ERRREQUEST
	}
	reqd, _ := proto.Marshal(req)
	callback, e := c.cc.Call(ctx, 0, RpcPathFirstWorld, reqd, metadata.GetAllMetadata(ctx))
	if e.(*error1.Error) != nil {
		return nil, e
	}
	resp := new(Worldresp)
	if e := proto.Unmarshal(callback, resp); e != nil {
		return nil, rpc.ERRRESPONSE
	}
	return resp, nil
}

// FirstRpcServer is the server API for First service.
type FirstRpcServer interface {
	Hello(context.Context, *Helloreq) (*Helloresp, error)
	World(context.Context, *Worldreq) (*Worldresp, error)
}

func _First_Hello_RpcHandler(handler func(context.Context, *Helloreq) (*Helloresp, error)) rpc.OutsideHandler {
	return func(ctx *rpc.Context) {
		req := new(Helloreq)
		if e := proto.Unmarshal(ctx.GetBody(), req); e != nil {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//gt check
		if float64(req.I32) <= 6 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//not in check
		if vv := strconv.FormatInt(int64(req.I32), 10); vv == "8" ||
			vv == "9" {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//empty check
		if len(req.Ri32) == 0 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//in check
		for _, v := range req.Ri32 {
			if vv := strconv.FormatInt(int64(v), 10); vv != "1" &&
				vv != "2" &&
				vv != "3" {
				ctx.Abort(rpc.ERRREQUEST)
				return
			}
		}
		//gt check
		if float64(req.Ui32) <= 6 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//not in check
		if vv := strconv.FormatUint(uint64(req.Ui32), 10); vv == "8" ||
			vv == "9" {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//empty check
		if len(req.Rui32) == 0 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//in check
		for _, v := range req.Rui32 {
			if vv := strconv.FormatUint(uint64(v), 10); vv != "1" &&
				vv != "2" &&
				vv != "3" {
				ctx.Abort(rpc.ERRREQUEST)
				return
			}
		}
		//empty check
		if len(req.Bs) == 0 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//empty check
		if len(req.Rbs) == 0 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//empty check
		if len(req.Ss) == 0 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//empty check
		if len(req.Rss) == 0 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//enum check
		if _, ok := second.ABC_name[int32(req.E)]; !ok {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//enum check
		for _, v := range req.Re {
			if _, ok := second.ABC_name[int32(v)]; !ok {
				ctx.Abort(rpc.ERRREQUEST)
				return
			}
		}
		if req.A != nil {
			//gt check
			if float64(req.A.Uid) <= 0 {
				ctx.Abort(rpc.ERRREQUEST)
				return
			}
		}
		for i := range req.Ra {
			if req.Ra[i] != nil {
				//gt check
				if float64(req.Ra[i].Uid) <= 0 {
					ctx.Abort(rpc.ERRREQUEST)
					return
				}
			}
		}
		//enum check
		for _, v := range req.Me {
			if _, ok := second.ABC_name[int32(v)]; !ok {
				ctx.Abort(rpc.ERRREQUEST)
				return
			}
		}
		//empty check
		if len(req.Me) == 0 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		//empty check
		if len(req.Mb) == 0 {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		resp, e := handler(ctx, req)
		if e != nil {
			ctx.Abort(e)
			return
		}
		if resp == nil {
			resp = new(Helloresp)
		}
		respd, _ := proto.Marshal(resp)
		ctx.Write(respd)
	}
}
func _First_World_RpcHandler(handler func(context.Context, *Worldreq) (*Worldresp, error)) rpc.OutsideHandler {
	return func(ctx *rpc.Context) {
		req := new(Worldreq)
		if e := proto.Unmarshal(ctx.GetBody(), req); e != nil {
			ctx.Abort(rpc.ERRREQUEST)
			return
		}
		resp, e := handler(ctx, req)
		if e != nil {
			ctx.Abort(e)
			return
		}
		if resp == nil {
			resp = new(Worldresp)
		}
		respd, _ := proto.Marshal(resp)
		ctx.Write(respd)
	}
}
func RegisterFirstRpcServer(engine *rpc.RpcServer, svc FirstRpcServer, allmids map[string]rpc.OutsideHandler) error {
	//avoid lint
	_ = allmids
	{
		requiredMids := []string{"auth", "limit"}
		mids := make([]rpc.OutsideHandler, 0)
		for _, v := range requiredMids {
			if mid, ok := allmids[v]; ok {
				mids = append(mids, mid)
			}
		}
		mids = append(mids, _First_Hello_RpcHandler(svc.Hello))
		if e := engine.RegisterHandler(RpcPathFirstHello, 200000000, mids...); e != nil {
			return e
		}
	}
	if e := engine.RegisterHandler(RpcPathFirstWorld, 0, _First_World_RpcHandler(svc.World)); e != nil {
		return e
	}
	return nil
}
