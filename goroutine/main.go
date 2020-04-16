package main

import (
	"fmt"
	"time"
)

//程序启动之后也会创建一个主goroutine（main） 去执行
//goroutine 什么时候结束？goroutine对应的函数结束了，goroutine也就结束了
//main 函数执行完了；由main函数创建的那些goroutine也就结束了
func main() {
	for i := 0; i < 100; i++ {
		//go hello(i) //开启一个单独的goroutine去执行hello函数（任务）
		go func() {
			fmt.Println(i) //相当于闭包，从外面拿取i值；而不是传进来的值；就会导致；同一数字多次打印执行
		}()
	}
	fmt.Println("main")
	time.Sleep(time.Second * 1) //自定义时间来等待goroutine结束；不够规范；具体的需要调用wg包来实现；详见waitGroup文件
	//main函数结束后 由main函数启动的goroutine也都会结束
}

func hello(i int) {
	fmt.Println("hello word", i)
}

/*
goroutine的本质
goroutine的调度模型：GMP
m:n  把m个goroutine分配个n个操作系统线程

goroutine 与操作系统的线程（OS线程）的区别
goroutine 是用户态的线程，比内核态的线程更加轻量级一点；初始时只占用2KB的内存；
可以轻松开启数十万的goroutine也不会崩内存

*/
