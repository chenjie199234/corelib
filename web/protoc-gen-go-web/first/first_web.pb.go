// Code generated by protoc-gen-go-web. DO NOT EDIT.

package first

import (
	second "./second"
	context "context"
	json "encoding/json"
	fmt "fmt"
	bufpool "github.com/chenjie199234/Corelib/bufpool"
	common "github.com/chenjie199234/Corelib/util/common"
	metadata "github.com/chenjie199234/Corelib/util/metadata"
	web "github.com/chenjie199234/Corelib/web"
	io "io"
	http "net/http"
	strconv "strconv"
	time "time"
)

var WebPathFirstHello = "/first.First/Hello"
var WebPathFirstWorld = "/first.First/World"

// FirstWebClient is the client API for First service.
type FirstWebClient interface {
	Hello(context.Context, *Helloreq) (*Helloresp, error)
	World(context.Context, *Worldreq) (*Worldresp, error)
}

type firstWebClient struct {
	cc *web.WebClient
}

//has race,will only return the first's call's client,the config will use the first call's config
func NewFirstWebClient(globaltimeout time.Duration, selfgroup, selfname string, picker web.PickHandler, discover web.DiscoveryHandler) (FirstWebClient, error) {
	cc, e := web.NewWebClient(globaltimeout, selfgroup, selfname, Group, Name, picker, discover)
	if e != nil {
		return nil, e
	}
	return &firstWebClient{cc: cc}, nil
}

func (c *firstWebClient) Hello(ctx context.Context, req *Helloreq) (*Helloresp, error) {
	if req == nil {
		return nil, fmt.Errorf("bad request:nil")
	}
	//gt check
	if float64(req.I32) <= 6 {
		return nil, fmt.Errorf("bad request:gt check failed: i32")
	}
	//not in check
	if vv := strconv.FormatInt(int64(req.I32), 10); vv == "8" ||
		vv == "9" {
		return nil, fmt.Errorf("bad request:notin check failed: i32")
	}
	//empty check
	if len(req.Ri32) == 0 {
		return nil, fmt.Errorf("bad request:empty check failed: ri32")
	}
	//in check
	for _, v := range req.Ri32 {
		if vv := strconv.FormatInt(int64(v), 10); vv != "1" &&
			vv != "2" &&
			vv != "3" {
			return nil, fmt.Errorf("bad request:in check failed: ri32")
		}
	}
	//gt check
	if float64(req.Ui32) <= 6 {
		return nil, fmt.Errorf("bad request:gt check failed: ui32")
	}
	//not in check
	if vv := strconv.FormatUint(uint64(req.Ui32), 10); vv == "8" ||
		vv == "9" {
		return nil, fmt.Errorf("bad request:notin check failed: ui32")
	}
	//empty check
	if len(req.Rui32) == 0 {
		return nil, fmt.Errorf("bad request:empty check failed: rui32")
	}
	//in check
	for _, v := range req.Rui32 {
		if vv := strconv.FormatUint(uint64(v), 10); vv != "1" &&
			vv != "2" &&
			vv != "3" {
			return nil, fmt.Errorf("bad request:in check failed: rui32")
		}
	}
	//empty check
	if len(req.Bs) == 0 {
		return nil, fmt.Errorf("bad request:empty check failed: bs")
	}
	//empty check
	if len(req.Rbs) == 0 {
		return nil, fmt.Errorf("bad request:empty check failed: rbs")
	}
	//empty check
	if len(req.Ss) == 0 {
		return nil, fmt.Errorf("bad request:empty check failed: ss")
	}
	//empty check
	if len(req.Rss) == 0 {
		return nil, fmt.Errorf("bad request:empty check failed: rss")
	}
	//enum check
	if _, ok := second.ABC_name[int32(req.E)]; !ok {
		return nil, fmt.Errorf("bad request:enum check failed: e")
	}
	//enum check
	for _, v := range req.Re {
		if _, ok := second.ABC_name[int32(v)]; !ok {
			return nil, fmt.Errorf("bad request:enum check failed: re")
		}
	}
	if req.A != nil {
		//gt check
		if float64(req.A.Uid) <= 0 {
			return nil, fmt.Errorf("bad request:gt check failed: uid")
		}
	}
	for i := range req.Ra {
		if req.Ra[i] != nil {
			//gt check
			if float64(req.Ra[i].Uid) <= 0 {
				return nil, fmt.Errorf("bad request:gt check failed: uid")
			}
		}
	}
	//enum check
	for _, v := range req.Me {
		if _, ok := second.ABC_name[int32(v)]; !ok {
			return nil, fmt.Errorf("bad request:enum check failed: me")
		}
	}
	//empty check
	if len(req.Me) == 0 {
		return nil, fmt.Errorf("bad request:empty check failed: me")
	}
	//empty check
	if len(req.Mb) == 0 {
		return nil, fmt.Errorf("bad request:empty check failed: mb")
	}
	var header http.Header
	if realcrx, ok := ctx.(*web.Context); ok {
		header = realcrx.GetHeaders()
	}
	if header == nil {
		header = make(http.Header)
	}
	if md := metadata.GetAllMetadata(ctx); len(md) != 0 {
		d, _ := json.Marshal(md)
		header.Set("Metadata", common.Byte2str(d))
	}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	buf := bufpool.GetBuffer()
	if req.I32 != 0 {
		buf.Append("&")
		buf.Append("i32")
		buf.Append("=")
		buf.Append(req.I32)
	}
	if len(req.Ri32) != 0 {
		buf.Append("&")
		buf.Append("ri32")
		buf.Append("=")
		buf.Append(req.Ri32)
	}
	if req.Ui32 != 0 {
		buf.Append("&")
		buf.Append("ui32")
		buf.Append("=")
		buf.Append(req.Ui32)
	}
	if len(req.Rui32) != 0 {
		buf.Append("&")
		buf.Append("rui32")
		buf.Append("=")
		buf.Append(req.Rui32)
	}
	if req.I64 != 0 {
		buf.Append("&")
		buf.Append("i64")
		buf.Append("=")
		buf.Append(req.I64)
	}
	if len(req.Ri64) != 0 {
		buf.Append("&")
		buf.Append("ri64")
		buf.Append("=")
		buf.Append(req.Ri64)
	}
	if req.Ui64 != 0 {
		buf.Append("&")
		buf.Append("ui64")
		buf.Append("=")
		buf.Append(req.Ui64)
	}
	if len(req.Rui64) != 0 {
		buf.Append("&")
		buf.Append("rui64")
		buf.Append("=")
		buf.Append(req.Rui64)
	}
	if len(req.Bs) != 0 {
		buf.Append("&")
		buf.Append("bs")
		buf.Append("=\"")
		buf.Append(common.Byte2str(req.Bs))
		buf.Append("\"")
	}
	if len(req.Rbs) != 0 {
		buf.Append("&")
		buf.Append("rbs")
		buf.Append("=[")
		for _, v := range req.Rbs {
			buf.Append("\"")
			buf.Append(common.Byte2str(v))
			buf.Append("\",")
		}
		buf.Bytes()[buf.Len()-1] = ']'
	}
	if len(req.Ss) != 0 {
		buf.Append("&")
		buf.Append("ss")
		buf.Append("=\"")
		buf.Append(req.Ss)
		buf.Append("\"")
	}
	if len(req.Rss) != 0 {
		buf.Append("&")
		buf.Append("rss")
		buf.Append("=[")
		for _, v := range req.Rss {
			buf.Append("\"")
			buf.Append(v)
			buf.Append("\",")
		}
		buf.Bytes()[buf.Len()-1] = ']'
	}
	if req.F != 0 {
		buf.Append("&")
		buf.Append("f")
		buf.Append("=")
		buf.Append(req.F)
	}
	if len(req.Rf) != 0 {
		buf.Append("&")
		buf.Append("rf")
		buf.Append("=")
		buf.Append(req.Rf)
	}
	if req.E != 0 {
		buf.Append("&")
		buf.Append("e")
		buf.Append("=")
		buf.Append(req.E)
	}
	if len(req.Re) != 0 {
		buf.Append("&")
		buf.Append("re")
		buf.Append("=")
		buf.Append(req.Re)
	}
	if req.A != nil {
		buf.Append("&")
		buf.Append("a")
		buf.Append("=")
		d, _ := json.Marshal(req.A)
		buf.Append(common.Byte2str(d))
	}
	if len(req.Ra) != 0 {
		buf.Append("&")
		buf.Append("ra")
		buf.Append("=")
		d, _ := json.Marshal(req.Ra)
		buf.Append(common.Byte2str(d))
	}
	if len(req.Me) != 0 {
		buf.Append("&")
		buf.Append("me")
		buf.Append("=")
		d, _ := json.Marshal(req.Me)
		buf.Append(common.Byte2str(d))
	}
	if len(req.Mb) != 0 {
		buf.Append("&")
		buf.Append("mb")
		buf.Append("=")
		d, _ := json.Marshal(req.Mb)
		buf.Append(common.Byte2str(d))
	}
	if buf.Len() > 0 {
		buf.Bytes()[0] = '?'
	}
	callback, e := c.cc.Get(ctx, 200000000, WebPathFirstHello+buf.String(), header)
	bufpool.PutBuffer(buf)
	if e != nil {
		return nil, fmt.Errorf("call error:" + e.Error())
	}
	defer callback.Body.Close()
	data, e := io.ReadAll(callback.Body)
	if e != nil {
		return nil, fmt.Errorf("read response error:" + e.Error())
	}
	if callback.StatusCode/100 == 5 || callback.StatusCode/100 == 4 {
		return nil, fmt.Errorf(common.Byte2str(data))
	}
	resp := new(Helloresp)
	if len(data) > 0 {
		if e = json.Unmarshal(data, resp); e != nil {
			return nil, fmt.Errorf("response data format errors" + e.Error())
		}
	}
	return resp, nil
}
func (c *firstWebClient) World(ctx context.Context, req *Worldreq) (*Worldresp, error) {
	if req == nil {
		return nil, fmt.Errorf("bad request:nil")
	}
	var header http.Header
	if realcrx, ok := ctx.(*web.Context); ok {
		header = realcrx.GetHeaders()
	}
	if header == nil {
		header = make(http.Header)
	}
	if md := metadata.GetAllMetadata(ctx); len(md) != 0 {
		d, _ := json.Marshal(md)
		header.Set("Metadata", common.Byte2str(d))
	}
	header.Set("Content-Type", "application/x-www-form-urlencoded")
	buf := bufpool.GetBuffer()
	if buf.Len() > 0 {
		buf.Bytes()[0] = '?'
	}
	callback, e := c.cc.Get(ctx, 0, WebPathFirstWorld+buf.String(), header)
	bufpool.PutBuffer(buf)
	if e != nil {
		return nil, fmt.Errorf("call error:" + e.Error())
	}
	defer callback.Body.Close()
	data, e := io.ReadAll(callback.Body)
	if e != nil {
		return nil, fmt.Errorf("read response error:" + e.Error())
	}
	if callback.StatusCode/100 == 5 || callback.StatusCode/100 == 4 {
		return nil, fmt.Errorf(common.Byte2str(data))
	}
	resp := new(Worldresp)
	if len(data) > 0 {
		if e = json.Unmarshal(data, resp); e != nil {
			return nil, fmt.Errorf("response data format errors" + e.Error())
		}
	}
	return resp, nil
}

// FirstWebServer is the server API for First service.
type FirstWebServer interface {
	Hello(context.Context, *Helloreq) (*Helloresp, error)
	World(context.Context, *Worldreq) (*Worldresp, error)
}

func _First_Hello_WebHandler(handler func(context.Context, *Helloreq) (*Helloresp, error)) web.OutsideHandler {
	return func(ctx *web.Context) {
		req := new(Helloreq)
		if ctx.GetMethod() != http.MethodGet && ctx.GetContentType() == "application/json" {
			data, e := ctx.GetBody()
			if e != nil {
				ctx.WriteString(http.StatusInternalServerError, "server error:read request body error:"+e.Error())
				return
			}
			if len(data) != 0 {
				if e := json.Unmarshal(data, req); e != nil {
					ctx.WriteString(http.StatusBadRequest, "bad request:json format error:"+e.Error())
					return
				}
			}
		} else {
			if e := ctx.ParseForm(); e != nil {
				ctx.WriteString(http.StatusBadRequest, "bad request:form format error:"+e.Error())
				return
			}
			buf := bufpool.GetBuffer()
			buf.Append("{")
			hasfields := false
			if temp := ctx.GetForm("i32"); len(temp) != 0 {
				buf.Append("\"i32\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("ri32"); len(temp) != 0 {
				buf.Append("\"ri32\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("ui32"); len(temp) != 0 {
				buf.Append("\"ui32\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("rui32"); len(temp) != 0 {
				buf.Append("\"rui32\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("i64"); len(temp) != 0 {
				buf.Append("\"i64\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("ri64"); len(temp) != 0 {
				buf.Append("\"ri64\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("ui64"); len(temp) != 0 {
				buf.Append("\"ui64\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("rui64"); len(temp) != 0 {
				buf.Append("\"rui64\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("bs"); len(temp) != 0 {
				buf.Append("\"bs\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("rbs"); len(temp) != 0 {
				buf.Append("\"rbs\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("ss"); len(temp) != 0 {
				buf.Append("\"ss\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("rss"); len(temp) != 0 {
				buf.Append("\"rss\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("f"); len(temp) != 0 {
				buf.Append("\"f\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("rf"); len(temp) != 0 {
				buf.Append("\"rf\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("e"); len(temp) != 0 {
				buf.Append("\"e\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("re"); len(temp) != 0 {
				buf.Append("\"re\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("a"); len(temp) != 0 {
				buf.Append("\"a\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("ra"); len(temp) != 0 {
				buf.Append("\"ra\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("me"); len(temp) != 0 {
				buf.Append("\"me\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if temp := ctx.GetForm("mb"); len(temp) != 0 {
				buf.Append("\"mb\":")
				buf.Append(temp)
				buf.Append(",")
				hasfields = true
			}
			if hasfields {
				buf.Bytes()[buf.Len()-1] = '}'
			} else {
				buf.Append("}")
			}
			if buf.Len() > 2 {
				if e := json.Unmarshal(buf.Bytes(), req); e != nil {
					ctx.WriteString(http.StatusBadRequest, "bad request:form format error:"+e.Error())
					return
				}
			}
			bufpool.PutBuffer(buf)
		}
		//header
		if temp := ctx.GetHeader("i32"); len(temp) != 0 {
			tempnum, e := strconv.ParseInt(temp, 10, 32)
			if e != nil {
				ctx.WriteString(http.StatusBadRequest, "bad request:header: i32 format error")
				return
			}
			req.I32 = int32(tempnum)
		}
		//gt check
		if float64(req.I32) <= 6 {
			ctx.WriteString(http.StatusBadRequest, "bad request:gt check failed: i32")
			return
		}
		//not in check
		if vv := strconv.FormatInt(int64(req.I32), 10); vv == "8" ||
			vv == "9" {
			ctx.WriteString(http.StatusBadRequest, "bad request:notin check failed: i32")
			return
		}
		//empty check
		if len(req.Ri32) == 0 {
			ctx.WriteString(http.StatusBadRequest, "bad request:empty check failed: ri32")
			return
		}
		//in check
		for _, v := range req.Ri32 {
			if vv := strconv.FormatInt(int64(v), 10); vv != "1" &&
				vv != "2" &&
				vv != "3" {
				ctx.WriteString(http.StatusBadRequest, "bad request:in check failed: ri32")
				return
			}
		}
		//header
		if temp := ctx.GetHeader("ui32"); len(temp) != 0 {
			tempnum, e := strconv.ParseUint(temp, 10, 32)
			if e != nil {
				ctx.WriteString(http.StatusBadRequest, "bad request:header: ui32 format error")
				return
			}
			req.Ui32 = uint32(tempnum)
		}
		//gt check
		if float64(req.Ui32) <= 6 {
			ctx.WriteString(http.StatusBadRequest, "bad request:gt check failed: ui32")
			return
		}
		//not in check
		if vv := strconv.FormatUint(uint64(req.Ui32), 10); vv == "8" ||
			vv == "9" {
			ctx.WriteString(http.StatusBadRequest, "bad request:notin check failed: ui32")
			return
		}
		//empty check
		if len(req.Rui32) == 0 {
			ctx.WriteString(http.StatusBadRequest, "bad request:empty check failed: rui32")
			return
		}
		//in check
		for _, v := range req.Rui32 {
			if vv := strconv.FormatUint(uint64(v), 10); vv != "1" &&
				vv != "2" &&
				vv != "3" {
				ctx.WriteString(http.StatusBadRequest, "bad request:in check failed: rui32")
				return
			}
		}
		//empty check
		if len(req.Bs) == 0 {
			ctx.WriteString(http.StatusBadRequest, "bad request:empty check failed: bs")
			return
		}
		//empty check
		if len(req.Rbs) == 0 {
			ctx.WriteString(http.StatusBadRequest, "bad request:empty check failed: rbs")
			return
		}
		//header
		if temp := ctx.GetHeader("ss"); len(temp) != 0 {
			req.Ss = temp
		}
		//empty check
		if len(req.Ss) == 0 {
			ctx.WriteString(http.StatusBadRequest, "bad request:empty check failed: ss")
			return
		}
		//header
		if temp := ctx.GetHeader("rss"); len(temp) != 0 {
			req.Rss = make([]string, 0)
			if e := json.Unmarshal([]byte(temp), &req.Rss); e != nil {
				ctx.WriteString(http.StatusBadRequest, "bad request:header: rss format error")
				return
			}
		}
		//empty check
		if len(req.Rss) == 0 {
			ctx.WriteString(http.StatusBadRequest, "bad request:empty check failed: rss")
			return
		}
		//enum check
		if _, ok := second.ABC_name[int32(req.E)]; !ok {
			ctx.WriteString(http.StatusBadRequest, "bad request:enum check failed: e")
			return
		}
		//enum check
		for _, v := range req.Re {
			if _, ok := second.ABC_name[int32(v)]; !ok {
				ctx.WriteString(http.StatusBadRequest, "bad request:enum check failed: re")
				return
			}
		}
		if req.A != nil {
			//header
			if temp := ctx.GetHeader("uid"); len(temp) != 0 {
				tempnum, e := strconv.ParseInt(temp, 10, 64)
				if e != nil {
					ctx.WriteString(http.StatusBadRequest, "bad request:header: uid format error")
					return
				}
				req.A.Uid = tempnum
			}
			//gt check
			if float64(req.A.Uid) <= 0 {
				ctx.WriteString(http.StatusBadRequest, "bad request:gt check failed: uid")
				return
			}
		}
		for i := range req.Ra {
			if req.Ra[i] != nil {
				//header
				if temp := ctx.GetHeader("uid"); len(temp) != 0 {
					tempnum, e := strconv.ParseInt(temp, 10, 64)
					if e != nil {
						ctx.WriteString(http.StatusBadRequest, "bad request:header: uid format error")
						return
					}
					req.Ra[i].Uid = tempnum
				}
				//gt check
				if float64(req.Ra[i].Uid) <= 0 {
					ctx.WriteString(http.StatusBadRequest, "bad request:gt check failed: uid")
					return
				}
			}
		}
		//enum check
		for _, v := range req.Me {
			if _, ok := second.ABC_name[int32(v)]; !ok {
				ctx.WriteString(http.StatusBadRequest, "bad request:enum check failed: me")
				return
			}
		}
		//empty check
		if len(req.Me) == 0 {
			ctx.WriteString(http.StatusBadRequest, "bad request:empty check failed: me")
			return
		}
		//empty check
		if len(req.Mb) == 0 {
			ctx.WriteString(http.StatusBadRequest, "bad request:empty check failed: mb")
			return
		}
		resp, e := handler(ctx, req)
		if e != nil {
			ctx.WriteString(http.StatusInternalServerError, e.Error())
		} else if resp == nil {
			ctx.WriteString(http.StatusOK, "{}")
		} else {
			respd, _ := json.Marshal(resp)
			ctx.Write(http.StatusOK, respd)
		}
	}
}
func _First_World_WebHandler(handler func(context.Context, *Worldreq) (*Worldresp, error)) web.OutsideHandler {
	return func(ctx *web.Context) {
		req := new(Worldreq)
		if ctx.GetMethod() != http.MethodGet && ctx.GetContentType() == "application/json" {
			data, e := ctx.GetBody()
			if e != nil {
				ctx.WriteString(http.StatusInternalServerError, "server error:read request body error:"+e.Error())
				return
			}
			if len(data) != 0 {
				if e := json.Unmarshal(data, req); e != nil {
					ctx.WriteString(http.StatusBadRequest, "bad request:json format error:"+e.Error())
					return
				}
			}
		} else {
			if e := ctx.ParseForm(); e != nil {
				ctx.WriteString(http.StatusBadRequest, "bad request:form format error:"+e.Error())
				return
			}
			buf := bufpool.GetBuffer()
			buf.Append("{")
			hasfields := false
			if hasfields {
				buf.Bytes()[buf.Len()-1] = '}'
			} else {
				buf.Append("}")
			}
			if buf.Len() > 2 {
				if e := json.Unmarshal(buf.Bytes(), req); e != nil {
					ctx.WriteString(http.StatusBadRequest, "bad request:form format error:"+e.Error())
					return
				}
			}
			bufpool.PutBuffer(buf)
		}
		resp, e := handler(ctx, req)
		if e != nil {
			ctx.WriteString(http.StatusInternalServerError, e.Error())
		} else if resp == nil {
			ctx.WriteString(http.StatusOK, "{}")
		} else {
			respd, _ := json.Marshal(resp)
			ctx.Write(http.StatusOK, respd)
		}
	}
}
func RegisterFirstWebServer(engine *web.WebServer, svc FirstWebServer, allmids map[string]web.OutsideHandler) error {
	//avoid lint
	_ = allmids
	{
		requiredMids := []string{"auth"}
		mids := make([]web.OutsideHandler, 0)
		for _, v := range requiredMids {
			if mid, ok := allmids[v]; ok {
				mids = append(mids, mid)
			}
		}
		mids = append(mids, _First_Hello_WebHandler(svc.Hello))
		if e := engine.Get(WebPathFirstHello, 200000000, mids...); e != nil {
			return e
		}
	}
	if e := engine.Get(WebPathFirstWorld, 0, _First_World_WebHandler(svc.World)); e != nil {
		return e
	}
	return nil
}
