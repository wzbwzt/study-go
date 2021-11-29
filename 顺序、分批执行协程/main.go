package main

import (
	"sync"
)

//协程按照顺序执行,eg.fun1/func2/fun3,先执行fun1再执行func2orfunc3
//分批执行
//sync.Cond 实现
//借助与Lock（mutex or RWMutex）,唤醒一个或者多个线程继续执行
//常用方法:
//sync.NewCond(&Mutex):需要传入一个mutex(实现阻塞和解除阻塞通知)
//sync.Wait();等待通知
//sync.Signal():发送单个通知
//sync.Broadcast():用于广播

type jobFunc func()

func main() {
	wg := do(job1, job2, job3)
	wg.Wait()
}

func job1() {
	println("job1")
}

func job2() {
	println("job2")
}
func job3() {
	println("job3")
}

func do(f ...jobFunc) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	for _, fu := range f {
		wg.Add(1)
		go func(f jobFunc) {
			defer wg.Done()
			f()
		}(fu)
	}
	return wg
}
