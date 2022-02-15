package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

/**
一个 Promise 对象代表一个在这个 promise 被创建出来时不一定已知的值。
它让您能够把异步操作最终的成功返回值或者失败原因和相应的处理程序关联起来。
这样使得异步方法可以像同步方法那样返回值：异步方法并不会立即返回最终的值，
而是会返回一个 promise，以便在未来某个时候把值交给使用者。
**/

type Resolve func(interface{})
type Reject func(error)
type PromiseFunc func(Resolve, Reject)
type PromiseOpt func(*Promise)
type PromiseOpts []PromiseOpt

type Promise struct {
	F       PromiseFunc
	Res     Resolve
	Rej     Reject
	Wg      sync.WaitGroup
	TimeOut time.Duration //超时处理
}

func NewPromise(f PromiseFunc) *Promise {
	return &Promise{F: f, TimeOut: time.Second * 3} //默认超时3秒
}
func (p *Promise) Then(res Resolve) *Promise {
	p.Res = res
	return p
}
func (p *Promise) Catch(rej Reject) *Promise {
	p.Rej = rej
	return p
}

func (ops PromiseOpts) apply(p *Promise) {
	for _, opt := range ops {
		opt(p)
	}
}

func (p *Promise) Done(opt ...PromiseOpt) {
	defer func() {
		if e := recover(); e != nil {
			log.Print(e)
		}

	}()
	PromiseOpts(opt).apply(p)

	ctx, _ := context.WithTimeout(context.Background(), p.TimeOut)

	p.Wg.Add(1)
	sign := make(chan struct{})
	go func() {
		defer p.Wg.Done()
		p.F(p.Res, p.Rej)
	}()
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		p.Wg.Wait()
	}()
	select {
	case <-sign:
		fmt.Printf("正常完成")
	case <-ctx.Done():
		panic("超时")
	}
}

func WithTimeOut(t time.Duration) PromiseOpt {
	return func(p *Promise) {
		p.TimeOut = t
	}
}

func main() {
	NewPromise(func(r1 Resolve, r2 Reject) {
		time.Sleep(time.Second * 3)
		if time.Now().Second()%2 == 0 {
			r1("OK")
		} else {
			r2(errors.New("failed"))
		}
	}).Then(func(i interface{}) {
		fmt.Println(i)
	}).Catch(func(e error) {
		log.Println(e)
	}).Done(WithTimeOut(time.Second * 5))

}
