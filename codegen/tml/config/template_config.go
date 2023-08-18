package config

import (
	"os"
	"strings"
	"text/template"
)

const txt = `package config

import (
	"os"
	"strconv"
	"time"

	"{{.}}/model"

	configsdk "github.com/chenjie199234/admin/sdk/config"
	"github.com/chenjie199234/Corelib/log"
)

// EnvConfig can't hot update,all these data is from system env setting
// nil field means that system env not exist
type EnvConfig struct {
	ConfigType *int
	RunEnv     *string
	DeployEnv  *string
}

// EC -
var EC *EnvConfig

// RemoteConfigSdk -
var RemoteConfigSdk *configsdk.Sdk

// notice is a sync function
// don't write block logic inside it
func Init(notice func(c *AppConfig)) {
	initenv()
	if EC.ConfigType != nil && *EC.ConfigType == 1 {
		tmer := time.NewTimer(time.Second * 2)
		waitapp := make(chan *struct{})
		waitsource := make(chan *struct{})
		initremoteapp(notice, waitapp)
		stopwatchsource := initremotesource(waitsource)
		appinit := false
		sourceinit := false
		for {
			select {
			case <-waitapp:
				appinit = true
			case <-waitsource:
				sourceinit = true
				stopwatchsource()
			case <-tmer.C:
				log.Error(nil, "[config.Init] timeout", nil)
				Close()
				os.Exit(1)
			}
			if appinit && sourceinit {
				break
			}
		}
	} else {
		initlocalapp(notice)
		initlocalsource()
	}
}

// Close -
func Close() {
	log.Close()
}

func initenv() {
	EC = &EnvConfig{}
	if str, ok := os.LookupEnv("CONFIG_TYPE"); ok && str != "<CONFIG_TYPE>" && str != "" {
		configtype, e := strconv.Atoi(str)
		if e != nil || (configtype != 0 && configtype != 1 && configtype != 2) {
			log.Error(nil, "[config.initenv] env CONFIG_TYPE must be number in [0,1,2]", nil)
			Close()
			os.Exit(1)
		}
		EC.ConfigType = &configtype
	} else {
		log.Warning(nil, "[config.initenv] missing env CONFIG_TYPE", nil)
	}
	if EC.ConfigType != nil && *EC.ConfigType == 1 {
		var e error
		if RemoteConfigSdk, e = configsdk.NewConfigSdk(model.Group, model.Name, nil); e != nil {
			log.Error(nil, "[config.initenv] new remote config sdk failed", map[string]interface{}{"error": e})
			Close()
			os.Exit(1)
		}
	}
	if str, ok := os.LookupEnv("RUN_ENV"); ok && str != "<RUN_ENV>" && str != "" {
		EC.RunEnv = &str
	} else {
		log.Warning(nil, "[config.initenv] missing env RUN_ENV", nil)
	}
	if str, ok := os.LookupEnv("DEPLOY_ENV"); ok && str != "<DEPLOY_ENV>" && str != "" {
		EC.DeployEnv = &str
	} else {
		log.Warning(nil, "[config.initenv] missing env DEPLOY_ENV", nil)
	}
}`
const apptxt = `package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/chenjie199234/Corelib/log"
	publicmids "github.com/chenjie199234/Corelib/mids"
	"github.com/chenjie199234/Corelib/util/common"
	"github.com/chenjie199234/Corelib/util/ctime"
	"github.com/fsnotify/fsnotify"
)

// AppConfig can hot update
// this is the config used for this app
type AppConfig struct {
	HandlerTimeout     map[string]map[string]ctime.Duration      $json:"handler_timeout"$      //first key path,second key method(GET,POST,PUT,PATCH,DELETE,CRPC,GRPC),value timeout
	WebPathRewrite     map[string]map[string]string              $json:"web_path_rewrite"$     //first key method(GET,POST,PUT,PATCH,DELETE),second key origin url,value new url
	HandlerRate        map[string][]*publicmids.PathRateConfig   $json:"handler_rate"$         //key path
	Accesses           map[string][]*publicmids.PathAccessConfig $json:"accesses"$             //key path
	TokenSecret        string                                    $json:"token_secret"$         //if don't need token check,this can be ingored
	SessionTokenExpire ctime.Duration                            $json:"session_token_expire"$ //if don't need session and token check,this can be ignored
	Service            *ServiceConfig                            $json:"service"$
}
type ServiceConfig struct {
	//add your config here
}

// every time update AppConfig will call this function
func validateAppConfig(ac *AppConfig) {
}

// AC -
var AC *AppConfig

var watcher *fsnotify.Watcher

func initlocalapp(notice func(*AppConfig)) {
	data, e := os.ReadFile("./AppConfig.json")
	if e != nil {
		log.Error(nil, "[config.local.app] read config file failed", map[string]interface{}{"error": e})
		Close()
		os.Exit(1)
	}
	AC = &AppConfig{}
	if e = json.Unmarshal(data, AC); e != nil {
		log.Error(nil, "[config.local.app] config file format wrong", map[string]interface{}{"error": e})
		Close()
		os.Exit(1)
	}
	validateAppConfig(AC)
	log.Info(nil, "[config.local.app] update success", map[string]interface{}{"config": AC})
	if notice != nil {
		notice(AC)
	}
	watcher, e = fsnotify.NewWatcher()
	if e != nil {
		log.Error(nil, "[config.local.app] create watcher for hot update failed", map[string]interface{}{"error": e})
		Close()
		os.Exit(1)
	}
	if e = watcher.Add("./"); e != nil {
		log.Error(nil, "[config.local.app] create watcher for hot update failed", map[string]interface{}{"error": e})
		Close()
		os.Exit(1)
	}
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if filepath.Base(event.Name) != "AppConfig.json" || (!event.Has(fsnotify.Create) && !event.Has(fsnotify.Write)) {
					continue
				}
				data, e := os.ReadFile("./AppConfig.json")
				if e != nil {
					log.Error(nil, "[config.local.app] hot update read config file failed", map[string]interface{}{"error": e})
					continue
				}
				c := &AppConfig{}
				if e = json.Unmarshal(data, c); e != nil {
					log.Error(nil, "[config.local.app] hot update config file format wrong", map[string]interface{}{"error": e})
					continue
				}
				validateAppConfig(c)
				log.Info(nil, "[config.local.app] update success", map[string]interface{}{"config": c})
				if notice != nil {
					notice(c)
				}
				AC = c
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Error(nil, "[config.local.app] hot update watcher failed", map[string]interface{}{"error": err})
			}
		}
	}()
}
func initremoteapp(notice func(*AppConfig), wait chan *struct{}) (stopwatch func()) {
	return RemoteConfigSdk.Watch("AppConfig", func(key, keyvalue, keytype string) {
		//only support json
		if keytype != "json" {
			log.Error(nil, "[config.remote.app] config data can only support json format", nil)
			return
		}
		c := &AppConfig{}
		if e := json.Unmarshal(common.Str2byte(keyvalue), c); e != nil {
			log.Error(nil, "[config.remote.app] config data format wrong", map[string]interface{}{"error": e})
			return
		}
		validateAppConfig(c)
		log.Info(nil, "[config.remote.app] update success", map[string]interface{}{"config": c})
		if notice != nil {
			notice(c)
		}
		AC = c
		select {
		case wait <- nil:
		default:
		}
	})
}`
const sourcetxt = `package config

import (
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"os"
	"time"

	"github.com/chenjie199234/Corelib/log"
	"github.com/chenjie199234/Corelib/redis"
	"github.com/chenjie199234/Corelib/util/common"
	"github.com/chenjie199234/Corelib/util/ctime"
	"github.com/go-sql-driver/mysql"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// sourceConfig can't hot update
type sourceConfig struct {
	CGrpcServer *CGrpcServerConfig      $json:"cgrpc_server"$
	CGrpcClient *CGrpcClientConfig      $json:"cgrpc_client"$
	CrpcServer  *CrpcServerConfig       $json:"crpc_server"$
	CrpcClient  *CrpcClientConfig       $json:"crpc_client"$
	WebServer   *WebServerConfig        $json:"web_server"$
	WebClient   *WebClientConfig        $json:"web_client"$
	Mongo       map[string]*MongoConfig $json:"mongo"$ //key example:xxx_mongo
	Sql         map[string]*SqlConfig   $json:"sql"$   //key example:xx_sql
	Redis       map[string]*RedisConfig $json:"redis"$ //key example:xx_redis
	KafkaPub    []*KafkaPubConfig       $json:"kafka_pub"$
	KafkaSub    []*KafkaSubConfig       $json:"kafka_sub"$
}

// CGrpcServerConfig
type CGrpcServerConfig struct {
	ConnectTimeout ctime.Duration    $json:"connect_timeout"$ //default 500ms,max time to finish the handshake
	GlobalTimeout  ctime.Duration    $json:"global_timeout"$  //default 500ms,max time to handle the request,unless the specific handle timeout is used in HandlerTimeout in AppConfig,handler's timeout will also be effected by caller's deadline
	HeartProbe     ctime.Duration    $json:"heart_probe"$     //default 1.5s
	Certs          map[string]string $json:"certs"$           //key cert path,value private key path,if this is not empty,tls will be used
}

// CGrpcClientConfig
type CGrpcClientConfig struct {
	ConnectTimeout ctime.Duration $json:"connect_timeout"$   //default 500ms,max time to finish the handshake
	GlobalTimeout  ctime.Duration $json:"global_timeout"$ //max time to handle the request,0 means no default timeout
	HeartProbe     ctime.Duration $json:"heart_probe"$    //default 1.5s
}

// CrpcServerConfig -
type CrpcServerConfig struct {
	ConnectTimeout ctime.Duration    $json:"connect_timeout"$ //default 500ms,max time to finish the handshake
	GlobalTimeout  ctime.Duration    $json:"global_timeout"$  //default 500ms,max time to handle the request,unless the specific handle timeout is used in HandlerTimeout in AppConfig,handler's timeout will also be effected by caller's deadline
	HeartProbe     ctime.Duration    $json:"heart_probe"$     //default 1.5s
	Certs          map[string]string $json:"certs"$           //key cert path,value private key path,if this is not empty,tls will be used
}

// CrpcClientConfig -
type CrpcClientConfig struct {
	ConnectTimeout ctime.Duration $json:"connect_timeout"$   //default 500ms,max time to finish the handshake
	GlobalTimeout  ctime.Duration $json:"global_timeout"$ //max time to handle the request,0 means no default timeout
	HeartProbe     ctime.Duration $json:"heart_probe"$    //default 1.5s
}

// WebServerConfig -
type WebServerConfig struct {
	CloseMode      int               $json:"close_mode"$
	ConnectTimeout ctime.Duration    $json:"connect_timeout"$ //default 500ms,max time to finish the handshake and read each whole request
	GlobalTimeout  ctime.Duration    $json:"global_timeout"$  //default 500ms,max time to handle the request,unless the specific handle timeout is used in HandlerTimeout in AppConfig,handler's timeout will also be effected by caller's deadline
	IdleTimeout    ctime.Duration    $json:"idle_timeout"$    //default 5s
	HeartProbe     ctime.Duration    $json:"heart_probe"$     //default 1.5s
	SrcRoot        string            $json:"src_root"$
	Certs          map[string]string $json:"certs"$ //key cert path,value private key path,if this is not empty,tls will be used
	//cors
	Cors *WebCorsConfig $json:"cors"$
}

// WebCorsConfig -
type WebCorsConfig struct {
	CorsOrigin []string $json:"cors_origin"$
	CorsHeader []string $json:"cors_header"$
	CorsExpose []string $json:"cors_expose"$
}

// WebClientConfig -
type WebClientConfig struct {
	ConnectTimeout ctime.Duration $json:"connect_timeout"$   //default 500ms,max time to finish the handshake
	GlobalTimeout  ctime.Duration $json:"global_timeout"$ //max time to handle the request,0 means no default timeout
	IdleTimeout    ctime.Duration $json:"idle_timeout"$   //default 5s
	HeartProbe     ctime.Duration $json:"heart_probe"$    //default 1.5s
}

// RedisConfig -
type RedisConfig struct {
	URL         string         $json:"url"$          //[redis/rediss]://[[username:]password@]host/[dbindex]
	MaxOpen     uint16         $json:"max_open"$     //if this is 0,means no limit //this will overwrite the param in url
	MaxIdle     uint16         $json:"max_idle"$     //defaule 100   //this will overwrite the param in url
	MaxIdletime ctime.Duration $json:"max_idletime"$ //default 10min //this will overwrite the param in url
	IOTimeout   ctime.Duration $json:"io_timeout"$   //default 500ms //this will overwrite the param in url
	ConnTimeout ctime.Duration $json:"conn_timeout"$ //default 250ms //this will overwrite the param in url
}

// SqlConfig -
type SqlConfig struct {
	URL         string         $json:"url"$          //[username:password@][protocol(address)]/[dbname][?param1=value1&...&paramN=valueN]
	MaxOpen     uint16         $json:"max_open"$     //if this is 0,means no limit //this will overwrite the param in url
	MaxIdle     uint16         $json:"max_idle"$     //default 100   //this will overwrite the param in url
	MaxIdletime ctime.Duration $json:"max_idletime"$ //default 10min //this will overwrite the param in url
	IOTimeout   ctime.Duration $json:"io_timeout"$   //default 500ms //this will overwrite the param in url
	ConnTimeout ctime.Duration $json:"conn_timeout"$ //default 250ms //this will overwrite the param in url
}

// MongoConfig -
type MongoConfig struct {
	URL         string         $json:"url"$          //[mongodb/mongodb+srv]://[username:password@]host1,...,hostN/[dbname][?param1=value1&...&paramN=valueN]
	MaxOpen     uint64         $json:"max_open"$     //if this is 0,means no limit //this will overwrite the param in url
	MaxIdletime ctime.Duration $json:"max_idletime"$ //default 10min //this will overwrite the param in url
	IOTimeout   ctime.Duration $json:"io_timeout"$   //default 500ms //this will overwrite the param in url
	ConnTimeout ctime.Duration $json:"conn_timeout"$ //default 250ms //this will overwrite the param in url
}

// KafkaPubConfig -
type KafkaPubConfig struct {
	Addrs          []string       $json:"addrs"$
	TLS            bool           $json:"tls"$
	Username       string         $json:"username"$
	Passwd         string         $json:"password"$
	AuthMethod     int            $json:"auth_method"$     //1-plain,2-scram sha256,3-scram sha512
	CompressMethod int            $json:"compress_method"$ //0-none,1-gzip,2-snappy,3-lz4,4-zstd
	TopicName      string         $json:"topic_name"$
	IOTimeout      ctime.Duration $json:"io_timeout"$   //default 500ms
	ConnTimeout    ctime.Duration $json:"conn_timeout"$ //default 250ms
}

// KafkaSubConfig -
type KafkaSubConfig struct {
	Addrs       []string       $json:"addrs"$
	TLS         bool           $json:"tls"$
	Username    string         $json:"username"$
	Passwd      string         $json:"password"$
	AuthMethod  int            $json:"auth_method"$ //1-plain,2-scram sha256,3-scram sha512
	TopicName   string         $json:"topic_name"$
	GroupName   string         $json:"group_name"$
	ConnTimeout ctime.Duration $json:"conn_timeout"$ //default 250ms
	//when there is no offset in a partition(add partition or first time to use the topic)
	//-1 will sub from the newest
	//-2 will sub from the firt
	//if this is 0,default -2 will be used
	StartOffset int64 $json:"start_offset"$
	//if this is 0,commit is synced,and effective is slow.
	//if this is not 0,commit is asynced,effective is high,but will cause duplicate sub when the program crash
	CommitInterval ctime.Duration $json:"commit_interval"$
}

// SC total source config instance
var sc *sourceConfig

var mongos map[string]*mongo.Client

var sqls map[string]*sql.DB

var rediss map[string]*redis.Pool

var kafkaSubers map[string]*kafka.Reader

var kafkaPubers map[string]*kafka.Writer

func initlocalsource() {
	data, e := os.ReadFile("./SourceConfig.json")
	if e != nil {
		log.Error(nil, "[config.local.source] read config file failed", map[string]interface{}{"error": e})
		Close()
		os.Exit(1)
	}
	sc = &sourceConfig{}
	if e = json.Unmarshal(data, sc); e != nil {
		log.Error(nil, "[config.local.source] config file format wrong", map[string]interface{}{"error": e})
		Close()
		os.Exit(1)
	}
	log.Info(nil, "[config.local.source] update success", map[string]interface{}{"config": sc})

	initgrpcserver()
	initgrpcclient()
	initcrpcserver()
	initcrpcclient()
	initwebserver()
	initwebclient()
	initredis()
	initmongo()
	initsql()
	initkafkapub()
	initkafkasub()
}
func initremotesource(wait chan *struct{}) (stopwatch func()) {
	return RemoteConfigSdk.Watch("SourceConfig", func(key, keyvalue, keytype string) {
		//only support json
		if keytype != "json" {
			log.Error(nil, "[config.remote.source] config data can only support json format", nil)
			return
		}
		//source config only init once
		if sc != nil {
			return
		}
		c := &sourceConfig{}
		if e := json.Unmarshal(common.Str2byte(keyvalue), c); e != nil {
			log.Error(nil, "[config.remote.source] config data format wrong", map[string]interface{}{"error": e})
			return
		}
		sc = c
		log.Info(nil, "[config.remote.source] update success", map[string]interface{}{"config": sc})
		initgrpcserver()
		initgrpcclient()
		initcrpcserver()
		initcrpcclient()
		initwebserver()
		initwebclient()
		initredis()
		initmongo()
		initsql()
		initkafkapub()
		initkafkasub()
		select {
		case wait <- nil:
		default:
		}
	})
}
func initgrpcserver() {
	if sc.CGrpcServer == nil {
		sc.CGrpcServer = &CGrpcServerConfig{
			ConnectTimeout: ctime.Duration(time.Millisecond * 500),
			GlobalTimeout:  ctime.Duration(time.Millisecond * 500),
			HeartProbe:     ctime.Duration(1500 * time.Millisecond),
		}
	} else {
		if sc.CGrpcServer.ConnectTimeout <= 0 {
			sc.CGrpcServer.ConnectTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.CGrpcServer.GlobalTimeout <= 0 {
			sc.CGrpcServer.GlobalTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.CGrpcServer.HeartProbe <= 0 {
			sc.CGrpcServer.HeartProbe = ctime.Duration(1500 * time.Millisecond)
		}
	}
}
func initgrpcclient() {
	if sc.CGrpcClient == nil {
		sc.CGrpcClient = &CGrpcClientConfig{
			ConnectTimeout: ctime.Duration(time.Millisecond * 500),
			GlobalTimeout:  ctime.Duration(time.Millisecond * 500),
			HeartProbe:     ctime.Duration(time.Millisecond * 1500),
		}
	} else {
		if sc.CGrpcClient.ConnectTimeout <= 0 {
			sc.CGrpcClient.ConnectTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.CGrpcClient.GlobalTimeout < 0 {
			sc.CGrpcClient.GlobalTimeout = 0
		}
		if sc.CGrpcClient.HeartProbe <= 0 {
			sc.CGrpcClient.HeartProbe = ctime.Duration(time.Millisecond * 1500)
		}
	}
}
func initcrpcserver() {
	if sc.CrpcServer == nil {
		sc.CrpcServer = &CrpcServerConfig{
			ConnectTimeout: ctime.Duration(time.Millisecond * 500),
			GlobalTimeout:  ctime.Duration(time.Millisecond * 500),
			HeartProbe:     ctime.Duration(1500 * time.Millisecond),
		}
	} else {
		if sc.CrpcServer.ConnectTimeout <= 0 {
			sc.CrpcServer.ConnectTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.CrpcServer.GlobalTimeout <= 0 {
			sc.CrpcServer.GlobalTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.CrpcServer.HeartProbe <= 0 {
			sc.CrpcServer.HeartProbe = ctime.Duration(1500 * time.Millisecond)
		}
	}
}
func initcrpcclient() {
	if sc.CrpcClient == nil {
		sc.CrpcClient = &CrpcClientConfig{
			ConnectTimeout: ctime.Duration(time.Millisecond * 500),
			GlobalTimeout:  ctime.Duration(time.Millisecond * 500),
			HeartProbe:     ctime.Duration(time.Millisecond * 1500),
		}
	} else {
		if sc.CrpcClient.ConnectTimeout <= 0 {
			sc.CrpcClient.ConnectTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.CrpcClient.GlobalTimeout < 0 {
			sc.CrpcClient.GlobalTimeout = 0
		}
		if sc.CrpcClient.HeartProbe <= 0 {
			sc.CrpcClient.HeartProbe = ctime.Duration(time.Millisecond * 1500)
		}
	}

}
func initwebserver() {
	if sc.WebServer == nil {
		sc.WebServer = &WebServerConfig{
			ConnectTimeout: ctime.Duration(time.Millisecond * 500),
			GlobalTimeout:  ctime.Duration(time.Millisecond * 500),
			IdleTimeout:    ctime.Duration(time.Second * 5),
			HeartProbe:     ctime.Duration(time.Millisecond * 1500),
			SrcRoot:        "./src",
			Cors: &WebCorsConfig{
				CorsOrigin: []string{"*"},
				CorsHeader: []string{"*"},
				CorsExpose: nil,
			},
		}
	} else {
		if sc.WebServer.ConnectTimeout <= 0 {
			sc.WebServer.ConnectTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.WebServer.GlobalTimeout <= 0 {
			sc.WebServer.GlobalTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.WebServer.IdleTimeout <= 0 {
			sc.WebServer.IdleTimeout = ctime.Duration(time.Second * 5)
		}
		if sc.WebServer.HeartProbe <= 0 {
			sc.WebServer.HeartProbe = ctime.Duration(time.Millisecond * 1500)
		}
		if sc.WebServer.Cors == nil {
			sc.WebServer.Cors = &WebCorsConfig{
				CorsOrigin: []string{"*"},
				CorsHeader: []string{"*"},
				CorsExpose: nil,
			}
		}
	}
}
func initwebclient() {
	if sc.WebClient == nil {
		sc.WebClient = &WebClientConfig{
			ConnectTimeout: ctime.Duration(time.Millisecond * 500),
			GlobalTimeout:  ctime.Duration(time.Millisecond * 500),
			IdleTimeout:    ctime.Duration(time.Second * 5),
			HeartProbe:     ctime.Duration(time.Millisecond * 1500),
		}
	} else {
		if sc.WebClient.ConnectTimeout <= 0 {
			sc.WebClient.ConnectTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sc.WebClient.GlobalTimeout < 0 {
			sc.WebClient.GlobalTimeout = 0
		}
		if sc.WebClient.IdleTimeout <= 0 {
			sc.WebClient.IdleTimeout = ctime.Duration(time.Second * 5)
		}
		if sc.WebClient.HeartProbe <= 0 {
			sc.WebClient.HeartProbe = ctime.Duration(time.Millisecond * 1500)
		}
	}
}
func initredis(){
	for k, redisc := range sc.Redis {
		if k == "example_redis" {
			continue
		}
		if redisc.MaxIdle == 0 {
			redisc.MaxIdle = 100
		}
		if redisc.MaxIdletime == 0 {
			redisc.MaxIdletime = ctime.Duration(time.Minute * 10)
		}
		if redisc.IOTimeout == 0 {
			redisc.IOTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if redisc.ConnTimeout == 0 {
			redisc.ConnTimeout = ctime.Duration(time.Millisecond * 250)
		}
	}
	rediss = make(map[string]*redis.Pool, len(sc.Redis))
	for k, redisc := range sc.Redis {
		if k == "example_redis" {
			continue
		}
		tempredis := redis.NewRedis(&redis.Config{
			RedisName:   k,
			URL:         redisc.URL,
			MaxIdle:     redisc.MaxIdle,
			MaxOpen:     redisc.MaxOpen,
			MaxIdletime: redisc.MaxIdletime.StdDuration(),
			ConnTimeout: redisc.ConnTimeout.StdDuration(),
			IOTimeout:   redisc.IOTimeout.StdDuration(),
		})
		if e := tempredis.Ping(context.Background()); e != nil {
			log.Error(nil, "[config.initredis] ping failed", map[string]interface{}{"redis": k, "error": e})
			Close()
			os.Exit(1)
		}
		rediss[k] = tempredis
	}
}
func initmongo(){
	for k, mongoc := range sc.Mongo {
		if k == "example_mongo" {
			continue
		}
		if mongoc.MaxIdletime == 0 {
			mongoc.MaxIdletime = ctime.Duration(time.Minute * 10)
		}
		if mongoc.IOTimeout == 0 {
			mongoc.IOTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if mongoc.ConnTimeout == 0 {
			mongoc.ConnTimeout = ctime.Duration(time.Millisecond * 250)
		}
	}
	mongos = make(map[string]*mongo.Client, len(sc.Mongo))
	for k, mongoc := range sc.Mongo {
		if k == "example_mongo" {
			continue
		}
		op := options.Client().ApplyURI(mongoc.URL)
		op = op.SetConnectTimeout(mongoc.ConnTimeout.StdDuration())
		op = op.SetMaxConnIdleTime(mongoc.MaxIdletime.StdDuration())
		op = op.SetMaxPoolSize(mongoc.MaxOpen)
		op = op.SetTimeout(mongoc.IOTimeout.StdDuration())
		tempdb, e := mongo.Connect(nil, op)
		if e != nil {
			log.Error(nil, "[config.initmongo] open failed", map[string]interface{}{"mongodb": k, "error": e})
			Close()
			os.Exit(1)
		}
		e = tempdb.Ping(context.Background(), readpref.Primary())
		if e != nil {
			log.Error(nil, "[config.initmongo] ping failed", map[string]interface{}{"mongodb": k, "error": e})
			Close()
			os.Exit(1)
		}
		mongos[k] = tempdb
	}
}
func initsql(){
	for _, sqlc := range sc.Sql {
		if sqlc.MaxIdle == 0 {
			sqlc.MaxIdle = 100
		}
		if sqlc.MaxIdletime == 0 {
			sqlc.MaxIdletime = ctime.Duration(time.Minute * 10)
		}
		if sqlc.IOTimeout == 0 {
			sqlc.IOTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if sqlc.ConnTimeout == 0 {
			sqlc.ConnTimeout = ctime.Duration(time.Millisecond * 250)
		}
	}
	sqls = make(map[string]*sql.DB, len(sc.Sql))
	for k, sqlc := range sc.Sql {
		if k == "example_sql" {
			continue
		}
		tmpc, e := mysql.ParseDSN(sqlc.URL)
		if e != nil {
			log.Error(nil, "[config.initsql] url format wrong", map[string]interface{}{"mysql": k, "error": e})
			Close()
			os.Exit(1)
		}
		tmpc.Timeout = sqlc.ConnTimeout.StdDuration()
		tmpc.ReadTimeout = sqlc.IOTimeout.StdDuration()
		tmpc.WriteTimeout = sqlc.IOTimeout.StdDuration()
		tempdb, e := sql.Open("mysql", tmpc.FormatDSN())
		if e != nil {
			log.Error(nil, "[config.initsql] open failed", map[string]interface{}{"mysql": k, "error": e})
			Close()
			os.Exit(1)
		}
		tempdb.SetMaxOpenConns(int(sqlc.MaxOpen))
		tempdb.SetMaxIdleConns(int(sqlc.MaxIdle))
		tempdb.SetConnMaxIdleTime(sqlc.MaxIdletime.StdDuration())
		e = tempdb.PingContext(context.Background())
		if e != nil {
			log.Error(nil, "[config.initsql] ping failed", map[string]interface{}{"mysql": k, "error": e})
			Close()
			os.Exit(1)
		}
		sqls[k] = tempdb
	}
}
func initkafkapub(){
	for _, pubc := range sc.KafkaPub {
		if pubc.TopicName == "example_topic" || pubc.TopicName == "" {
			continue
		}
		if len(pubc.Addrs) == 0 {
			pubc.Addrs = []string{"127.0.0.1:9092"}
		}
		if (pubc.AuthMethod == 1 || pubc.AuthMethod == 2 || pubc.AuthMethod == 3) && (pubc.Username == "" || pubc.Passwd == "") {
			log.Error(nil, "[config.initkafkapub] username or password missing when auth_method != 0", map[string]interface{}{"topic": pubc.TopicName})
			Close()
			os.Exit(1)
		}
		if pubc.IOTimeout == 0 {
			pubc.IOTimeout = ctime.Duration(time.Millisecond * 500)
		}
		if pubc.ConnTimeout == 0 {
			pubc.IOTimeout = ctime.Duration(time.Millisecond * 250)
		}
	}
	kafkaPubers = make(map[string]*kafka.Writer, len(sc.KafkaPub))
	for _, pubc := range sc.KafkaPub {
		if pubc.TopicName == "example_topic" || pubc.TopicName == "" {
			continue
		}
		dialer := &kafka.Dialer{
			Timeout:   pubc.ConnTimeout.StdDuration(),
			DualStack: true,
		}
		if pubc.TLS {
			dialer.TLS = &tls.Config{}
		}
		var e error
		switch pubc.AuthMethod {
		case 1:
			dialer.SASLMechanism = plain.Mechanism{Username: pubc.Username, Password: pubc.Passwd}
		case 2:
			dialer.SASLMechanism, e = scram.Mechanism(scram.SHA256, pubc.Username, pubc.Passwd)
		case 3:
			dialer.SASLMechanism, e = scram.Mechanism(scram.SHA512, pubc.Username, pubc.Passwd)
		}
		if e != nil {
			log.Error(nil, "[config.initkafkapub] username and password wrong",map[string]interface{}{"topic": pubc.TopicName, "error": e})
			Close()
			os.Exit(1)
		}
		var compressor kafka.CompressionCodec
		switch pubc.CompressMethod {
		case 1:
			compressor = kafka.Gzip.Codec()
		case 2:
			compressor = kafka.Snappy.Codec()
		case 3:
			compressor = kafka.Lz4.Codec()
		case 4:
			compressor = kafka.Zstd.Codec()
		}
		writer := kafka.NewWriter(kafka.WriterConfig{
			Brokers:          pubc.Addrs,
			Topic:            pubc.TopicName,
			Dialer:           dialer,
			ReadTimeout:      pubc.IOTimeout.StdDuration(),
			WriteTimeout:     pubc.IOTimeout.StdDuration(),
			Balancer:         &kafka.Hash{},
			MaxAttempts:      3,
			RequiredAcks:     int(kafka.RequireAll),
			Async:            false,
			CompressionCodec: compressor,
		})
		kafkaPubers[pubc.TopicName] = writer
	}
}
func initkafkasub(){
	for _, subc := range sc.KafkaSub {
		if subc.TopicName == "example_topic" || subc.TopicName == "" {
			continue
		}
		if len(subc.Addrs) == 0 {
			subc.Addrs = []string{"127.0.0.1:9092"}
		}
		if (subc.AuthMethod == 1 || subc.AuthMethod == 2 || subc.AuthMethod == 3) && (subc.Username == "" || subc.Passwd == "") {
			log.Error(nil, "[config.initkafkasub] username or password missing when auth_method != 0", map[string]interface{}{"topic": subc.TopicName})
			Close()
			os.Exit(1)
		}
		if subc.GroupName == "" {
			log.Error(nil, "[config.initkafkasub] groupname missing", map[string]interface{}{"topic": subc.TopicName})
			Close()
			os.Exit(1)
		}
		if subc.ConnTimeout == 0 {
			subc.ConnTimeout = ctime.Duration(time.Millisecond * 250)
		}
	}
	kafkaSubers = make(map[string]*kafka.Reader, len(sc.KafkaSub))
	for _, subc := range sc.KafkaSub {
		if subc.TopicName == "example_topic" || subc.TopicName == "" {
			continue
		}
		dialer := &kafka.Dialer{
			Timeout:   subc.ConnTimeout.StdDuration(),
			DualStack: true,
		}
		if subc.TLS {
			dialer.TLS = &tls.Config{}
		}
		var e error
		switch subc.AuthMethod {
		case 1:
			dialer.SASLMechanism = plain.Mechanism{Username: subc.Username, Password: subc.Passwd}
		case 2:
			dialer.SASLMechanism, e = scram.Mechanism(scram.SHA256, subc.Username, subc.Passwd)
		case 3:
			dialer.SASLMechanism, e = scram.Mechanism(scram.SHA512, subc.Username, subc.Passwd)
		}
		if e != nil {
			log.Error(nil, "[config.initkafkasub] username and password wrong", map[string]interface{}{"topic": subc.TopicName, "error": e})
			Close()
			os.Exit(1)
		}
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:                subc.Addrs,
			Dialer:                 dialer,
			Topic:                  subc.TopicName,
			GroupID:                subc.GroupName,
			StartOffset:            subc.StartOffset,
			MinBytes:               1,
			MaxBytes:               1024 * 1024 * 10,
			CommitInterval:         time.Duration(subc.CommitInterval),
			IsolationLevel:         kafka.ReadCommitted,
			PartitionWatchInterval: time.Second,
			WatchPartitionChanges:  true,
			MaxAttempts:            3,
		})
		kafkaSubers[subc.TopicName+subc.GroupName] = reader
	}
}

// GetCGrpcServerConfig get the grpc net config
func GetCGrpcServerConfig() *CGrpcServerConfig {
	return sc.CGrpcServer
}

// GetCGrpcClientConfig get the grpc net config
func GetCGrpcClientConfig() *CGrpcClientConfig {
	return sc.CGrpcClient
}

// GetCrpcServerConfig get the crpc net config
func GetCrpcServerConfig() *CrpcServerConfig {
	return sc.CrpcServer
}

// GetCrpcClientConfig get the crpc net config
func GetCrpcClientConfig() *CrpcClientConfig {
	return sc.CrpcClient
}

// GetWebServerConfig get the web net config
func GetWebServerConfig() *WebServerConfig {
	return sc.WebServer
}

// GetWebClientConfig get the web net config
func GetWebClientConfig() *WebClientConfig {
	return sc.WebClient
}

// GetMongo get a mongodb client by db's instance name
// return nil means not exist
func GetMongo(mongoname string) *mongo.Client {
	return mongos[mongoname]
}

// GetSql get a mysql db client by db's instance name
// return nil means not exist
func GetSql(mysqlname string) *sql.DB {
	return sqls[mysqlname]
}

// GetRedis get a redis client by redis's instance name
// return nil means not exist
func GetRedis(redisname string) *redis.Pool {
	return rediss[redisname]
}

// GetKafkaSuber get a kafka sub client by topic and groupid
func GetKafkaSuber(topic string, groupid string) *kafka.Reader {
	return kafkaSubers[topic+groupid]
}

// GetKafkaPuber get a kafka pub client by topic name
func GetKafkaPuber(topic string) *kafka.Writer {
	return kafkaPubers[topic]
}`

func CreatePathAndFile(packagename string) {
	if e := os.MkdirAll("./config/", 0755); e != nil {
		panic("mkdir ./config/ error: " + e.Error())
	}
	//./config/config.go
	configtemplate, e := template.New("./config/config.go").Parse(txt)
	if e != nil {
		panic("parse ./config/config.go template error: " + e.Error())
	}
	configfile, e := os.OpenFile("./config/config.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic("open ./config/config.go error: " + e.Error())
	}
	if e := configtemplate.Execute(configfile, packagename); e != nil {
		panic("write ./config/config.go error: " + e.Error())
	}
	if e := configfile.Sync(); e != nil {
		panic("sync ./config/config.go error: " + e.Error())
	}
	if e := configfile.Close(); e != nil {
		panic("close ./config/config.go error: " + e.Error())
	}
	//./config/app_config.go
	appfile, e := os.OpenFile("./config/app_config.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic("open ./config/app_config.go error: " + e.Error())
	}
	if _, e := appfile.WriteString(strings.ReplaceAll(apptxt, "$", "`")); e != nil {
		panic("write ./config/app_config.go error: " + e.Error())
	}
	if e := appfile.Sync(); e != nil {
		panic("sync ./config/app_config.go error: " + e.Error())
	}
	if e := appfile.Close(); e != nil {
		panic("close ./config/app_config.go error: " + e.Error())
	}
	//./config/source_config.go
	sourcefile, e := os.OpenFile("./config/source_config.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic("open ./config/source_config.go error: " + e.Error())
	}
	if _, e := sourcefile.WriteString(strings.ReplaceAll(sourcetxt, "$", "`")); e != nil {
		panic("write ./config/source_config.go error: " + e.Error())
	}
	if e := sourcefile.Sync(); e != nil {
		panic("sync ./config/source_config.go error: " + e.Error())
	}
	if e := sourcefile.Close(); e != nil {
		panic("close ./config/source_config.go error: " + e.Error())
	}
}
