package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"go_redis_test/cli"
)

func redisTest() {

	/**
	create redis server:
	docker run --name redis507 -p 6379:6379 -d redis:5.0.7

	*/
	client := cli.RedisInit(1)
	err := client.Set("salary", 1000, 0).Err()
	if err != nil {
		fmt.Printf("set salary failed, err:%v", err)
		return
	}
	salary, err := client.Get("salary").Result()
	if err != nil {
		fmt.Printf("get salary failed, err:%v", err)
		return
	}
	fmt.Printf("salary is:%v\n", salary)

	err = client.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	val, err := client.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := client.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}

}

func main() {
	redisTest()
}
