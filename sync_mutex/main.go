package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x  int64
	wg sync.WaitGroup
	//互斥锁  确保每次同时只有一个goroutine可以访问共享资源；
	lock sync.Mutex
	//读写互斥锁  确保对共享资源进行写操作时，只会有一个goroutine可以访问，其他goroutine需要等待
	//但是当不同的goroutine 对共享资源都进行读时不会产生锁；读写锁非常适合读多写少的场景
	//总结：
	//多个写操作之间是互斥的.
	//写操作与读操作之间也是互斥的.
	//多个读操作之间不是互斥的.
	rwlock sync.RWMutex
)

//互斥锁

func main() {
	//竞态问题
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)

	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(x)

}
func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	defer wg.Done()
}

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	rwlock.RUnlock()

}
func write() {
	defer wg.Done()
	rwlock.Lock()
	x++
	rwlock.Unlock()
	// fmt.Println(x)
}

/*
Locker表示可以锁定和解锁的对象。

type Locker interface {
   Lock()
   Unlock()
}

//锁定当前的互斥量
//如果锁已被使用，则调用goroutine
//阻塞直到互斥锁可用。
func (m *Mutex) Lock()

// 对当前互斥量进行解锁
// 如果在进入解锁时未锁定m，则为运行时错误。
// 锁定的互斥锁与特定的goroutine无关。
// 允许一个goroutine锁定Mutex然后安排另一个goroutine来解锁它。
func (m *Mutex) Unlock()
*/

//以防忘记及时解开已被锁住的锁，从而导致流程异常。但Go由于存在defer，所以此类问题出现的概率极低
//我们知道如果遇到panic，可以使用recover方法进行恢复，但是如果对重复解锁互斥锁引发的panic却是无用的（Go 1.8及以后）。

//如果对一个已经上锁的对象再次上锁，那么就会导致该锁定操作被阻塞，直到该互斥锁回到被解锁状态.
func test1() {
	var mutex sync.Mutex
	fmt.Println("begin lock")
	mutex.Lock()
	fmt.Println("get locked")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Println("begin lock ", i)
			mutex.Lock()
			fmt.Println("get locked ", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	fmt.Println("get unlocked")
	time.Sleep(time.Second)

	/*
		 output:
			begin lock
			get locked
			begin lock  3
			begin lock  1
			begin lock  2
			Unlock the lock
			get unlocked
			get locked  3
	*/
	/*
		在for循环之前开始加锁，然后在每一次循环中创建一个协程，并对其加锁，
		但是由于之前已经加锁了，所以这个for循环中的加锁会陷入阻塞直到main中的锁被解锁
	*/

}
