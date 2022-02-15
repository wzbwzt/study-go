package main

import (
	"fmt"
	"time"
)

//扇出：将一条channel数据发散到多个channel中处理

func fanOut(source <-chan interface{}, chs ...chan interface{}) {
	go func() {
		defer func() {
			for _, ch := range chs {
				close(ch)
			}
		}()
		for v := range source {
			for _, ch := range chs {
				ch <- v
			}
		}
	}()
}

func redisJob() chan interface{} {
	var res = make(chan interface{})
	go func() {
		for v := range res {
			time.Sleep(time.Second * 1)
			fmt.Printf("data %d insert to redis,success\n", v)
		}
	}()
	return res
}

func mysqlJob() chan interface{} {
	var res = make(chan interface{})
	go func() {
		for v := range res {
			time.Sleep(time.Second * 1)
			fmt.Printf("data %d insert to mysql,success\n", v)
		}
	}()
	return res
}

//##############################################################################
func main() {
	sourceCh := make(chan interface{})
	fanOut(sourceCh, redisJob(), mysqlJob())

	for i := 0; i < 6; i++ {
		sourceCh <- i
	}
	close(sourceCh)

	for {
	}
}
