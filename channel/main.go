package main

import (
	"fmt"
	"sync"
)

//channle就三种操作
//发送（send）、接收(receive）和关闭（close）三种操作。
//通道关闭后还是可以receive
/*
为什么需要channel?
通过通道实现多个goroutine之间的通信
CSP:通过通信来共享内存；
而通过变量来传递参数的方法是通过共享内存来实现通信
*/
/*
channel是一种引用类型；需要make初始化才可以使用；
需要用到make初始化的也就只有slice\map\channel
*/

/*
channel 无值时往外取会阻塞；channel缓存已满再往里面send值也会阻塞；阻塞即系统会一直处于等待状态，死锁（前提是通道没有关闭）
只有当通道关闭时才会往外取对应类型的零值；
所以当循环往通道外取值时，使用for x:=range y{};系统会自动判断是否结束；
而使用for循环时  则需要判断是否OK；
for {
	c,ok:=<-ch1
	if !ok{
		break
	}
}

*/
// 	}
// }

// */

var c chan int
var wg sync.WaitGroup

func main() {
	noBufChan()
	// bufChan()
	//channelDemo()
}

func noBufChan() {
	//chan必须使用make来实例化分配一块内存才可以使用，无缓冲区，必须先定义一个接受方，才可以发送
	c = make(chan int)
	//c <- 10                //hang住了 程序无法执行下去了  值发送不进去，没有接受的缓存区；所以必须先定义一个goroutine来接受，再发送
	//wg.Add(1)
	go func() {
		//defer wg.Done()
		x := <-c
		fmt.Println("X:", x)
	}()
	c <- 10
	fmt.Println("C:", c)
	//wg.Wait()
	close(c)

}

func bufChan() {
	//有缓冲区 ，可以往里面先存值,通道可以先直接返送，而不需接受方
	c = make(chan int, 1) //定义1个就是说只能往通道中发送一个值，发送多个就会报错，原理同无buf
	c <- 10
	fmt.Println(c)
	close(c)
}

//channel练习
//1.启动一个goroutine,生成100个数发送到ch1
//2.启动一个goroutine,从ch1中取值，计算其平方发送到ch2
//3.在main中 从ch2取值打印出来

func channelDemo() {
	var ch1, ch2 chan int
	ch1 = make(chan int)
	ch2 = make(chan int, 10)
	//wg.Add(1)
	go func() {
		//defer wg.Done()
		for i := 0; i < 10; i++ {
			x := <-ch1
			x = x * x
			ch2 <- x
		}
		close(ch1)
		close(ch2)

	}()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}

	for i := 0; i < 10; i++ {
		y := <-ch2
		fmt.Println(y)
	}
	//wg.Wait()
}

//单项通道  规定通道只能用来send 或者receive;一般用于函数参数的传递，来保证通道的操作唯一
func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
