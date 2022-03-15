// +build ignore

package main

import (
	"fmt"
	"sync"
	"time"
)

//交替打印
func main1() {
	//fmt.Println("123")
	c1 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- 1
			if i%2 == 0 {
				fmt.Print(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-c1
			if i%2 == 1 {
				fmt.Print(i)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func main() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("start working")
		time.Sleep(time.Second * 1)
		<-ch
	}()

	ch <- struct{}{}

	fmt.Println("finished")
	mapstructure.WeakDecode()
}
