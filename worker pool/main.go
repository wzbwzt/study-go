package main

import (
	"fmt"
	"time"
)

//worker pool（goroutine池）

func main() {
	var jobs = make(chan int, 100)
	var result = make(chan int, 100)
	//5个工作
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	//开启3个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, result)
	}
	//输出结果
	for r := 0; r < 5; r++ {
		<-result
	}
}

//示例
func worker(work int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d;start:%d\n", work, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d;end:%d\n", work, j)
		result <- j * 2
	}
}
