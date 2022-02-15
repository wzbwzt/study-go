package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。
//针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操
//作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。
//Go语言中原子操作由内置的标准库sync/atomic提供。

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func main() {
	for i := 0; i < 200; i++ {
		wg.Add(1)
		//go add()
		go atoAdd()
	}
	wg.Wait()
	fmt.Println(x)
}

//加锁
func add() {
	defer wg.Done()
	lock.Lock()
	x++
	lock.Unlock()
}

//atomic  效能更高
func atoAdd() {
	defer wg.Done()
	atomic.AddInt64(&x, 1)
}
