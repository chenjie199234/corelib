package discovery

import (
	"fmt"
	"testing"
	"time"

	"github.com/chenjie199234/Corelib/stream"
)

func Test_Client1(t *testing.T) {
	//go func() {
	//        tker := time.NewTicker(time.Second)
	//        for {
	//                <-tker.C
	//                if clientinstance != nil {
	//                        clientinstance.lker.RLock()
	//                        if server, ok := clientinstance.servers["server1:127.0.0.1:9234"]; ok {
	//                                fmt.Println(server.allapps)
	//                        } else {
	//                                fmt.Println("server1: not online")
	//                        }
	//                        if server, ok := clientinstance.servers["server2:127.0.0.1:9235"]; ok {
	//                                fmt.Println(server.allapps)
	//                        } else {
	//                                fmt.Println("server2: not online")
	//                        }
	//                        clientinstance.lker.RUnlock()
	//                }
	//        }
	//}()
	NewDiscoveryClient(&stream.InstanceConfig{
		SelfName:           "client",
		VerifyTimeout:      500,
		HeartbeatTimeout:   5000,
		HeartprobeInterval: 2000,
		GroupNum:           1,
		TcpC: &stream.TcpConfig{
			ConnectTimeout:       500,
			SocketReadBufferLen:  1024,
			SocketWriteBufferLen: 1024,
			AppWriteBufferNum:    256,
		},
	}, []byte{'t', 'e', 's', 't'}, "http://127.0.0.1:8080/discoveryservers")
	time.Sleep(time.Second)
	gch, e := NoticeGrpcChanges("client")
	if e != nil {
		panic("notice grpc change error:" + e.Error())
	}
	hch, e := NoticeHttpChanges("client")
	if e != nil {
		panic("notice http change error:" + e.Error())
	}
	tch, e := NoticeTcpChanges("client")
	if e != nil {
		panic("notice tcp change error:" + e.Error())
	}
	wch, e := NoticeWebsocketChanges("client")
	if e != nil {
		panic("notice websocket change error:" + e.Error())
	}
	go func() {
		for {
			select {
			case <-gch:
				r := GetGrpcInfos("client")
				fmt.Println(r)
			case <-hch:
				r := GetHttpInfos("client")
				fmt.Println(r)
			case <-tch:
				r := GetTcpInfos("client")
				fmt.Println(r)
			case <-wch:
				r := GetWebsocketInfos("client")
				fmt.Println(r)
			}
		}
	}()
	RegisterSelf(&RegMsg{
		GrpcIp:      "",
		GrpcPort:    9000,
		HttpIp:      "",
		HttpPort:    8000,
		TcpIp:       "",
		TcpPort:     7000,
		WebSockIp:   "",
		WebSockPort: 6000,
		WebSockUrl:  "/root",
	})
	select {}
}
func Test_Client2(t *testing.T) {
	//go func() {
	//        tker := time.NewTicker(time.Second)
	//        for {
	//                <-tker.C
	//                if clientinstance != nil {
	//                        clientinstance.lker.RLock()
	//                        if server, ok := clientinstance.servers["server1:127.0.0.1:9234"]; ok {
	//                                fmt.Println(server.allapps)
	//                        } else {
	//                                fmt.Println("server1: not online")
	//                        }
	//                        if server, ok := clientinstance.servers["server2:127.0.0.1:9235"]; ok {
	//                                fmt.Println(server.allapps)
	//                        } else {
	//                                fmt.Println("server2: not online")
	//                        }
	//                        clientinstance.lker.RUnlock()
	//                }
	//        }
	//}()
	NewDiscoveryClient(&stream.InstanceConfig{
		SelfName:           "client",
		VerifyTimeout:      500,
		HeartbeatTimeout:   5000,
		HeartprobeInterval: 2000,
		GroupNum:           1,
		TcpC: &stream.TcpConfig{
			ConnectTimeout:       500,
			SocketReadBufferLen:  1024,
			SocketWriteBufferLen: 1024,
			AppWriteBufferNum:    256,
		},
	}, []byte{'t', 'e', 's', 't'}, "http://127.0.0.1:8080/discoveryservers")
	time.Sleep(time.Second)
	gch, e := NoticeGrpcChanges("client")
	if e != nil {
		panic("notice grpc change error:" + e.Error())
	}
	hch, e := NoticeHttpChanges("client")
	if e != nil {
		panic("notice http change error:" + e.Error())
	}
	tch, e := NoticeTcpChanges("client")
	if e != nil {
		panic("notice tcp change error:" + e.Error())
	}
	wch, e := NoticeWebsocketChanges("client")
	if e != nil {
		panic("notice websocket change error:" + e.Error())
	}
	go func() {
		for {
			select {
			case <-gch:
				r := GetGrpcInfos("client")
				fmt.Println(r)
			case <-hch:
				r := GetHttpInfos("client")
				fmt.Println(r)
			case <-tch:
				r := GetTcpInfos("client")
				fmt.Println(r)
			case <-wch:
				r := GetWebsocketInfos("client")
				fmt.Println(r)
			}
		}
	}()
	RegisterSelf(&RegMsg{
		GrpcIp:      "",
		GrpcPort:    9001,
		HttpIp:      "",
		HttpPort:    8001,
		TcpIp:       "",
		TcpPort:     7001,
		WebSockIp:   "",
		WebSockPort: 6001,
		WebSockUrl:  "/root",
	})
	select {}
}
