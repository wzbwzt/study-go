package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done() //没执行一次减1
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //没启动一次加1
		go f1(i)  //启动10个goroutine
	}
	wg.Wait() //等待wg的计数器减为0时结束goroutine
}
