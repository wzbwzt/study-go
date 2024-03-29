# 协程和线程和进程的区别

- 进程

进程是程序的一次执行过程，是程序在执行过程中的分配和管理资源的基本单位，每个进程都有自己的地址空间,进程是系统进行资源分配和调度的一个独立单位。

每个进程都有自己的独立内存空间，不同进程通过 IPC（Inter-Process Communication）进程间通信来通信。由于进程比较重量，占据独立的内存，所以上下文进程间的切换开销（栈、寄存器、虚拟内存、文件句柄等）比较大，但相对比较稳定安全。

进程可以使用以下两种方式相互通信：共享内存和消息解析

> - 共享内存:
>   有两个进程：生产者和消费者。 生产者往该区域存入信息，消费者消费该信息。 这两个进程共享一个被称为缓冲区的公共空间或内存位置，生产者将信息存储在该缓冲区中，消费者在需要时从该缓冲区中消费消息。 这个问题有两个版本：第一个被称为无界缓冲区问题，其中生产者可以一直往该缓冲区存储信息并且缓冲区大小没有限制，第二个被称为有界缓冲区问题，其中生产者最多可以存储一定数量的信息，然后它开始等待消费者消费它。
> - 消息协议方法:1.建立通信链路（如果链路已经存在，则无需重新建立。） 2.使用基本原语开始交换消息。 我们至少需要两个原语：发送（消息，目的地）或发送（消息）和接收（消息，主机）或接收（消息）

- 线程

线程是进程的一个实体,线程是内核态,而且是 CPU 调度和分派的基本单位,它是比进程更小的能独立运行的基本单位.线程自己基本上不拥有系统资源,只拥有一点在运行中必不可少的资源(如程序计数器,一组寄存器和栈),但是它可与同属一个进程的其他的线程共享进程所拥有的全部资源。

线程间通信主要通过共享内存，上下文切换很快，资源开销较少，但相比进程不够稳定容易丢失数据。

- 协程

协程是一种用户态的轻量级线程，协程的调度完全由用户控制。协程拥有自己的寄存器上下文和栈。

协程调度切换时，将寄存器上下文和栈保存到其他地方，在切回来的时候，恢复先前保存的寄存器上下文和栈，直接操作栈则基本没有内核切换的开销，可以不加锁的访问全局变量，所以上下文的切换非常快。

### 进程、线程、协程的关系和区别：

- 进程拥有自己独立的堆和栈，既不共享堆，亦不共享栈，进程由操作系统调度。

- 线程拥有自己独立的栈和共享的堆，共享堆，不共享栈，线程亦由操作系统调度(标准线程是的)。

- 协程和线程一样共享堆，不共享栈，协程由程序开发者在协程的代码里显示调度。

---

---

# 怎么限制 Goroutine 的数量

在 Golang 中，Goroutine 虽然很好，但是数量太多了，往往会带来很多麻烦，比如耗尽系统资源导致程序崩溃，或者 CPU 使用率过高导致系统忙不过来。

所以我们可以限制下 Goroutine 的数量,这样就需要在每一次执行 go 之前判断 goroutine 的数量，如果数量超了，就要阻塞 go 的执行。

所以通常我们第一时间想到的就是使用通道。每次执行的 go 之前向通道写入值，直到通道满的时候就阻塞了，

```go
package main

import "fmt"

var ch chan  int

func elegance(){
	<-ch
	fmt.Println("the ch value receive",ch)
}

func main(){
	ch = make(chan int,5)
	for i:=0;i<10;i++{
		ch <-1
		fmt.Println("the ch value send",ch)
		go elegance()
		fmt.Println("the result i",i)
	}

}
```

运行:

```go
> go run goroutine.go
the ch value send 0xc00009c000
the result i 0
the ch value send 0xc00009c000
the result i 1
the ch value send 0xc00009c000
the result i 2
the ch value send 0xc00009c000
the result i 3
the ch value send 0xc00009c000
the result i 4
the ch value send 0xc00009c000
the result i 5
the ch value send 0xc00009c000
the ch value receive 0xc00009c000
the result i 6
the ch value receive 0xc00009c000
the ch value send 0xc00009c000
the result i 7
the ch value send 0xc00009c000
the result i 8
the ch value send 0xc00009c000
the result i 9
the ch value send 0xc00009c000
the ch value receive 0xc00009c000
the ch value receive 0xc00009c000
the ch value receive 0xc00009c000
the result i 10
the ch value send 0xc00009c000
the result i 11
the ch value send 0xc00009c000
the result i 12
the ch value send 0xc00009c000
the result i 13
the ch value send 0xc00009c000
the ch value receive 0xc00009c000
the ch value receive 0xc00009c000
the ch value receive 0xc00009c000
the ch value receive 0xc00009c000
the result i 14
the ch value receive 0xc00009c000
```

```
> go run goroutine.go
the ch value send 0xc00007e000
the result i 0
the ch value send 0xc00007e000
the result i 1
the ch value send 0xc00007e000
the result i 2
the ch value send 0xc00007e000
the result i 3
the ch value send 0xc00007e000
the ch value receive 0xc00007e000
the result i 4
the ch value send 0xc00007e000
the ch value receive 0xc00007e000
the result i 5
the ch value send 0xc00007e000
the ch value receive 0xc00007e000
the result i 6
the ch value send 0xc00007e000
the result i 7
the ch value send 0xc00007e000
the ch value receive 0xc00007e000
the ch value receive 0xc00007e000
the ch value receive 0xc00007e000
the result i 8
the ch value send 0xc00007e000
the result i 9
```

这样每次同时运行的 goroutine 就被限制为 5 个了。但是新的问题于是就出现了，因为并不是所有的 goroutine 都执行完了，在 main 函数退出之后，还有一些 goroutine 没有执行完就被强制结束了。这个时候我们就需要用到 sync.WaitGroup。使用 WaitGroup 等待所有的 goroutine 退出。

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Pool Goroutine Pool
type Pool struct {
	queue chan int
	wg *sync.WaitGroup
}

// New 新建一个协程池
func NewPool(size int) *Pool{
	if size <=0{
		size = 1
	}
	return &Pool{
		queue:make(chan int,size),
		wg:&sync.WaitGroup{},
	}
}

// Add 新增一个执行
func (p *Pool)Add(delta int){
	// delta为正数就添加
	for i :=0;i<delta;i++{
		p.queue <-1
	}
	// delta为负数就减少
	for i:=0;i>delta;i--{
		<-p.queue
	}
	p.wg.Add(delta)
}

// Done 执行完成减一
func (p *Pool) Done(){
	<-p.queue
	p.wg.Done()
}

// Wait 等待Goroutine执行完毕
func (p *Pool) Wait(){
	p.wg.Wait()
}

func main(){
	// 这里限制5个并发
	pool := NewPool(5)
	fmt.Println("the NumGoroutine begin is:",runtime.NumGoroutine())
	for i:=0;i<20;i++{
		pool.Add(1)
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("the NumGoroutine continue is:",runtime.NumGoroutine())
			pool.Done()
		}(i)
	}
	pool.Wait()
	fmt.Println("the NumGoroutine done is:",runtime.NumGoroutine())
}
```

运行:

```go
the NumGoroutine begin is: 1
the NumGoroutine continue is: 6
the NumGoroutine continue is: 7
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 6
the NumGoroutine continue is: 3
the NumGoroutine continue is: 2
the NumGoroutine done is: 1
```

其中，Go 的`GOMAXPROCS`默认值已经设置为 CPU 的核数， 这里允许我们的 Go 程序充分使用机器的每一个 CPU,最大程度的提高我们程序的并发性能。`runtime.NumGoroutine`函数在被调用后，会返回系统中的处于特定状态的 Goroutine 的数量。这里的特指是指`Grunnable\Gruning\Gsyscall\Gwaition`。处于这些状态的 Groutine 即被看做是活跃的或者说正在被调度。

这里需要注意下：垃圾回收所在 Groutine 的状态也处于这个范围内的话，也会被纳入该计数器。

---

---

# 线程有几种模型？Goroutine 的原理你了解过吗，讲一下实现和原理

## 线程模型

1. 内核线程模型
1. 用户级线程模型
1. 混合型线程模型

### Linux 历史上线程的 3 种实现模型： 线程的实现曾有 3 种模型：

多对一(M:1)的用户级线程模型
一对一(1:1)的内核级线程模型
多对多(M:N)的两级线程模型

## goroutine 的原理

基于 CSP（communicating sequential processes)并发模型开发了 GMP 调度器，其中

- G（Goroutine） : 每个 Goroutine 对应一个 G 结构体，G 存储 Goroutine 的运行堆栈、状态以及任务函数
- M（Machine）: 对 OS 内核级线程的封装，数量对应真实的 CPU 数(真正干活的对象).
- P (Processor): 逻辑处理器,即为 G 和 M 的调度对象，用来调度 G 和 M 之间的关联关系，其数量可通过 GOMAXPROCS()来设置，默认为核心数。

<p align="center">
<img width="800" align="center" src="../images/198.jpeg" />
</p>

### 调度流程:

- 每个 P 有个局部队列，局部队列保存待执行的 goroutine（流程 2），当 M 绑定的 P 的的局部队列已经满了之后就会把 goroutine 放到全局队列（流 程 2-1）

- 每个 P 和一个 M 绑定，M 是真正的执行 P 中 goroutine 的实体（流程 3）， M 从绑定的 P 中的局部队列获取 G 来执行

- 当 M 绑定的 P 的局部队列为空时，M 会从全局队列获取到本地队列来执行 G （流程 3.1），当从全局队列中没有获取到可执行的 G 时候，M 会从其他 P 的局部队列中偷取 G 来执行（流程 3.2），这种从其他 P 偷的方式称为 work stealing

- 当 G 因系统调用（syscall）阻塞时会阻塞 M，此时 P 会和 M 解绑即 hand off，并寻找新的 idle 的 M，若没有 idle 的 M 就会新建一个 M（流程 5.1）

- 当 G 因 channel 或者 network I/O 阻塞时，不会阻塞 M，M 会寻找其他 runnable 的 G；当阻塞的 G 恢复后会重新进入 runnable 进入 P 队列等待执 行（流程 5.3）

## GMP 中 work stealing 机制

获取 P 本地队列，当从绑定 P 本地 runq 上找不到可执行的 g，尝试从全局链 表中拿，再拿不到从 netpoll 和事件池里拿，最后会从别的 P 里偷任务。P 此时去唤醒一个 M。P 继续执行其它的程序。M 寻找是否有空闲的 P，如果有则 将该 G 对象移动到它本身。接下来 M 执行一个调度循环（调用 G 对象->执行-> 清理线程 → 继续找新的 Goroutine 执行）

## GMP 中 hand off 机制

当本线程 M 因为 G 进行的系统调用阻塞时，线程释放绑定的 P，把 P 转移给其 他空闲的 M 执行。 细节：当发生上线文切换时，需要对执行现场进行保护，以便下次被调度执行 时进行现场恢复。Go 调度器 M 的栈保存在 G 对象上，只需要将 M 所需要的寄存 器（SP、PC 等）保存到 G 对象上就可以实现现场保护。当这些寄存器数据被保 护起来，就随时可以做上下文切换了，在中断之前把现场保存起来。如果此时 G 任务还没有执行完，M 可以将任务重新丢到 P 的任务队列，等待下一次被调度 执行。当再次被调度执行时，M 通过访问 G 的 vdsoSP、vdsoPC 寄存器进行现场 恢复（从上次中断位置继续执行）。

## GMP 调度过程中存在哪些阻塞

- I/O，select
- block on syscall
- channel
- 等待锁
- runtime.Gosched()

## Goroutine 的优势

- **上下文切换代价小**：从 GMP 调度器可以看出，避免了用户态和内核态线程切换，所以上下文切换代价小
- **内存占用少**：线程栈空间通常是 2M，Goroutine 栈空间最小 2K；

# 在 GPM 调度模型，goroutine 有哪几种状态？线程呢？

有 9 种状态

- \_Gidle：刚刚被分配并且还没有被初始化
- \_Grunnable：没有执行代码，没有栈的所有权，存储在运行队列中
- \_Grunning：可以执行代码，拥有栈的所有权，被赋予了内核线程 M 和处理器 P
- \_Gsyscall：正在执行系统调用，拥有栈的所有权，没有执行用户代码，被赋予了内核线程 M 但是不在运行队列上
- \_Gwaiting：由于运行时而被阻塞，没有执行用户代码并且不在运行队列上，但是可能存在于 Channel 的等待队列上
- \_Gdead：没有被使用，没有执行代码，可能有分配的栈
- \_Gcopystack：栈正在被拷贝，没有执行代码，不在运行队列上
- \_Gpreempted：由于抢占而被阻塞，没有执行用户代码并且不在运行队列上，等待唤醒
- \_Gscan：GC 正在扫描栈空间，没有执行代码，可以与其他状态同时存在

# 如果 goroutine 一直占用资源怎么办，GMP 模型怎么解决这个问题

如果有一个 goroutine 一直占用资源的话，GMP 模型会从正常模式转为饥饿模式，通过信号协作强制处理在最前的 goroutine 去分配使用

# 如果若干个线程发生 OOM，会发生什么？Goroutine 中内存泄漏的发现与排查？项目出现过 OOM 吗，怎么解决

线程:
如果线程发生 OOM，也就是内存溢出，发生 OOM 的线程会被 kill 掉，其它线程不受影响。

## Goroutine 中内存泄漏的发现与排查

go 中的内存泄漏一般都是 goroutine 泄露，就是 goroutine 没有被关闭，或者没有添加超时控制，让 goroutine 一只处于阻塞状态，不能被 GC。

场景
在 Go 中内存泄露分为暂时性内存泄露和永久性内存泄露

暂时性内存泄露

- 获取长字符串中的一段导致长字符串未释放
- 获取长 slice 中的一段导致长 slice 未释放
- 在长 slice 新建 slice 导致泄漏

string 相比切片少了一个容量的 cap 字段，可以把 string 当成一个只读的切片类型。获取长 string 或者切片中的一段内容，由于新生成的对象和老的 string 或者切片共用一个内存空间，会导致老的 string 和切片资源暂时得不到释放，造成短暂的内存泄漏

永久性内存泄露

- goroutine 永久阻塞而导致泄漏
- time.Ticker 未关闭导致泄漏
- 不正确使用 Finalizer 导致泄漏

## **使用 pprof 排查**
