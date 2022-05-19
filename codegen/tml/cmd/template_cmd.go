package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

const textbash = `#!/bin/bash
#      Warning!!!!!!!!!!!This file is readonly!Don't modify this file!

cd $(dirname $0)

help() {
	echo "cmd.sh — every thing you need"
	echo "         please install git"
	echo "         please install golang"
	echo "         please install protoc           (github.com/protocolbuffers/protobuf)"
	echo "         please install protoc-gen-go    (github.com/protocolbuffers/protobuf-go)"
	echo "         please install codegen          (github.com/chenjie199234/Corelib)"
	echo ""
	echo "Usage:"
	echo "   ./cmd.sh <option>"
	echo ""
	echo "Options:"
	echo "   pb                        Generate the proto in this program"
	echo "   new <sub service name>    Create a new sub service"
	echo "   kube                      Update or add kubernetes config"
	echo "   h/-h/help/-help/--help    Show this message"
}

pb() {
	rm ./api/*.pb.go
	rm ./api/*.md
	go mod tidy
	corelib=$(go list -m -f {{.GoListFormat}} github.com/chenjie199234/Corelib)
	workdir=$(pwd)
	cd $corelib
	go install ./...
	cd $workdir
	protoc -I ./ -I $corelib --go_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I $corelib --go-pbex_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I $corelib --go-cgrpc_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I $corelib --go-crpc_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I $corelib --go-web_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I $corelib --go-markdown_out=paths=source_relative:. ./api/*.proto
	go mod tidy
}

new() {
	codegen -n {{.ProjectName}} -p {{.PackageName}} -s $1
}

kube() {
	codegen -n {{.ProjectName}} -p {{.PackageName}} -k
}

if !(type git >/dev/null 2>&1);then
	echo "missing dependence: git"
	exit 0
fi

if !(type go >/dev/null 2>&1);then
	echo "missing dependence: golang"
	exit 0
fi

if !(type protoc >/dev/null 2>&1);then
	echo "missing dependence: protoc"
	exit 0
fi

if !(type protoc-gen-go >/dev/null 2>&1);then
	echo "missing dependence: protoc-gen-go"
	exit 0
fi

if !(type codegen >/dev/null 2>&1);then
	echo "missing dependence: codegen"
	exit 0
fi

if [[ $# == 0 ]] || [[ "$1" == "h" ]] || [[ "$1" == "help" ]] || [[ "$1" == "-h" ]] || [[ "$1" == "-help" ]] || [[ "$1" == "--help" ]]; then
	help
	exit 0
fi

if [[ "$1" == "pb" ]];then
	pb
	exit 0
fi

if [[ "$1" == "kube" ]];then
	kube
	exit 0
fi

if [[ $# == 2 ]] && [[ "$1" == "new" ]];then
	new $2
	exit 0
fi

echo "option unsupport"
help`
const textbat = `@echo off
REM      Warning!!!!!!!!!!!This file is readonly!Don't modify this file!

cd %~dp0

where /q git.exe
if %errorlevel% == 1 (
	echo "missing dependence: git"
	goto :end
)

where /q go.exe
if %errorlevel% == 1 (
	echo "missing dependence: golang"
	goto :end
)

where /q protoc.exe
if %errorlevel% == 1 (
	echo "missing dependence: protoc"
	goto :end
)

where /q protoc-gen-go.exe
if %errorlevel% == 1 (
	echo "missing dependence: protoc-gen-go"
	goto :end
)

where /q codegen.exe
if %errorlevel% == 1 (
	echo "missing dependence: codegen"
	goto :end
)

if "%1" == "" (
	goto :help
)
if %1 == "" (
	goto :help
)
if %1 == "h" (
	goto :help
)
if "%1" == "h" (
	goto :help
)
if %1 == "-h" (
	goto :help
)
if "%1" == "-h" (
	goto :help
)
if %1 == "help" (
	goto :help
)
if "%1" == "help" (
	goto :help
)
if %1 == "-help" (
	goto :help
)
if "%1" == "-help" (
	goto :help
)
if %1 == "pb" (
	goto :pb
)
if "%1" == "pb" (
	goto :pb
)
if %1 == "kube" (
	goto :kube
)
if "%1" ==  "kube" (
	goto :kube
)
if %1 == "new" (
	if "%2" == "" (
		goto :help
	)
	if %2 == "" (
		goto :help
	)
	goto :new
)
if "%1" == "new" (
	if "%2" == "" (
		goto :help
	)
	if %2 == "" (
		goto :help
	)
	goto :new
)

:pb
	del >nul 2>nul .\api\*.pb.go
	del >nul 2>nul .\api\*.md
	go mod tidy
	for /F %%i in ('go list -m -f {{.GoListFormat}} github.com/chenjie199234/Corelib') do ( set corelib=%%i )
	set workdir=%cd%
	cd %corelib%
	go install ./...
	cd %workdir%
	protoc -I ./ -I %corelib% --go_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I %corelib% --go-pbex_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I %corelib% --go-cgrpc_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I %corelib% --go-crpc_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I %corelib% --go-web_out=paths=source_relative:. ./api/*.proto
	protoc -I ./ -I %corelib% --go-markdown_out=paths=source_relative:. ./api/*.proto
	go mod tidy
goto :end

:kube
	codegen -n {{.ProjectName}} -p {{.PackageName}} -k
goto :end

:new
	codegen -n {{.ProjectName}} -p {{.PackageName}} -s %2
goto :end

:help
	echo cmd.bat — every thing you need
	echo           please install git
	echo           please install golang
	echo           please install protoc           (github.com/protocolbuffers/protobuf)
	echo           please install protoc-gen-go    (github.com/protocolbuffers/protobuf-go)
	echo           please install codegen          (github.com/chenjie199234/Corelib)
	echo
	echo Usage:
	echo    ./cmd.bat <option^>
	echo
	echo Options:
	echo    pb                        Generate the proto in this program.
	echo    new <sub service name^>    Create a new sub service.
	echo    kube                      Update or add kubernetes config.
	echo    h/-h/help/-help/--help    Show this message.

:end
pause
exit /b 0`

const path = "./"
const namebash = "cmd.sh"
const namebat = "cmd.bat"

var tmlbash *template.Template
var tmlbat *template.Template
var filebash *os.File
var filebat *os.File

type Data struct {
	PackageName  string
	ProjectName  string
	GoListFormat string
}

func init() {
	var e error
	tmlbash, e = template.New("bash").Parse(textbash)
	if e != nil {
		panic(fmt.Sprintf("create template error:%s", e))
	}
	tmlbat, e = template.New("bat").Parse(strings.Replace(textbat, "\n", "\r\n", -1))
	if e != nil {
		panic(fmt.Sprintf("create template error:%s", e))
	}
}
func CreatePathAndFile() {
	var e error
	if e = os.MkdirAll(path, 0755); e != nil {
		panic(fmt.Sprintf("make dir:%s error:%s", path, e))
	}
	filebash, e = os.OpenFile(path+namebash, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic(fmt.Sprintf("make file:%s error:%s", path+namebash, e))
	}
	e = os.Chmod(path+namebash, 0755)
	if e != nil {
		panic(fmt.Sprintf("change file:%s execute right error:%s", path+namebash, e))
	}
	filebat, e = os.OpenFile(path+namebat, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic(fmt.Sprintf("make file:%s error:%s", path+namebat, e))
	}
}
func Execute(PackageName, ProjectName string) {
	if e := tmlbash.Execute(filebash, &Data{PackageName: PackageName, ProjectName: ProjectName, GoListFormat: "\"{{.Dir}}\""}); e != nil {
		panic(fmt.Sprintf("write content into file:%s error:%s", path+namebash, e))
	}
	if e := tmlbat.Execute(filebat, &Data{PackageName: PackageName, ProjectName: ProjectName, GoListFormat: "\"{{.Dir}}\""}); e != nil {
		panic(fmt.Sprintf("write content into file:%s error:%s", path+namebat, e))
	}
}
