package datasource

import "github.com/go-redis/redis"

// 声明一个全局的redisdb变量
var redisdb *redis.Client

// 初始化连接
func init() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "23.238.49.110:6379",
		Password: "710069741", // no password set
		DB:       0,  // use default DB
	})
}

func GetRedis() *redis.Client{
	return redisdb
}
