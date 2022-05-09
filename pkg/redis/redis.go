package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"goapi/pkg/config"
	"goapi/pkg/helpers"
	"sync"
	"time"
)

// Client *Client 对象
var Client *Redis

// Redis 对象
type Redis struct {
	mutex      sync.Mutex    // 锁
	Client     *redis.Client // 客户端
	expireTime int           // 过期时间
}

// ConnectDB 初始化模型
func (rdb *Redis) ConnectDB(selectDB string) {
	rdb.expireTime = 600 // 过期时间
	// 初始化 Redis 连接信息
	var (
		err       error
		RedisIp   = config.GetString("redis.host")
		RedisPort = config.GetString("redis.port")
		DefaultDB = config.GetInt("redis.db")
		Pw        = config.GetString("redis.password")
	)
	if len(selectDB) > 0 {
		// 选择库
		DefaultDB = helpers.StringToInt(selectDB)
	}
	if len(Pw) > 0 {
		rdb.Client = redis.NewClient(&redis.Options{
			Addr:     RedisIp + ":" + RedisPort,
			DB:       DefaultDB, // use default DB
			Password: Pw,        // no password set
		})
	} else {
		rdb.Client = redis.NewClient(&redis.Options{
			Addr: RedisIp + ":" + RedisPort,
			DB:   DefaultDB, // use default DB
		})
	}
	_, err = rdb.Client.Ping().Result()
	if err != nil {
		fmt.Println("redis连接错误" + err.Error())
	}
	// 赋值
	Client = rdb
}

func (rdb *Redis) CheckExist(key string) bool {
	a, err := rdb.Client.Exists(key).Result()
	if err != nil {
		fmt.Println("判断key存在失败")
		return false
	}
	if a == 1 {
		fmt.Println("key存在")
		return true
	}
	return false
}

func (rdb *Redis) Add(key string, value interface{}, exTime int) (bool, error) {
	if exTime >= 0 {
		rdb.expireTime = exTime
	}
	err := rdb.Client.Set(key, value, time.Duration(rdb.expireTime)*time.Second).Err()
	if err != nil {
		fmt.Println("设置key失败")
		return false, err
	}
	return true, nil
}

func (rdb *Redis) Incr(key string) int64 {
	return rdb.Client.Incr(key).Val()
}

func (rdb *Redis) Lpush(key string, values ...interface{}) (bool, error) {
	_, err := rdb.Client.LPush(key, values...).Result()
	if err != nil {
		fmt.Println("Lpush 失败")
		return false, err
	}
	return true, nil
}

func (rdb *Redis) RPop(key string) string {
	return rdb.Client.RPop(key).Val()
}

func (rdb *Redis) Delete(key string) bool {
	err := rdb.Client.Del(key).Err()
	if err != nil {
		fmt.Println("删除key失败" + err.Error())
		return false
	}
	return true
}

func (rdb *Redis) Get(key string) (string, error) {
	defer rdb.mutex.Unlock()
	rdb.mutex.Lock()
	value, err := rdb.Client.Get(key).Result()
	return value, err
}

func (rdb *Redis) SelectDB(db int) {
	rdb.Client.Do("select", db)
}

func (rdb *Redis) GetKeys(key string) ([]string, error) {
	value, err := rdb.Client.Keys(key).Result()
	return value, err
}

// Redis Keys 命令 - 查找所有符合给定模式( pattern)的 key
// https://www.redis.net.cn/order/3535.html

func (rdb *Redis) Keys(key string) ([]string, error) {
	value, err := rdb.Client.Keys(key).Result()
	return value, err
}

func (rdb *Redis) SubExpireEvent(channels string) *redis.PubSub {
	// 订阅key过期事件
	//sub := Redis.Subscribe("__keyevent@0__:expired")
	return rdb.Client.Subscribe(channels)
	// 这里通过一个for循环监听redis-server发来的消息。
	// 当客户端接收到redis-server发送的事件通知时，
	// 客户端会通过一个channel告知我们。我们再根据
	// msg的channel字段来判断是不是我们期望收到的消息，
	// 然后再进行业务处理。
	//for {
	//	msg := <-sub.Channel()
	//	fmt.Println("Channel ", msg.Channel)
	//	fmt.Println("pattern ", msg.Pattern)
	//	fmt.Println("pattern ", msg.Payload)
	//}
}

func (rdb *Redis) Close() {
	err := rdb.Client.Close()
	if err != nil {
		fmt.Println("RedisCloseError", err.Error())
	}
}
