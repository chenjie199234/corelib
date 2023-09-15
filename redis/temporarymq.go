package redis

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/chenjie199234/Corelib/log"
	"github.com/chenjie199234/Corelib/util/common"

	gredis "github.com/redis/go-redis/v9"
)

// name_0(redis list):[data1,data2,data3...]
// {name_0}_exist(redis string):1
// name_1(redis list):[data1,data2,data3...]
// {name_1}_exist(redis string):1
// ...
// name_n(redis list):[data1,data2,data3...]
// {name_n}_exist(redis string):1

var ErrTemporaryMQMissingName = errors.New("temporary mq missing name")
var ErrTemporaryMQMissingGroup = errors.New("temporary mq missing group")
var ErrTemporaryMQMissingReceiver = errors.New("temporary mq missing receiver")

var expireTMQ *gredis.Script
var pubTMQ *gredis.Script

func init() {
	expireTMQ = gredis.NewScript(`redis.call("SETEX",KEYS[2],16,1)
redis.call("EXPIRE",KEYS[1],16)
return "OK"`)

	pubTMQ = gredis.NewScript(`if(redis.call("EXISTS",KEYS[2])==0)
then
	return -1
end
redis.call("EXPIRE",KEYS[1],16)
redis.call("RPUSH",KEYS[1],unpack(ARGV))
redis.call("EXPIRE",KEYS[1],16)
return #ARGV`)
}

// Warning!this module will take group*2 redis connections,be careful of the client's MaxOpen
// in redis cluster mode,group is used to split data into different redis node
// in redis slave master mode,group is better to be 1
// sub and pub's mqname and group should be same
// stop will stop the sub immediately,even if there are datas int the mq,the left datas will be expired within 16s
func (c *Client) TemporaryMQSub(mqname string, group uint64, subhandler func([]byte)) (stop func(), e error) {
	if mqname == "" {
		return nil, ErrTemporaryMQMissingName
	}
	if group == 0 {
		return nil, ErrTemporaryMQMissingGroup
	}
	if e := c.temporaryMQSubRefresh(context.Background(), mqname, group); e != nil {
		return nil, e
	}
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	stop = func() {
		cancel()
		wg.Wait()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		tmer := time.NewTimer(time.Second * 5)
		for {
			select {
			case <-ctx.Done():
				//stop
				tmer.Stop()
				c.temporaryMQSubClean(context.Background(), mqname, group)
				return
			case <-tmer.C:
				//refresh
				if e := c.temporaryMQSubRefresh(ctx, mqname, group); e != nil && ctx.Err() != nil {
					break
				} else if e != nil {
					tmer.Reset(time.Millisecond * 500)
				} else {
					tmer.Reset(time.Second * 5)
				}
			}
		}
	}()
	for i := uint64(0); i < group; i++ {
		wg.Add(1)
		go func(index uint64) {
			defer wg.Done()
			c.temporaryMQSubHandle(ctx, mqname, index, subhandler)
		}(i)
	}
	return
}
func (c *Client) temporaryMQSubRefresh(ctx context.Context, mqname string, group uint64) error {
	var err error
	wg := sync.WaitGroup{}
	for i := uint64(0); i < group; i++ {
		wg.Add(1)
		go func(index uint64) {
			defer wg.Done()
			listname := mqname + "_" + strconv.FormatUint(index, 10)
			listexist := "{" + listname + "}_exist"
			if _, e := expireTMQ.Run(ctx, c, []string{listname, listexist}).Result(); e != nil {
				log.Error(ctx, "[redis.temporaryMQSubRefresh] failed", map[string]interface{}{"group": index, "error": e})
				err = e
			}
		}(i)
	}
	wg.Wait()
	return err
}
func (c *Client) temporaryMQSubClean(ctx context.Context, mqname string, group uint64) error {
	var err error
	wg := sync.WaitGroup{}
	for i := uint64(0); i < group; i++ {
		wg.Add(1)
		go func(index uint64) {
			defer wg.Done()
			listname := mqname + "_" + strconv.FormatUint(index, 10)
			listexist := "{" + listname + "}_exist"
			if _, e := c.Del(ctx, listexist).Result(); e != nil {
				log.Error(ctx, "[redis.TemporaryMQSubClean] failed", map[string]interface{}{"group": index, "error": e})
				err = e
			}
		}(i)
	}
	wg.Wait()
	return err
}
func (c *Client) temporaryMQSubHandle(ctx context.Context, mqname string, index uint64, handler func([]byte)) {
	listname := mqname + "_" + strconv.FormatUint(index, 10)
	var result []string
	var e error
	for {
		if e != nil && ctx.Err() == nil {
			time.Sleep(time.Millisecond * 10)
		}
		if ctx.Err() != nil {
			//stopped
			return
		}
		if result, e = c.BLPop(ctx, time.Second, listname).Result(); e == nil {
			handler(common.Str2byte(result[1]))
		} else if ee, ok := e.(interface{ Timeout() bool }); (!ok || !ee.Timeout()) && e != gredis.Nil {
			log.Error(ctx, "[redis.temporaryMQSubHandle] failed", map[string]interface{}{"group": index, "error": e})
		} else {
			e = nil
		}
	}
}

// in redis cluster mode,group is used to split data into different redis node
// in redis slave master mode,group is better to be 1
// sub and pub's mqname and group should be same
// key is only used to caculate the data's group(hash)
func (c *Client) TemporaryMQPub(ctx context.Context, mqname string, group uint64, key string, values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}
	if mqname == "" {
		return ErrTemporaryMQMissingName
	}
	if group == 0 {
		return ErrTemporaryMQMissingGroup
	}
	listname := mqname + "_" + strconv.FormatUint(common.BkdrhashString(key, group), 10)
	listexist := "{" + listname + "}_exist"
	r, e := pubTMQ.Run(ctx, c, []string{listname, listexist}, values...).Int()
	if r == -1 {
		e = ErrTemporaryMQMissingReceiver
	}
	return e
}
