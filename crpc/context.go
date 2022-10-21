package crpc

import (
	"context"

	"github.com/chenjie199234/Corelib/cerror"
	"github.com/chenjie199234/Corelib/stream"
	"github.com/chenjie199234/Corelib/util/common"
)

func (s *CrpcServer) getContext(c context.Context, p *stream.Peer, msg *Msg, handlers []OutsideHandler) *Context {
	ctx, ok := s.ctxpool.Get().(*Context)
	if !ok {
		ctx = &Context{
			Context:  c,
			peer:     p,
			msg:      msg,
			metadata: msg.Metadata,
			handlers: handlers,
			status:   0,
		}
		if msg.Metadata == nil {
			ctx.metadata = make(map[string]string)
		}
		return ctx
	}
	ctx.Context = c
	ctx.peer = p
	ctx.msg = msg
	if msg.Metadata != nil {
		ctx.metadata = msg.Metadata
	}
	ctx.handlers = handlers
	ctx.status = 0
	return ctx
}

func (s *CrpcServer) putContext(ctx *Context) {
	for k := range ctx.metadata {
		delete(ctx.metadata, k)
	}
	s.ctxpool.Put(ctx)
}

type Context struct {
	context.Context
	msg      *Msg
	metadata map[string]string
	peer     *stream.Peer
	handlers []OutsideHandler
	status   int8
}

func (c *Context) run() {
	for _, handler := range c.handlers {
		handler(c)
		if c.status != 0 {
			break
		}
	}
}

// has race
func (c *Context) Abort(e error) {
	c.msg.Error = cerror.ConvertStdError(e)
	if c.msg.Error != nil && (c.msg.Error.Httpcode < 400 || c.msg.Error.Httpcode > 999) {
		panic("[crpc.Context.Abort] httpcode must in [400,999]")
	}
	c.msg.Path = ""
	c.msg.Deadline = 0
	c.msg.Body = nil
	c.msg.Metadata = nil
	c.msg.Tracedata = nil
	c.status = -1
}

// has race
func (c *Context) Write(resp []byte) {
	c.msg.Path = ""
	c.msg.Deadline = 0
	c.msg.Body = resp
	c.msg.Error = nil
	c.msg.Metadata = nil
	c.msg.Tracedata = nil
	c.status = 1
}

func (c *Context) WriteString(resp string) {
	c.Write(common.Str2byte(resp))
}

func (c *Context) GetMethod() string {
	return "CRPC"
}
func (c *Context) GetPath() string {
	return c.msg.Path
}
func (c *Context) GetBody() []byte {
	return c.msg.Body
}

// get the direct peer's addr(maybe a proxy)
func (c *Context) GetRemoteAddr() string {
	return c.peer.GetRemoteAddr()
}

// this function try to return the first caller's ip(mostly time it will be the user's ip)
// if can't get the first caller's ip,try to return the real peer's ip which will not be confused by proxy
// if failed,the direct peer's ip will be returned(maybe a proxy)
func (c *Context) GetClientIp() string {
	return c.metadata["Client-IP"]
}
func (c *Context) GetPeerMaxMsgLen() uint32 {
	return c.peer.GetPeerMaxMsgLen()
}
func (c *Context) GetMetadata() map[string]string {
	return c.metadata
}
