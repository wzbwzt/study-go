package main

import (
	"fmt"
	"sync"
)

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex //互斥锁  确保每次同时只有一个goroutine可以访问共享资源；
	//读写互斥锁  确保对共享资源进行写操作时，只会有一个goroutine可以访问，其他goroutine需要等待
	//但是当不同的goroutine 对共享资源都进行读时不会产生锁；读写锁非常适合读多写少的场景
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
