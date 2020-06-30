package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//redis 示例
var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		fmt.Println("redis connect failed err:", err)
		return
	}

	return nil
}

func main() {
	initRedis()

}
