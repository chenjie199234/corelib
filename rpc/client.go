package rpc

import (
	"bytes"
	"context"
	"math"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/chenjie199234/Corelib/common"
	cerror "github.com/chenjie199234/Corelib/error"
	"github.com/chenjie199234/Corelib/log"
	"github.com/chenjie199234/Corelib/stream"

	"google.golang.org/protobuf/proto"
)

type PickHandler func(servers []*ServerForPick) *ServerForPick
type DiscoveryHandler func(appname string, client *RpcClient)

//appuniquename = appname:addr
type RpcClient struct {
	c          *stream.InstanceConfig
	timeout    time.Duration
	appname    string
	verifydata []byte
	instance   *stream.Instance

	lker    *sync.RWMutex
	servers []*ServerForPick

	callid  uint64
	reqpool *sync.Pool

	pick      PickHandler
	discovery DiscoveryHandler
}

type ServerForPick struct {
	addr             string
	discoveryservers map[string]struct{} //this app registered on which discovery server
	lker             *sync.Mutex
	peer             *stream.Peer
	starttime        uint64
	status           int //0-idle,1-start,2-verify,3-connected,4-closing

	//active calls
	reqs map[uint64]*req //all reqs to this server

	Pickinfo *pickinfo
}
type pickinfo struct {
	Lastcall                   int64   //last call timestamp nanosecond
	Cpu                        float64 //cpuinfo
	Activecalls                int32   //current active calls
	DiscoveryServers           int32   //this server registered on how many discoveryservers
	DiscoveryServerOfflineTime int64   //
	Addition                   []byte  //addition info register on register center
}

func (s *ServerForPick) Pickable() bool {
	return s.status == 3
}

var lker *sync.Mutex
var all map[string]*RpcClient

func init() {
	rand.Seed(time.Now().UnixNano())
	lker = &sync.Mutex{}
	all = make(map[string]*RpcClient)
}

func NewRpcClient(c *stream.InstanceConfig, globaltimeout time.Duration, appname string, vdata []byte, pick PickHandler, discovery DiscoveryHandler) *RpcClient {
	if e := common.NameCheck(appname, true); e != nil {
		panic("[rpc.client]" + e.Error())
	}
	lker.Lock()
	defer lker.Unlock()
	if client, ok := all[appname]; ok {
		return client
	}
	if pick == nil {
		pick = defaultPicker
	}
	if discovery == nil {
		discovery = defaultdiscovery
	}
	client := &RpcClient{
		timeout:    globaltimeout,
		appname:    appname,
		verifydata: vdata,
		lker:       &sync.RWMutex{},
		servers:    make([]*ServerForPick, 0, 10),
		callid:     0,
		reqpool:    &sync.Pool{},
		pick:       pick,
		discovery:  discovery,
	}
	//tcp instalce
	dupc := *c //duplicate to remove the callback func race
	dupc.Verifyfunc = client.verifyfunc
	dupc.Onlinefunc = client.onlinefunc
	dupc.Userdatafunc = client.userfunc
	dupc.Offlinefunc = client.offlinefunc
	client.c = &dupc
	client.instance = stream.NewInstance(&dupc)
	go discovery(appname, client)
	all[appname] = client
	return client
}

//first key:addr
//second key:discovery server
//value:addition data
func (c *RpcClient) UpdateDiscovery(allapps map[string]map[string]struct{}, addition []byte) {
	//offline app
	c.lker.Lock()
	defer c.lker.Unlock()
	for _, server := range c.servers {
		if _, ok := allapps[server.addr]; !ok {
			//this app unregistered
			server.discoveryservers = nil
			server.Pickinfo.DiscoveryServers = 0
			server.Pickinfo.DiscoveryServerOfflineTime = time.Now().Unix()
		}
	}
	//online app or update app's discoveryservers
	for addr, discoveryservers := range allapps {
		var exist *ServerForPick
		//appuniquename := fmt.Sprintf("%s:%s", c.appname, addr)
		for _, existserver := range c.servers {
			if existserver.addr == addr {
				exist = existserver
				break
			}
		}
		if exist == nil {
			//this is a new register
			c.servers = append(c.servers, &ServerForPick{
				addr:             addr,
				discoveryservers: discoveryservers,
				peer:             nil,
				starttime:        0,
				status:           1,
				reqs:             make(map[uint64]*req, 10),
				lker:             &sync.Mutex{},
				Pickinfo: &pickinfo{
					Lastcall:                   0,
					Cpu:                        1,
					Activecalls:                0,
					DiscoveryServers:           int32(len(discoveryservers)),
					DiscoveryServerOfflineTime: 0,
					Addition:                   addition,
				},
			})
			go c.start(addr)
			continue
		}
		//this is not a new register
		//unregister on which discovery server
		for dserver := range exist.discoveryservers {
			if _, ok := discoveryservers[dserver]; !ok {
				delete(exist.discoveryservers, dserver)
				exist.Pickinfo.DiscoveryServerOfflineTime = time.Now().UnixNano()
			}
		}
		//register on which new discovery server
		for dserver := range discoveryservers {
			if _, ok := exist.discoveryservers[dserver]; !ok {
				exist.discoveryservers[dserver] = struct{}{}
				exist.Pickinfo.DiscoveryServerOfflineTime = 0
			}
		}
		exist.Pickinfo.Addition = addition
		exist.Pickinfo.DiscoveryServers = int32(len(exist.discoveryservers))
	}
}
func (c *RpcClient) start(addr string) {
	tempverifydata := common.Byte2str(c.verifydata) + "|" + c.appname
	if r := c.instance.StartTcpClient(addr, common.Str2byte(tempverifydata)); r == "" {
		c.lker.RLock()
		var exist *ServerForPick
		for _, existserver := range c.servers {
			if existserver.addr == addr {
				exist = existserver
				break
			}
		}
		if exist == nil {
			//app removed
			c.lker.RUnlock()
			return
		}
		if len(exist.discoveryservers) == 0 {
			exist.lker.Lock()
			c.lker.RUnlock()
			exist.status = 0
			exist.lker.Unlock()
			c.unregister(addr)
		} else {
			exist.lker.Lock()
			c.lker.RUnlock()
			exist.status = 1
			exist.lker.Unlock()
			time.Sleep(100 * time.Millisecond)
			go c.start(addr)
		}
	}
}
func (c *RpcClient) unregister(addr string) {
	c.lker.Lock()
	var exist *ServerForPick
	var index int
	for i, existserver := range c.servers {
		if existserver.addr == addr {
			exist = existserver
			index = i
			break
		}
	}
	if exist == nil {
		//already removed
		c.lker.Unlock()
		return
	}
	//check again
	exist.lker.Lock()
	if len(exist.discoveryservers) == 0 && exist.status == 0 {
		//remove app
		c.servers[index], c.servers[len(c.servers)-1] = c.servers[len(c.servers)-1], c.servers[index]
		c.servers = c.servers[:len(c.servers)-1]
		exist.lker.Unlock()
		c.lker.Unlock()
	} else if len(exist.discoveryservers) != 0 && exist.status == 0 {
		exist.status = 1
		exist.lker.Unlock()
		c.lker.Unlock()
		time.Sleep(100 * time.Millisecond)
		go c.start(addr)
	}
}
func (c *RpcClient) verifyfunc(ctx context.Context, appuniquename string, peerVerifyData []byte) ([]byte, bool) {
	if !bytes.Equal(peerVerifyData, c.verifydata) {
		return nil, false
	}
	c.lker.RLock()
	var exist *ServerForPick
	for _, existserver := range c.servers {
		if existserver.addr == appuniquename[strings.Index(appuniquename, ":")+1:] {
			exist = existserver
			break
		}
	}
	if exist == nil {
		//this is impossible
		c.lker.RUnlock()
		return nil, false
	}
	exist.lker.Lock()
	c.lker.RUnlock()
	exist.status = 2
	exist.lker.Unlock()
	return nil, true
}
func (c *RpcClient) onlinefunc(p *stream.Peer, appuniquename string, starttime uint64) {
	c.lker.RLock()
	var exist *ServerForPick
	for _, existserver := range c.servers {
		if existserver.addr == appuniquename[strings.Index(appuniquename, ":")+1:] {
			exist = existserver
			break
		}
	}
	if exist == nil {
		//this is impossible
		p.Close()
		c.lker.RUnlock()
		return
	}
	exist.lker.Lock()
	c.lker.RUnlock()
	exist.peer = p
	exist.starttime = starttime
	exist.status = 3
	p.SetData(unsafe.Pointer(exist))
	log.Info("[rpc.client.onlinefunc] server:", appuniquename, "online")
	exist.lker.Unlock()
}

func (c *RpcClient) userfunc(p *stream.Peer, appuniquename string, data []byte, starttime uint64) {
	server := (*ServerForPick)(p.GetData())
	msg := &Msg{}
	if e := proto.Unmarshal(data, msg); e != nil {
		//this is impossible
		log.Error("[rpc.client.userfunc] server data format error:", e)
		return
	}
	server.lker.Lock()
	e := cerror.ErrorstrToError(msg.Error)
	if e != nil && e.Code == ERRCLOSING {
		server.status = 4
	}
	req, ok := server.reqs[msg.Callid]
	if !ok {
		server.lker.Unlock()
		return
	}
	if msg.Cpu < 1 {
		atomic.StoreUint64((*uint64)(unsafe.Pointer(&server.Pickinfo.Cpu)), 1)
	} else {
		atomic.StoreUint64((*uint64)(unsafe.Pointer(&server.Pickinfo.Cpu)), math.Float64bits(msg.Cpu))
	}
	if req.callid == msg.Callid {
		req.resp = msg.Body
		req.err = e
		req.finish <- struct{}{}
		atomic.AddInt32(&server.Pickinfo.Activecalls, -1)
	}
	server.lker.Unlock()
}
func (c *RpcClient) offlinefunc(p *stream.Peer, appuniquename string, starttime uint64) {
	server := (*ServerForPick)(p.GetData())
	server.lker.Lock()
	server.peer = nil
	server.starttime = 0
	//all req failed
	for _, req := range server.reqs {
		if req.callid != 0 {
			req.resp = nil
			req.err = ERR[ERRCLOSED]
			req.finish <- struct{}{}
		}
	}
	server.reqs = make(map[uint64]*req, 10)
	if len(server.discoveryservers) == 0 {
		server.status = 0
		server.lker.Unlock()
		c.unregister(appuniquename)
	} else {
		server.status = 1
		server.lker.Unlock()
		time.Sleep(100 * time.Millisecond)
		go c.start(appuniquename[strings.Index(appuniquename, ":")+1:])
	}
	log.Info("[Mrpc.client.offlinefunc] server:", appuniquename, "offline")
}

func (c *RpcClient) Call(ctx context.Context, functimeout time.Duration, path string, in []byte) ([]byte, error) {
	var min time.Duration
	if c.timeout < functimeout {
		min = c.timeout
	} else {
		min = functimeout
	}
	if min != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, min)
		defer cancel()
	}
	dl, ok := ctx.Deadline()
	if ok && dl.UnixNano() <= time.Now().UnixNano()+int64(5*time.Millisecond) {
		//ttl + server logic time
		return nil, ERR[ERRCTXTIMEOUT]
	}
	msg := &Msg{
		Callid:   atomic.AddUint64(&c.callid, 1),
		Path:     path,
		Deadline: dl.UnixNano(),
		Body:     in,
		Metadata: GetAllMetadata(ctx),
	}
	d, _ := proto.Marshal(msg)
	if len(d) > c.c.TcpC.MaxMessageLen {
		return nil, ERR[ERRMSGLARGE]
	}
	var server *ServerForPick
	r := c.getreq(msg.Callid)
	for {
		//pick server
		for {
			c.lker.RLock()
			server = c.pick(c.servers)
			if server == nil {
				c.lker.RUnlock()
				c.putreq(r)
				return nil, ERR[ERRNOSERVER]
			}
			server.lker.Lock()
			c.lker.RUnlock()
			if !server.Pickable() {
				server.lker.Unlock()
				continue
			}
			if msg.Deadline != 0 && msg.Deadline <= time.Now().UnixNano()+int64(5*time.Millisecond) {
				server.lker.Unlock()
				return nil, ERR[ERRCTXTIMEOUT]
			}
			if e := server.peer.SendMessage(d, server.starttime, false); e != nil {
				if e == stream.ERRMSGLENGTH {
					server.lker.Unlock()
					return nil, ERR[ERRMSGLARGE]
				}
				if e == stream.ERRCONNCLOSED {
					server.status = 4
				}
				server.lker.Unlock()
				continue
			}
			atomic.StoreInt64(&server.Pickinfo.Lastcall, time.Now().UnixNano())
			server.reqs[msg.Callid] = r
			atomic.AddInt32(&server.Pickinfo.Activecalls, 1)
			server.lker.Unlock()
			break
		}
		select {
		case <-r.finish:
			if r.err != nil && r.err.Code == ERRCLOSING {
				r.resp = nil
				r.err = nil
				continue
			}
			resp := r.resp
			err := r.err
			server.lker.Lock()
			delete(server.reqs, msg.Callid)
			server.lker.Unlock()
			c.putreq(r)
			//resp and err maybe both nil
			return resp, err
		case <-ctx.Done():
			server.lker.Lock()
			delete(server.reqs, msg.Callid)
			server.lker.Unlock()
			c.putreq(r)
			e := ctx.Err()
			if e == context.Canceled {
				return nil, ERR[ERRCTXCANCEL]
			} else if e == context.DeadlineExceeded {
				return nil, ERR[ERRCTXTIMEOUT]
			} else {
				return nil, ERR[ERRUNKNOWN]
			}
		}
	}
}

type req struct {
	callid    uint64
	finish    chan struct{}
	resp      []byte
	err       *cerror.Error
	starttime int64
}

func (r *req) reset() {
	r.callid = 0
	for len(r.finish) > 0 {
		<-r.finish
	}
	r.resp = nil
	r.err = nil
	r.starttime = 0
}
func (c *RpcClient) getreq(callid uint64) *req {
	r, ok := c.reqpool.Get().(*req)
	if ok {
		r.reset()
		r.callid = callid
		r.starttime = time.Now().UnixNano()
		return r
	}
	return &req{
		callid:    callid,
		finish:    make(chan struct{}),
		resp:      nil,
		err:       nil,
		starttime: time.Now().UnixNano(),
	}
}
func (c *RpcClient) putreq(r *req) {
	r.reset()
	c.reqpool.Put(r)
}
