package redis

import "github.com/redis/go-redis/v9"

var (
	rc *redis.Client
	RC RedisClient
)

const (
	addr      = "121.37.143.160:6379"
	pass_word = ""
	db        = 0
	shop_key  = "SHOP_STATUS"
)

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass_word, // 没有密码，默认值
		DB:       db,        // 默认DB 0
	})
	rc = rdb
	RC = &redisClient{}
}
