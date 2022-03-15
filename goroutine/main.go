package main

import (
	"fmt"
	"time"
)

//程序启动之后也会创建一个主goroutine（main） 去执行
//goroutine 什么时候结束？goroutine对应的函数结束了，goroutine也就结束了
//main 函数执行完了；由main函数创建的那些goroutine也就结束了

/*
协程的主动退出：通过 runtime.Goexit()可以做到提前结束协程，且结束前还能执行到defer的内容•
runtime.Goexit()其实是对goexit0的封装，只要执行 goexit0 这个函数，当前协程就会退出，
同时还能调度下一个可执行的协程出来跑
*/
func main() {
	for i := 0; i < 100; i++ {
		//go hello(i) //开启一个单独的goroutine去执行hello函数（任务）
		go func(ii int) {
			fmt.Println(ii) //相当于闭包，从外面拿取i值；而不是传进来的值；就会导致；同一数字多次打印执行
		}(i)
	}
	fmt.Println("main")
	// Test()//注意：是main函数结束后所有的goroutine才会结束，并不是上级的函数结束后结束子goruntine
	time.Sleep(time.Second * 3) //自定义时间来等待goroutine结束；不够规范；具体的需要调用wg包来实现；详见waitGroup文件
	//main函数结束后 由main函数启动的goroutine也都会结束
}

func hello(i int) {
	fmt.Println("hello word", i)
}

func Test() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("hello word")
	return
}

/*
goroutine的本质
goroutine的调度模型：GMP
m:n  把m个goroutine分配个n个操作系统线程

goroutine 与操作系统的线程（OS线程）的区别
goroutine 是用户态的线程，比内核态的线程更加轻量级一点；初始时只占用2KB的内存；
可以轻松开启数十万的goroutine也不会崩内存

*/
