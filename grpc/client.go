package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	cerror "github.com/chenjie199234/Corelib/error"
	"github.com/chenjie199234/Corelib/trace"
	"github.com/chenjie199234/Corelib/util/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	gmetadata "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type ClientConfig struct {
	ConnTimeout   time.Duration
	GlobalTimeout time.Duration //global timeout for every rpc call(including connection establish time)
	HeartPorbe    time.Duration
	SocketRBuf    uint32
	SocketWBuf    uint32
	MaxMsgLen     uint32
	UseTLS        bool     //grpc or grpcs
	SkipVerifyTLS bool     //don't verify the server's cert
	CAs           []string //CAs' path,specific the CAs need to be used,this will overwrite the default behavior:use the system's certpool
	ResolverName  string
	BalancerName  string
}

func (c *ClientConfig) validate() {
	if c.ConnTimeout <= 0 {
		c.ConnTimeout = time.Millisecond * 500
	}
	if c.GlobalTimeout < 0 {
		c.GlobalTimeout = 0
	}
	if c.HeartPorbe < time.Second {
		c.HeartPorbe = 1500 * time.Millisecond
	}
	if c.SocketRBuf == 0 {
		c.SocketRBuf = 1024
	}
	if c.SocketRBuf > 65535 {
		c.SocketRBuf = 65535
	}
	if c.SocketWBuf == 0 {
		c.SocketWBuf = 1024
	}
	if c.SocketWBuf > 65535 {
		c.SocketWBuf = 65535
	}
	if c.MaxMsgLen < 1024 {
		c.MaxMsgLen = 65535
	}
	if c.MaxMsgLen > 65535 {
		c.MaxMsgLen = 65535
	}
}

type GrpcClient struct {
	c           *ClientConfig
	selfappname string
	appname     string
	conn        *grpc.ClientConn
}

func NewGrpcClient(c *ClientConfig, selfgroup, selfname, group, name string) (*GrpcClient, error) {
	if e := common.NameCheck(selfname, false, true, false, true); e != nil {
		return nil, e
	}
	if e := common.NameCheck(name, false, true, false, true); e != nil {
		return nil, e
	}
	if e := common.NameCheck(selfgroup, false, true, false, true); e != nil {
		return nil, e
	}
	if e := common.NameCheck(group, false, true, false, true); e != nil {
		return nil, e
	}
	appname := group + "." + name
	if e := common.NameCheck(appname, true, true, false, true); e != nil {
		return nil, e
	}
	selfappname := selfgroup + "." + selfname
	if e := common.NameCheck(selfappname, true, true, false, true); e != nil {
		return nil, e
	}
	if c == nil {
		c = &ClientConfig{}
	}
	c.validate()
	clientinstance := &GrpcClient{
		c:           c,
		selfappname: selfappname,
		appname:     appname,
	}
	opts := make([]grpc.DialOption, 0)
	if !c.UseTLS {
		opts = append(opts, grpc.WithInsecure())
	} else {
		var certpool *x509.CertPool
		if len(c.CAs) != 0 {
			certpool = x509.NewCertPool()
			for _, cert := range c.CAs {
				certPEM, e := os.ReadFile(cert)
				if e != nil {
					return nil, errors.New("[grpc.client] read cert file:" + cert + " error:" + e.Error())
				}
				if !certpool.AppendCertsFromPEM(certPEM) {
					return nil, errors.New("[grpc.client] load cert file:" + cert + " error:" + e.Error())
				}
			}
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: c.SkipVerifyTLS,
			RootCAs:            certpool,
		})))
	}
	opts = append(opts, grpc.WithReadBufferSize(int(c.SocketRBuf)))
	opts = append(opts, grpc.WithWriteBufferSize(int(c.SocketWBuf)))
	if c.ConnTimeout != 0 {
		dialer := &net.Dialer{Timeout: c.ConnTimeout}
		opts = append(opts, grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return dialer.DialContext(ctx, "tcp", addr)
		}))
	}
	opts = append(opts, grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: c.HeartPorbe, Timeout: c.HeartPorbe*3 + c.HeartPorbe/3}))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(c.MaxMsgLen))))
	opts = append(opts, grpc.WithBalancerName(c.BalancerName))
	conn, e := grpc.Dial(c.ResolverName+":///"+selfname+"."+selfgroup, opts...)
	if e != nil {
		return nil, e
	}
	clientinstance.conn = conn
	return clientinstance, nil
}
func (c *GrpcClient) Call(ctx context.Context, functimeout time.Duration, path string, req interface{}, resp interface{}, metadata map[string]string) error {
	start := time.Now()
	var min time.Duration
	if c.c.GlobalTimeout != 0 {
		min = c.c.GlobalTimeout
	}
	if functimeout != 0 {
		if min == 0 {
			min = functimeout
		} else if functimeout < min {
			min = functimeout
		}
	}
	if min != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, min)
		defer cancel()
	}
	dl, ok := ctx.Deadline()
	if ok && dl.UnixNano() <= start.UnixNano()+int64(5*time.Millisecond) {
		return cerror.ErrDeadlineExceeded
	}
	md := gmetadata.New(nil)
	if len(metadata) != 0 {
		d, _ := json.Marshal(metadata)
		md.Set("core_metadata", common.Byte2str(d))
	}
	traceid, _, _, selfmethod, selfpath := trace.GetTrace(ctx)
	if traceid != "" {
		md.Set("core_tracedata", traceid, c.selfappname, selfmethod, selfpath)
	}
	if md.Len() > 0 {
		ctx = gmetadata.NewOutgoingContext(ctx, md)
	}
	p := &peer.Peer{}
	e := transGrpcError(c.conn.Invoke(ctx, path, req, resp, grpc.Peer(p)))
	end := time.Now()
	trace.Trace(ctx, trace.CLIENT, c.appname, p.Addr.String(), "GRPC", path, &start, &end, e)
	return e
}
func transGrpcError(e error) *cerror.Error {
	if e == nil {
		return nil
	}
	s, _ := status.FromError(e)
	if s == nil {
		return nil
	}
	switch s.Code() {
	case codes.OK:
		return nil
	case codes.Canceled:
		return cerror.ErrCanceled
	case codes.DeadlineExceeded:
		return cerror.ErrDeadlineExceeded
	case codes.Unknown:
		return cerror.ConvertErrorstr(s.Message())
	case codes.Unimplemented:
		return ErrNoapi
	case codes.Unavailable:
		if strings.Contains(s.Message(), "zero addresses") {
			return ErrNoserver
		} else {
			return ErrClosed
		}
	case codes.Unauthenticated:
		return cerror.ErrAuth
	default:
		return cerror.MakeError(int32(s.Code()), http.StatusInternalServerError, s.Message())
	}
}
