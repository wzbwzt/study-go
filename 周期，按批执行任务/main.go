package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

var maxnum chan struct{} = make(chan struct{}, 5)

//周期性的开启指定数量的携程并发完成，并保证当前批次完成后，再开启同样批次的携程完成任务
func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(5)
	setPool(maxnum)
	go func() {
		for {
			wg.Wait()
			fmt.Println("发送5个任务")
			setPool(maxnum)
			wg.Add(5)
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			<-maxnum
			go doJob()
		}
	}()
	select {}
}

func setPool(pool chan struct{}) {
	for i := 0; i < 5; i++ {
		pool <- struct{}{}
	}
}

func doJob() {
	defer wg.Done()
	fmt.Println(rand.Int63n(100))
}
