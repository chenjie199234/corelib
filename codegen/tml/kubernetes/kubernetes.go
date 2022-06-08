package kubernetes

import (
	"fmt"
	"os"
	"text/template"
)

const dockerfiletext = `FROM golang:1.18-buster as builder
ENV GOSUMDB='off' \
	GOOS='linux' \
	GOARCH='amd64' \
	CGO_ENABLED=0

RUN mkdir /code
ADD . /code
WORKDIR /code
RUN echo "start build" && go mod tidy && go build -o main && echo "end build"

FROM debian:buster
RUN apt-get update && apt-get install -y ca-certificates curl inetutils-telnet inetutils-ping inetutils-traceroute dnsutils iproute2 procps net-tools neovim && mkdir /root/app
WORKDIR /root/app
EXPOSE 6060 8000 9000 10000
COPY --from=builder /code/main /code/AppConfig.json /code/SourceConfig.json ./
ENTRYPOINT ["./main"]`

const deploymenttext = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.ProjectName}}-deployment
  namespace: <GROUP>
  labels:
    app: {{.ProjectName}}
spec:
  replicas: 2
  revisionHistoryLimit: 5
  minReadySeconds: 5
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  selector:
    matchLabels:
      app: {{.ProjectName}}
  template:
    metadata:
      labels:
        app: {{.ProjectName}}
    spec:
      containers:
        - name: {{.ProjectName}}
          image: <IMAGE>
          imagePullPolicy: IfNotPresent
          resources:
            limits:
              memory: 4096Mi
              cpu: 4000m
            requests:
              memory: 256Mi
              cpu: 250m
          env:
            - name: HOSTIP
              valueFrom:
                fieldRef:
                  fieldPath: status.podIP
            - name: GROUP
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
            - name: LOG_LEVEL
              value: <LOG_LEVEL>
            - name: LOG_TRACE
              value: <LOG_TRACE>
            - name: LOG_TARGET
              value: <LOG_TARGET>
            - name: DEPLOY_ENV
              value: <DEPLOY_ENV>
            - name: RUN_ENV
              value: <RUN_ENV>
            - name: MONITOR
              value: <MONITOR>
            - name: CONFIG_TYPE
              value: <CONFIG_TYPE>
            - name: REMOTE_CONFIG_SERVICE_GROUP
              value: <REMOTE_CONFIG_SERVICE_GROUP>
            - name: REMOTE_CONFIG_SERVICE_HOST
              value: <REMOTE_CONFIG_SERVICE_HOST>
          livenessProbe:
            tcpSocket:
              port: 8000
            initialDelaySeconds: 5
            timeoutSeconds: 1
            periodSeconds: 1
            successThreshold: 1
            failureThreshold: 3
          readinessProbe:
            tcpSocket:
              port: 8000
            initialDelaySeconds: 5
            timeoutSeconds: 1
            periodSeconds: 1
            successThreshold: 1
            failureThreshold: 3
      imagePullSecrets:
        - name: <GROUP>-secret
---
apiVersion: autoscaling/v2beta2
kind: HorizontalPodAutoscaler
metadata:
  name: {{.ProjectName}}-hpa
  namespace: <GROUP>
spec:
  scaleTargetRef:   
    apiVersion: apps/v1
    kind: Deployment  
    name: {{.ProjectName}}-deployment
  maxReplicas: 10
  minReplicas: 2
  metrics:
  - type: Resource
    resource:
      name: memory
      target:
        type: AverageValue
        averageValue: 3500Mi
  - type: Resource
    resource:
      name: cpu
      target:
        type: AverageValue
        averageValue: 3400m{{ if .NeedService }}
---
apiVersion: v1
kind: Service
metadata:
  name: {{.ProjectName}}-headless
  namespace: <GROUP>
  labels:
    app: {{.ProjectName}}
spec:
  type: ClusterIP
  clusterIP: None
  selector:
    app: {{.ProjectName}}
---
apiVersion: v1
kind: Service
metadata:
  name: {{.ProjectName}}
  namespace: <GROUP>
  labels:
    app: {{.ProjectName}}
spec:
  type: ClusterIP
  ports:
  - name: web
    protocol: TCP
    port: 8000
  selector:
    app: {{.ProjectName}}{{ end }}{{ if .NeedIngress}}
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{.ProjectName}}-ingress
  namespace: <GROUP>
  annotations:
    nginx.ingress.kubernetes.io/use-regex: 'true'
spec:
  rules: 
  - host: <HOST>
    http:
      paths:
      - path: /{{.ProjectName}}.*
        backend:
          serviceName: {{.ProjectName}}
          servicePort: 8000{{ end }}`

const path = "./"
const dockerfilename = "Dockerfile"
const deploymentname = "deployment.yaml"

var dockerfiletml *template.Template
var dockerfilefile *os.File

var deploymenttml *template.Template
var deploymentfile *os.File

func init() {
	var e error
	dockerfiletml, e = template.New("dockerfile").Parse(dockerfiletext)
	if e != nil {
		panic(fmt.Sprintf("create template error:%s", e))
	}

	deploymenttml, e = template.New("deployment").Parse(deploymenttext)
	if e != nil {
		panic(fmt.Sprintf("create template error:%s", e))
	}
}
func CreatePathAndFile() {
	var e error
	if e = os.MkdirAll(path, 0755); e != nil {
		panic(fmt.Sprintf("make dir:%s error:%s", path, e))
	}
	dockerfilefile, e = os.OpenFile(path+dockerfilename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic(fmt.Sprintf("make file:%s error:%s", path+dockerfilename, e))
	}

	deploymentfile, e = os.OpenFile(path+deploymentname, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic(fmt.Sprintf("make file:%s error:%s", path+deploymentname, e))
	}
}

type data struct {
	ProjectName string
	NeedService bool
	NeedIngress bool
}

func Execute(projectname string, needservice bool, needingress bool) {
	if e := dockerfiletml.Execute(dockerfilefile, projectname); e != nil {
		panic(fmt.Sprintf("write content into file:%s error:%s", path+dockerfilename, e))
	}
	tempdata := &data{
		ProjectName: projectname,
		NeedService: needservice,
		NeedIngress: needingress,
	}
	if e := deploymenttml.Execute(deploymentfile, tempdata); e != nil {
		panic(fmt.Sprintf("write content into file:%s error:%s", path+deploymentname, e))
	}
}
