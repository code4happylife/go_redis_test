package middleware

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"go_redis_test/common"
)

var (
	RedisCli *redis.Client
)

func InitRedis(db int) *redis.Client {
	addr := common.GetConfigValue("redis", "addr")
	password := common.GetConfigValue("redis", "password")
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return RedisCli
}

func InitRedisForCustomer(db int, addr, pwd string) (*redis.Client, error) {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	})
	return RedisCli, RedisCli.Ping().Err()
}

func redisExample() {
	fmt.Printf("Redis connect demo")
}
