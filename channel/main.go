package main

import (
	"fmt"
	"sync"
)

//Golang的Channel, 发送一个数据到Channel和从Channel接收一个数据都是原子性的。

/*
Channel是异步进行的, channel存在3种状态：

nil，未初始化的状态，只进行了声明，或者手动赋值为nil
active，正常的channel，可读或者可写
closed，已关闭，千万不要误认为关闭channel后，channel的值是nil
*/

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

/*
1. 对于无缓冲区的channel，往channel发送数据和从channel接收数据都会阻塞。
2.对于nil channel和有缓冲区的channel

channel	nil	空的	非空非满	满了
--------------------------------------
发送数据	阻塞	  ok	  ok		阻塞
--------------------------------------
接收数据	阻塞	  阻塞   ok		 ok
--------------------------------------
close()    panic   ok    ok        ok
--------------------------------------

channel被关闭后：
往被关闭的channel发送数据会触发panic。
从被关闭的channel接收数据，会先读完channel里的数据。如果数据读完了，
继续从channel读数据会拿到channel里存储的元素类型的零值。
data, ok := <- c
对于上面的代码，如果channel c关闭了，继续从c里读数据，当c里还有数据时，data就是对应读到的值，
ok的值是true。如果c的数据已经读完了，那data就是零值，ok的值是false。
channel被关闭后，如果再次关闭，会引发panic。
*/

/*
chan 作为参数传递时，不完全是引用传递
*/

var c chan int
var wg sync.WaitGroup

func main() {
	// noBufChan()
	// bufChan()
	//channelDemo()
	// channelDemo2()
	// test()
	test1()

}
func test() {
	defer fmt.Println(123)
	defer fmt.Println(456)
	//panic(func()int{
	// n,_:=fmt.Println("done")
	// return  n
	//}())
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

// channel 练习
func channelDemo2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

//证明chan 不完全是引用传递
func test1() {
	var c chan int

	initchan(c)

	if c == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is not nil", c)
	}
}

func initchan(c chan int) {
	c = make(chan int)
}
