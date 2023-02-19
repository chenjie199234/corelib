package status

import (
	"os"
	"text/template"
)

const txt = `package status

import (
	"context"
	"time"

	//"{{.}}/config"
	"{{.}}/api"
	statusdao "{{.}}/dao/status"
	//"{{.}}/ecode"

	//"github.com/chenjie199234/Corelib/cgrpc"
	//"github.com/chenjie199234/Corelib/crpc"
	//"github.com/chenjie199234/Corelib/log"
	//"github.com/chenjie199234/Corelib/web"
	"github.com/chenjie199234/Corelib/util/graceful"
)

// Service subservice for status business
type Service struct {
	stop *graceful.Graceful

	statusDao *statusdao.Dao
}

// Start -
func Start() *Service {
	return &Service{
		stop: graceful.New(),

		//statusDao: statusdao.NewDao(config.GetSql("status_sql"), config.GetRedis("status_redis"), config.GetMongo("status_mongo")),
		statusDao: statusdao.NewDao(nil, nil, nil),
	}
}

// Ping -
func (s *Service) Ping(ctx context.Context, in *api.Pingreq) (*api.Pingresp, error) {
	//if _, ok := ctx.(*crpc.Context); ok {
	//        log.Info("this is a crpc call")
	//}
	//if _, ok := ctx.(*cgrpc.Context); ok {
	//        log.Info("this is a cgrpc call")
	//}
	//if _, ok := ctx.(*web.Context); ok {
	//        log.Info("this is a web call")
	//}
	return &api.Pingresp{ClientTimestamp: in.Timestamp, ServerTimestamp: time.Now().UnixNano()}, nil
}

// Stop -
func (s *Service) Stop() {
	s.stop.Close(nil, nil)
}`

func CreatePathAndFile(packagename string) {
	if e := os.MkdirAll("./service/status/", 0755); e != nil {
		panic("mkdir ./service/status/ error: " + e.Error())
	}
	servicetemplate, e := template.New("./service/status/service.go").Parse(txt)
	if e != nil {
		panic("parse ./service/status/service.go template error: " + e.Error())
	}
	file, e := os.OpenFile("./service/status/service.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic("open ./service/status/service.go error: " + e.Error())
	}
	if e := servicetemplate.Execute(file, packagename); e != nil {
		panic("write ./service/status/service.go error: " + e.Error())
	}
	if e := file.Sync(); e != nil {
		panic("sync ./service/status/service.go error: " + e.Error())
	}
	if e := file.Close(); e != nil {
		panic("close ./service/status/service.go error: " + e.Error())
	}
}
