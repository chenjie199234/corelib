package stream

import (
	"bufio"
	"context"
	"errors"
	"net"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"

	"github.com/chenjie199234/Corelib/bufpool"
	"github.com/chenjie199234/Corelib/log"
	"github.com/chenjie199234/Corelib/util/common"
)

func (this *Instance) StartTcpServer(listenaddr string) error {
	laddr, e := net.ResolveTCPAddr("tcp", listenaddr)
	if e != nil {
		return errors.New("[Stream.TCP.StartTcpServer] resolve addr:" + listenaddr + " error:" + e.Error())
	}
	this.tcplistener, e = net.ListenTCP(laddr.Network(), laddr)
	if e != nil {
		return errors.New("[Stream.TCP.StartTcpServer] listen addr:" + listenaddr + " error:" + e.Error())
	}
	for {
		p := this.getPeer(TCP, CLIENT, this.conf.TcpC.MaxBufferedWriteMsgNum, this.conf.TcpC.MaxMsgLen, this.selfname)
		conn, e := this.tcplistener.AcceptTCP()
		if e != nil {
			return errors.New("[Stream.TCP.Accept] accept connect error:" + e.Error())
		}
		if atomic.LoadInt32(&this.stop) == 1 {
			conn.Close()
			this.putPeer(p)
			this.tcplistener.Close()
			return errors.New("[Stream.TCP.Accept] server closed")
		}
		p.conn = unsafe.Pointer(conn)
		p.setbuffer(int(this.conf.TcpC.SocketRBufLen), int(this.conf.TcpC.SocketWBufLen))
		if p.reader == nil {
			p.reader = bufio.NewReaderSize(conn, int(this.conf.TcpC.SocketRBufLen))
		} else {
			p.reader.Reset(conn)
		}
		go this.sworker(p, this.conf.TcpC.MaxMsgLen)
	}
}
func (this *Instance) StartUnixServer(listenaddr string) error {
	laddr, e := net.ResolveUnixAddr("unix", listenaddr)
	if e != nil {
		return errors.New("[Stream.UNIX.StartUnixServer] resolve addr:" + listenaddr + " error:" + e.Error())
	}
	this.unixlistener, e = net.ListenUnix("unix", laddr)
	if e != nil {
		return errors.New("[Stream.UNIX.StartUnixServer] listening addr:" + listenaddr + " error:" + e.Error())
	}
	for {
		p := this.getPeer(UNIX, CLIENT, this.conf.UnixC.MaxBufferedWriteMsgNum, this.conf.UnixC.MaxMessageLen, this.selfname)
		conn, e := this.unixlistener.AcceptUnix()
		if e != nil {
			return errors.New("[Stream.UNIX.Accept] accept connect error:" + e.Error())
		}
		if atomic.LoadInt32(&this.stop) == 1 {
			conn.Close()
			this.putPeer(p)
			this.unixlistener.Close()
			return errors.New("[Stream.UNIX.Accept] server closed")
		}
		p.conn = unsafe.Pointer(conn)
		p.setbuffer(int(this.conf.UnixC.SocketReadBufferLen), int(this.conf.UnixC.SocketWriteBufferLen))
		if p.reader == nil {
			p.reader = bufio.NewReaderSize(conn, int(this.conf.UnixC.SocketReadBufferLen))
		} else {
			p.reader.Reset(conn)
		}
		go this.sworker(p, this.conf.UnixC.MaxMessageLen)
	}
}

func (this *Instance) sworker(p *Peer, maxlen uint) {
	//read first verify message from client
	verifydata := this.verifypeer(p, maxlen)
	if p.clientname == "" {
		p.closeconn()
		this.putPeer(p)
		return
	}
	if !this.addPeer(p) {
		switch p.protocoltype {
		case TCP:
			log.Error("[Stream.TCP.sworker] duplicate connect from client:", p.getpeeruniquename())
		case UNIX:
			log.Error("[Stream.UNIX.sworker] duplicate connect from client:", p.getpeeruniquename())
		}
		p.closeconn()
		this.putPeer(p)
		return
	}
	if atomic.LoadInt32(&this.stop) == 1 {
		p.closeconn()
		//after addpeer should use this way to delete this peer
		this.noticech <- p
		return
	}
	//verify client success,send self's verify message to client
	verifymsg := makeVerifyMsg(p.servername, verifydata, p.starttime, true)
	send := 0
	num := 0
	var e error
	for send < verifymsg.Len() {
		switch p.protocoltype {
		case TCP:
			if num, e = (*net.TCPConn)(p.conn).Write(verifymsg.Bytes()[send:]); e != nil {
				log.Error("[Stream.TCP.sworker] write first verify msg to client:", p.getpeeruniquename(), "error:", e)
			}
		case UNIX:
			if num, e = (*net.UnixConn)(p.conn).Write(verifymsg.Bytes()[send:]); e != nil {
				log.Error("[Stream.UNIX.sworker] write first verify msg to client:", p.getpeeruniquename(), "error:", e)
			}
		}
		if e != nil {
			bufpool.PutBuffer(verifymsg)
			p.closeconn()
			this.putPeer(p)
			return
		}
		send += num
	}
	bufpool.PutBuffer(verifymsg)
	if this.conf.Onlinefunc != nil {
		this.conf.Onlinefunc(p, p.getpeeruniquename(), p.starttime)
	}
	go this.read(p, maxlen)
	go this.write(p)
	return
}

// success return peeruniquename
// fail return empty
func (this *Instance) StartTcpClient(serveraddr string, verifydata []byte) string {
	if atomic.LoadInt32(&this.stop) == 1 {
		return ""
	}
	dialer := net.Dialer{
		Timeout: this.conf.TcpC.ConnectTimeout,
		Control: func(network, address string, c syscall.RawConn) error {
			c.Control(func(fd uintptr) {
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF, int(this.conf.TcpC.SocketRBufLen))
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_SNDBUF, int(this.conf.TcpC.SocketWBufLen))
			})
			return nil
		},
	}
	conn, e := dialer.Dial("tcp", serveraddr)
	if e != nil {
		log.Error("[Stream.TCP.StartTcpClient] dial error:", e)
		return ""
	}
	p := this.getPeer(TCP, SERVER, this.conf.TcpC.MaxBufferedWriteMsgNum, this.conf.TcpC.MaxMsgLen, this.selfname)
	p.conn = unsafe.Pointer(conn.(*net.TCPConn))
	p.setbuffer(int(this.conf.TcpC.SocketRBufLen), int(this.conf.TcpC.SocketWBufLen))
	if p.reader == nil {
		p.reader = bufio.NewReaderSize(conn, int(this.conf.TcpC.SocketRBufLen))
	} else {
		p.reader.Reset(conn)
	}
	return this.cworker(p, this.conf.TcpC.MaxMsgLen, verifydata)
}

// success return peeruniquename
// fail return empty
func (this *Instance) StartUnixClient(serveraddr string, verifydata []byte) string {
	if atomic.LoadInt32(&this.stop) == 1 {
		return ""
	}
	dialer := net.Dialer{
		Timeout: this.conf.UnixC.ConnectTimeout,
		Control: func(network, address string, c syscall.RawConn) error {
			c.Control(func(fd uintptr) {
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_RCVBUF, int(this.conf.UnixC.SocketReadBufferLen))
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_SNDBUF, int(this.conf.UnixC.SocketWriteBufferLen))
			})
			return nil
		},
	}
	conn, e := dialer.Dial("unix", serveraddr)
	if e != nil {
		log.Error("[Stream.UNIX.StartUnixClient] dial error:", e)
		return ""
	}
	p := this.getPeer(UNIX, SERVER, this.conf.UnixC.MaxBufferedWriteMsgNum, this.conf.UnixC.MaxMessageLen, this.selfname)
	p.conn = unsafe.Pointer(conn.(*net.UnixConn))
	p.setbuffer(int(this.conf.UnixC.SocketReadBufferLen), int(this.conf.UnixC.SocketWriteBufferLen))
	if p.reader == nil {
		p.reader = bufio.NewReaderSize(conn, int(this.conf.UnixC.SocketReadBufferLen))
	} else {
		p.reader.Reset(conn)
	}
	return this.cworker(p, this.conf.UnixC.MaxMessageLen, verifydata)
}

func (this *Instance) cworker(p *Peer, maxlen uint, verifydata []byte) string {
	//send self's verify message to server
	verifymsg := makeVerifyMsg(p.clientname, verifydata, 0, true)
	send := 0
	num := 0
	var e error
	for send < verifymsg.Len() {
		switch p.protocoltype {
		case TCP:
			if num, e = (*net.TCPConn)(p.conn).Write(verifymsg.Bytes()[send:]); e != nil {
				log.Error("[Stream.TCP.cworker] write first verify msg to server addr:", p.getpeeraddr(), "error:", e)
			}
		case UNIX:
			if num, e = (*net.UnixConn)(p.conn).Write(verifymsg.Bytes()[send:]); e != nil {
				log.Error("[Stream.UNIX.cworker] write first verify msg to server addr:", p.getpeeraddr(), "error:", e)
			}
		}
		if e != nil {
			bufpool.PutBuffer(verifymsg)
			p.closeconn()
			this.putPeer(p)
			return ""
		}
		send += num
	}
	bufpool.PutBuffer(verifymsg)
	//read first verify message from server
	_ = this.verifypeer(p, maxlen)
	if p.servername == "" {
		p.closeconn()
		this.putPeer(p)
		return ""
	}
	//verify server success
	if !this.addPeer(p) {
		switch p.protocoltype {
		case TCP:
			log.Error("[Stream.TCP.cworker] duplicate connect to server:", p.getpeeruniquename())
		case UNIX:
			log.Error("[Stream.UNIX.cworker] duplicate connect to server:", p.getpeeruniquename())
		}
		p.closeconn()
		this.putPeer(p)
		return ""
	}
	if atomic.LoadInt32(&this.stop) == 1 {
		p.closeconn()
		//after addpeer should use this way to delete this peer
		this.noticech <- p
		return ""
	}
	if this.conf.Onlinefunc != nil {
		this.conf.Onlinefunc(p, p.getpeeruniquename(), p.starttime)
	}
	uniquename := p.getpeeruniquename()
	go this.read(p, maxlen)
	go this.write(p)
	return uniquename
}
func (this *Instance) verifypeer(p *Peer, maxlen uint) []byte {
	var ctx context.Context
	var cancel context.CancelFunc
	switch p.protocoltype {
	case TCP:
		(*net.TCPConn)(p.conn).SetReadDeadline(time.Now().Add(this.conf.TcpC.ConnectTimeout))
		ctx, cancel = context.WithTimeout(p, this.conf.TcpC.ConnectTimeout)
	case UNIX:
		(*net.UnixConn)(p.conn).SetReadDeadline(time.Now().Add(this.conf.UnixC.ConnectTimeout))
		ctx, cancel = context.WithTimeout(p, this.conf.UnixC.ConnectTimeout)
	}
	defer cancel()
	data, e := p.readMessage(maxlen)
	if data != nil {
		defer bufpool.PutBuffer(data)
	}
	var sender string
	var peerverifydata []byte
	var starttime uint64
	if e == nil {
		if data == nil {
			e = errors.New("empty message")
		} else {
			var msgtype int
			if msgtype, e = getMsgType(data.Bytes()); e == nil {
				if msgtype != VERIFY {
					e = errors.New("first msg is not verify msg")
				} else if sender, peerverifydata, starttime, e = getVerifyMsg(data.Bytes()); e == nil && (sender == "" || sender == p.getselfname()) {
					e = errors.New("sender name illegal")
				}
			}
		}
	}
	if e != nil {
		switch p.protocoltype {
		case TCP:
			switch p.peertype {
			case CLIENT:
				log.Error("[Stream.TCP.verifypeer] read first verify msg from client addr:", p.getpeeraddr(), "error:", e)
			case SERVER:
				log.Error("[Stream.TCP.verifypeer] read first verify msg from server addr:", p.getpeeraddr(), "error:", e)
			}
		case UNIX:
			switch p.peertype {
			case CLIENT:
				log.Error("[Stream.UNIX.verifypeer] read first verify msg from client addr:", p.getpeeraddr(), "error:", e)
			case SERVER:
				log.Error("[Stream.UNIX.verifypeer] read first verify msg from server addr:", p.getpeeraddr(), "error:", e)
			}
		}
		return nil
	}
	p.lastactive = uint64(time.Now().UnixNano())
	p.recvidlestart = p.lastactive
	p.sendidlestart = p.lastactive
	switch p.peertype {
	case CLIENT:
		dup := make([]byte, 0, len(sender))
		dup = append(dup, sender...)
		p.clientname = common.Byte2str(dup)
		p.starttime = p.lastactive
	case SERVER:
		dup := make([]byte, 0, len(sender))
		dup = append(dup, sender...)
		p.servername = common.Byte2str(dup)
		p.starttime = starttime
	}
	response, success := this.conf.Verifyfunc(ctx, p.getpeeruniquename(), peerverifydata)
	if !success {
		switch p.protocoltype {
		case TCP:
			switch p.peertype {
			case CLIENT:
				log.Error("[Stream.TCP.verifypeer] client:", p.getpeeruniquename(), "verify failed with data:", common.Byte2str(peerverifydata))
			case SERVER:
				log.Error("[Stream.TCP.verifypeer] server:", p.getpeeruniquename(), "verify failed with data:", common.Byte2str(peerverifydata))
			}
		case UNIX:
			switch p.peertype {
			case CLIENT:
				log.Error("[Stream.UNIX.verifypeer] client:", p.getpeeruniquename(), "verify failed with data:", common.Byte2str(peerverifydata))
			case SERVER:
				log.Error("[Stream.UNIX.verifypeer] server:", p.getpeeruniquename(), "verify failed with data:", common.Byte2str(peerverifydata))
			}
		}
		p.clientname = ""
		p.servername = ""
		return nil
	}
	return response
}
func (this *Instance) read(p *Peer, maxlen uint) {
	defer func() {
		//every connection will have two goruntine to work for it
		if old := atomic.SwapUint32(&p.status, 0); old > 0 {
			//cause write goruntine return,this will be useful when there is nothing in writebuffer
			p.closeconn()
			//prevent write goruntine block on read channel
			p.writerbuffer <- (*bufpool.Buffer)(nil)
			p.heartbeatbuffer <- (*bufpool.Buffer)(nil)
		} else {
			if this.conf.Offlinefunc != nil {
				this.conf.Offlinefunc(p, p.getpeeruniquename(), p.starttime)
			}
			//when second goruntine return,put connection back to the pool
			this.noticech <- p
		}
	}()
	//after verify,the read timeout is useless,heartbeat will work for this
	switch p.protocoltype {
	case TCP:
		(*net.TCPConn)(p.conn).SetReadDeadline(time.Time{})
	case UNIX:
		(*net.UnixConn)(p.conn).SetReadDeadline(time.Time{})
	}
	for {
		var msgtype int
		data, e := p.readMessage(maxlen)
		if e == nil {
			if data == nil {
				e = errors.New("empty message")
			} else {
				msgtype, e = getMsgType(data.Bytes())
				if e == nil && msgtype != HEART && msgtype != USER && msgtype != CLOSEREAD && msgtype != CLOSEWRITE {
					e = errors.New("unknown msg type")
				}
			}
		}
		if e != nil {
			switch p.protocoltype {
			case TCP:
				switch p.peertype {
				case CLIENT:
					log.Error("[Stream.TCP.read] read msg from client:", p.getpeeruniquename(), "error:", e)
				case SERVER:
					log.Error("[Stream.TCP.read] read msg from server:", p.getpeeruniquename(), "error:", e)
				}
			case UNIX:
				switch p.peertype {
				case CLIENT:
					log.Error("[Stream.UNIX.read] read msg from client:", p.getpeeruniquename(), "error:", e)
				case SERVER:
					log.Error("[Stream.UNIX.read] read msg from server:", p.getpeeruniquename(), "error:", e)
				}
			}
			if data != nil {
				bufpool.PutBuffer(data)
			}
			return
		}
		//deal message
		switch msgtype {
		case HEART:
			//update lastactive time
			p.lastactive = uint64(time.Now().UnixNano())
		case USER:
			if !p.closewrite {
				userdata, starttime, _ := getUserMsg(data.Bytes())
				if starttime == p.starttime {
					//update lastactive time
					p.lastactive = uint64(time.Now().UnixNano())
					p.recvidlestart = p.lastactive
					this.conf.Userdatafunc(p, p.getpeeruniquename(), userdata, starttime)
				}
				//drop race data
			}
		case CLOSEREAD:
			starttime, _ := getCloseReadMsg(data.Bytes())
			if starttime == p.starttime {
				p.closeread = true
				if p.closewrite {
					return
				}
				p.writerbuffer <- makeCloseWriteMsg(p.starttime, true)
			}
			//drop race data
		case CLOSEWRITE:
			starttime, _ := getCloseWriteMsg(data.Bytes())
			if starttime == p.starttime {
				p.closewrite = true
				if p.closeread {
					return
				}
				p.writerbuffer <- makeCloseReadMsg(p.starttime, true)
			}
			//drop race data
		}
		bufpool.PutBuffer(data)
	}
}

func (this *Instance) write(p *Peer) {
	defer func() {
		//drop all data,prevent close read goruntine block on send empty data to these channel
		for len(p.writerbuffer) > 0 {
			if v := <-p.writerbuffer; v != nil {
				bufpool.PutBuffer(v)
			}
		}
		for len(p.heartbeatbuffer) > 0 {
			if v := <-p.heartbeatbuffer; v != nil {
				bufpool.PutBuffer(v)
			}
		}
		//every connection will have two goruntine to work for it
		if old := atomic.SwapUint32(&p.status, 0); old > 0 {
			//when first goruntine return,close the connection,cause read goruntine return
			p.closeconn()
		} else {
			if this.conf.Offlinefunc != nil {
				this.conf.Offlinefunc(p, p.getpeeruniquename(), p.starttime)
			}
			//when second goruntine return,put connection back to the pool
			this.noticech <- p
		}
	}()
	var data *bufpool.Buffer
	var ok bool
	var send int
	var num int
	var e error
	for {
		status := atomic.LoadUint32(&p.status)
		if (status == 0 || status == 2) && len(p.writerbuffer) == 0 {
			return
		}
		if len(p.heartbeatbuffer) > 0 {
			data, ok = <-p.heartbeatbuffer
		} else {
			select {
			case data, ok = <-p.writerbuffer:
			case data, ok = <-p.heartbeatbuffer:
			}
		}
		if p.closeread && p.closewrite {
			if data != nil {
				bufpool.PutBuffer(data)
			}
			return
		}
		if !ok || data == nil {
			continue
		}
		if !p.closeread {
			p.sendidlestart = uint64(time.Now().UnixNano())
			send = 0
			for send < data.Len() {
				switch p.protocoltype {
				case TCP:
					num, e = (*net.TCPConn)(p.conn).Write(data.Bytes()[send:])
				case UNIX:
					num, e = (*net.UnixConn)(p.conn).Write(data.Bytes()[send:])
				}
				if e != nil {
					switch p.protocoltype {
					case TCP:
						switch p.peertype {
						case CLIENT:
							log.Error("[Stream.TCP.write] write msg to client:", p.getpeeruniquename(), "error:", e)
						case SERVER:
							log.Error("[Stream.TCP.write] write msg to server:", p.getpeeruniquename(), "error:", e)
						}
					case SERVER:
						switch p.peertype {
						case CLIENT:
							log.Error("[Stream.UNIX.write] write msg to client:", p.getpeeruniquename(), "error:", e)
						case SERVER:
							log.Error("[Stream.UNIX.write] write msg to server:", p.getpeeruniquename(), "error:", e)
						}
					}
					return
				}
				send += num
			}
		}
		bufpool.PutBuffer(data)
	}
}
