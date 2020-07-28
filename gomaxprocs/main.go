package main

import (
	"fmt"
	"runtime"
	"sync"
)

//goroutine的调度模型GMP
//M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一
//映射的关系， 一个groutine最终是要放到M上执行的；
//G很好理解，就是个goroutine的，里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
//P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地
//址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的gorout
//ine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里
//也消费完了会去其他P的队列里抢任务。（P的个数是通过runtime.GOMAXPROCS设定（最大256））
//P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M
//上时，runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者
//认为其已经死掉时 回收旧的M。

//goroutine初始栈的内存是大小是2K，可以按需增加或缩小；最大1G
//OS线程有固定的栈内存，为2MB
var wg sync.WaitGroup

func main() {
	//默认使用物理线程数；当指定1个OS线程时；下面函数按照顺序执行；当使用os线程数大于1时，下面的函数会同时执行
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU()) //查看物理线程数，即CPU的逻辑核心数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}



