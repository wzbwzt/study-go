package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//redis 示例
var redisdb *redis.Client

//普通连接
func initRedis1() (err error) {
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

//哨兵模式
func initRedis2() (err error) {
	redisdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{":26379", ":26379", ":26379"},
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		fmt.Println("redis connect failed err:", err)
		return
	}
	return nil
}

//集群模式
func initRedis3() (err error) {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = redisdb.Ping().Result()
	if err != nil {
		fmt.Println("redis connect failed err:", err)
		return
	}
	return nil
}

//连接池连接	"github.com/garyburd/redigo/redis"
//"github.com/go-redis/redis" 不支持连接池
func initRedis() (err error) {
	// redisPool := &redis.Pool{
	// 	MaxIdle:     10,                               //最大闲置数
	// 	MaxActive:   100,                              //最大连接数
	// 	IdleTimeout: time.Duration(100) * time.Second, //连接超时
	// 	Dial: func() (redis.Conn, error) {
	// 		return redis.Dial("tcp", secKillConfig.redisConf.redisAddr)
	// 	},
	// }

	// //测试连接池
	// conn := redisPool.Get()
	// _, err = conn.Do("ping")
	// if err != nil {
	// 	return
	// }
	// return
}
func main() {
	initRedis1()

}
