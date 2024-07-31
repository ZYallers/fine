package fredis

import (
	"fmt"

	"github.com/go-redis/redis"
	"gitlab.sys.hxsapp.net/hxs/fine/frame/fmsg"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/consts"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/instance"
	"gitlab.sys.hxsapp.net/hxs/fine/internal/intlog"
	"gitlab.sys.hxsapp.net/hxs/fine/os/flog"
)

func init() {
	redis.SetLogger(intlog.NewGoRedisLogger(flog.Use("redis")))
}

func ClientFunc(name string) func() *redis.Client {
	return func() *redis.Client { return Client(name) }
}

func Client(name string) *redis.Client {
	if name == "" {
		return nil
	}
	instanceKey := fmt.Sprintf("%s.%s", consts.FrameComponentDatabaseRedis, name)
	result := instance.GetOrSetFuncLock(instanceKey, func() interface{} {
		flog.Use().Info(fmt.Sprintf("redis.%s.init", name))
		options := getRedisOptions(name)
		client := redis.NewClient(options)
		if err := client.Ping().Err(); err != nil {
			return fmt.Errorf("redis.%s.ping.error: %s", name, err)
		}
		return client
	})
	switch v := result.(type) {
	case *redis.Client:
		return v
	case error:
		instance.Remove(instanceKey)
		flog.Use().Error(v.Error())
		fmsg.Sender().Simple(v.Error(), true)
		return nil
	default:
		instance.Remove(instanceKey)
		msg := fmt.Sprintf("redis.%s.error: %s", name, "unknown error")
		flog.Use().Error(msg)
		fmsg.Sender().Simple(msg, true)
		return nil
	}
}
