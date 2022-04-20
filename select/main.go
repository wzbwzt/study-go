//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"runtime"
)

//select 多路复用
//同一时刻对多个通道实现操作的场景；
/*
可处理一个或多个channel的发送/接收操作。

select的运行机制如下：
- 选取一个可执行不阻塞的case分支，如果多个case分支都不阻塞，会随机选一个case分支执行，和case分支在代码里写的顺序没关系。
- 如果所有case分支都阻塞，会进入default分支执行。
- 如果没有default分支，那select会阻塞，直到有一个case分支不阻塞。
- 对于 select 语句，在进入该语句时，会按源码的顺序对每一个 case 子句进行求值：这个求值只针对发送或接收操作的额外表达式。（详见notion-Journal-GO）
*/

func main() {
	var ch1 = make(chan int, 1) //0,2,4,6,8;因为只有一个缓存区;如果缓存大于1的话，case随机执行；无法预测结果
	for i := 0; i < 10; i++ {
		select { //哪个case可以执行就执行哪个
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:

		}
	}

	audition1()

}

//面试题1
//程序可能panic，也可能打印"CLOSED, "
func audition1() {
	data := make(chan int)
	shutdown := make(chan int)
	close(shutdown)
	close(data)

	select {
	case <-shutdown:
		fmt.Print("CLOSED, ")
	case data <- 1:
		fmt.Print("HAS WRITTEN, ")
	default:
		fmt.Print("DEFAULT, ")
	}
	runtime.GC()
}
