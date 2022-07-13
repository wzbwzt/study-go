package main

import (
	"errors"
	"sync"
	"time"
)

/*
启动容易停止难！当程序要退出时，最粗暴的方法就是不管三七二十一，
main goroutine直接退出；优雅些的方式，也是*nix系统通常的作法是：
通知一下各个Goroutine要退出了，然后等待一段时间后再真正退出。
粗暴地直接退出的方式可能会导致业务数据的损坏、不完整或丢失。
等待超时的方式虽然不能完全避免“损失”，但是它给了各个goroutine
一个“挽救数据”的机会，可以尽可能地减少损失的程度。
*/

func main() {

}

type GracefullyShutdowner interface {
	Shutdown(waitTimeout time.Duration) error
}

type ShutdownerFunc func(time.Duration) error

func (f ShutdownerFunc) Shutdown(waitTimeout time.Duration) error {
	return f(waitTimeout)
}

//并发退出
func ConcurrencyShutdown(waitTimeout time.Duration, shutdowners ...GracefullyShutdowner) error {
	c := make(chan struct{})

	go func() {
		var wg sync.WaitGroup
		for _, g := range shutdowners {
			wg.Add(1)
			go func(shutdowner GracefullyShutdowner) {
				shutdowner.Shutdown(waitTimeout)
				wg.Done()
			}(g)
		}
		wg.Wait()
		c <- struct{}{}
	}()

	select {
	case <-c:
		return nil
	case <-time.After(waitTimeout):
		return errors.New("wait timeout")
	}
}

//串行退出
//串行退出的一个问题是waitTimeout的确定，因为这个超时时间是所有goroutine的退出时间之和。
//在上述代码里，我把每次的lefttime传入下一个要执行的goroutine的Shutdown方法中，
//外部select也同样使用这个left作为timeout的值。
func SequentialShutdown(waitTimeout time.Duration, shutdowners ...GracefullyShutdowner) error {
	start := time.Now()
	var left time.Duration

	for _, g := range shutdowners {
		elapsed := time.Since(start)
		left = waitTimeout - elapsed

		c := make(chan struct{})
		go func(shutdowner GracefullyShutdowner) {
			shutdowner.Shutdown(left)
			c <- struct{}{}
		}(g)

		select {
		case <-c:
			continue
		case <-time.After(left):
			return errors.New("wait timeout")
		}
	}

	return nil
}
