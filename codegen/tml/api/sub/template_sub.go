package sub

import (
	"os"
	"text/template"
)

const txt = `syntax="proto3";

//this is the proto package name,all proto in this project must use this name as the proto package name
package {{.ProjectName}};
//this is the golang's package name,all proto in this project must use this name as the golang's package name
option go_package="{{.PackageName}}/api;api";
//https://github.com/chenjie199234/Corelib/blob/main/pbex/pbex.proto
import "pbex/pbex.proto";

//this is the proto file for {{.Sname}} service
service {{.Sname}}{
	//rpc example(examplereq)returns(exampleresp){
	//	option (pbex.method)="get";//can be set to get,delete,post,put,patch
	//	option (pbex.web_midwares)="b";
	//	option (pbex.web_midwares)="c";
	//	option (pbex.web_midwares)="a";//this function on web protocol has 3 midwares,it's order is b,c,a
	//	option (pbex.crpc_midwares)="b";
	//	option (pbex.crpc_midwares)="c";
	//	option (pbex.crpc_midwares)="a";//this function on crpc protocol has 3 midwares,it's order is b,c,a
	//	option (pbex.cgrpc_midwares)="b";
	//	option (pbex.cgrpc_midwares)="c";
	//	option (pbex.cgrpc_midwares)="a";//this function on grpc protocol has 3 midwares,it's order is b,c,a
	//}
}
//req can be set with pbex extentions
//message examplereq{
	//int64 example_for_extentions=1[(pbex.int_gt)=1,(pbex.int_lt)=100];
//}
//resp's pbex extentions will be ignore
//message exampleresp{
	//int64 example_resp=1;
//}`

type data struct {
	PackageName string
	ProjectName string
	Sname       string
}

func CreatePathAndFile(packagename, projectname, sname string) {
	tmp := &data{
		PackageName: packagename,
		ProjectName: projectname,
		Sname:       sname,
	}
	if e := os.MkdirAll("./api/", 0755); e != nil {
		panic("mkdir ./api/ error: " + e.Error())
	}
	prototemplate, e := template.New("./api/" + sname + ".proto").Parse(txt)
	if e != nil {
		panic("parse ./api/" + sname + ".proto error: " + e.Error())
	}
	file, e := os.OpenFile("./api/"+sname+".proto", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic("open ./api/" + sname + ".proto error: " + e.Error())
	}
	if e := prototemplate.Execute(file, tmp); e != nil {
		panic("write ./api/" + sname + ".proto error: " + e.Error())
	}
	if e := file.Sync(); e != nil {
		panic("sync ./api/" + sname + ".proto error: " + e.Error())
	}
	if e := file.Close(); e != nil {
		panic("close ./api/" + sname + ".proto error: " + e.Error())
	}
}
