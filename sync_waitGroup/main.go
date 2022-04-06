package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*

在WaitGroup里主要有三个方法:

- Add, 可以添加或减少 goroutine的数量.
- Done, 相当于Add(-1).
- Wait, 执行后会堵塞主线程，直到WaitGroup 里的值减至0.
在主goroutine 中 Add(delta int) 索要等待goroutine 的数量。在每一个goroutine
完成后Done()表示这一个goroutine 已经完成，当所有的 goroutine 都完成后，
在主 goroutine 中 WaitGroup 返回。

注意:
在Golang官网中对于WaitGroup介绍是A WaitGroup must not be copied after first use,在 WaitGroup 第一次使用后，不能被拷贝。

WaitGroup 实现原理:
WaitGroup 主要维护了 2 个计数器，一个是请求计数器 v，一个是等待计数器 w，二者组成一个 64bit 的值，请求计数器占高 32bit，等待计数器占低32bit。
每次 Add 执行，请求计数器 v 加 1，Done 方法执行，等待计数器减 1，v 为0 时通过信号量唤醒 Wait()
*/
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

func test() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Printf("i:%d", i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
	//会报错：
	/*
		i:1i:3i:2i:0i:4fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [semacquire]:
		sync.runtime_Semacquire(0xc000094018)
		        /home/keke/soft/go/src/runtime/sema.go:56 +0x39
		sync.(*WaitGroup).Wait(0xc000094010)
		        /home/keke/soft/go/src/sync/waitgroup.go:130 +0x64
		main.main()
		        /home/keke/go/Test/wait.go:17 +0xab
		exit status 2
	*/

	/*
		它提示所有的 goroutine 都已经睡眠了，出现了死锁。这是因为 wg 给拷贝传递到了 goroutine 中，导致只有 Add 操作，其实 Done操作是在 wg 的副本执行的。

		因此 Wait 就会死锁。

		这个第一个修改方式: 将匿名函数中 wg 的传入类型改为 *sync.WaitGroup,这样就能引用到正确的WaitGroup了。

		这个第二个修改方式: 将匿名函数中的 wg 的传入参数去掉，因为Go支持闭包类型，在匿名函数中可以直接使用外面的 wg 变量.
	*/
}
