package cli

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/go-redis/redis/v7"
	"go_redis_test/middleware"
	"log"

	"reflect"
	"strings"
)

//go:embed redis_info.txt
var file embed.FS

var RedisCli *redis.Client

func RedisInit(db int) *redis.Client {
	fileContent, err := file.ReadFile("redis_info.txt")
	if err != nil {
		log.Fatal("read redis config fail:", err)
	}
	redisInfo := strings.Split(string(fileContent), ",")
	redisCli, err := middleware.InitRedisForCustomer(db, redisInfo[0], redisInfo[1])
	fmt.Printf("redis server ip:%v\n", redisInfo[0])
	fmt.Printf("redis server password:%v\n", redisInfo[1])
	if err != nil {
		log.Fatalf("redis init fail: %v", err)
	}
	log.Println("redis init success")
	RedisCli = redisCli
	return RedisCli
}

func RedisInitFuntest(db int) {
	redisCli := middleware.InitRedis(db)
	fmt.Println(reflect.TypeOf(redisCli))
}
