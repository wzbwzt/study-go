package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//1.开启一个goroutine生成int64类型的随机数，发送到jobChan
//2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//3.主goroutine从resultChan取出结果并打印到终端输出

//需要同时从多个通道接收数据  使用select多路复用 详见select

var jobChan chan int64
var resultChan chan int64
var wg sync.WaitGroup

func main() {
	jobChan = make(chan int64, 100)
	resultChan = make(chan int64, 50)
	go proInt(jobChan)
	for i := 1; i <= 24; i++ { //worker pool
		go spendInt(jobChan, resultChan)
	}
	for x := range resultChan {
		time.Sleep(time.Second)
		x = <-resultChan
		fmt.Println(x)
	}
}

//生产者 生成int64类型的随机数，发送到jobChan
func proInt(jobChan chan<- int64) {
	for {
		jobChan <- rand.Int63()
	}

}

//消费者 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan

func spendInt(jobChan <-chan int64, resultChan chan<- int64) {
	for {
		job := <-jobChan
		var sum = int64(0)
		for job > 0 {
			i := job % 10
			job = job / 10
			sum += i
		}
		resultChan <- sum
	}
}
