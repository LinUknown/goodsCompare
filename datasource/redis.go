package datasource

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的redisdb变量
var redisdb *redis.Client

// 初始化连接
func init() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "47.112.216.17:6379",
		Password: "710069741", // no password set
		DB:       0,  // use default DB
	})
	pong, err := redisdb.Ping().Result()
	fmt.Printf("ping = %v,err=%v",pong, err)
}

func GetRedis() *redis.Client{
	return redisdb
}
