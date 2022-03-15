package main

import (
	"fmt"
	"sync"
)

//并发对象模式：解藕方法的调用和执行，让调用和执行在不同线程（协程）中执行，使得程序具备并发执行的能力
// 譬如有个业务场景：
// 1.统计某个用户的积分
// 2.购买商品+1分
// 3.发帖被删-1分

//并发不安全的写法
//##############################################################################
//当数据操作耗时长的时候,并不是并发安全的
// type UserInfo struct {
// 	UserID int
// 	Score  int
// }

// func NewUserInfo(userid, score int) *UserInfo {
// 	return &UserInfo{UserID: userid, Score: score}

// }
// func (this *UserInfo) Add() {
// 	this.Score++
// }

// func (this *UserInfo) Minus() {
// 	this.Score--
// }

//##############################################################################

//并发对象模式写法，使用chan(并发安全的) 将调用与执行操作分开执行
//抽离出执行方法
type OperatorFunc func()
type UserInfo struct {
	UserID int
	Score  int
	ch     chan OperatorFunc
}

func (this *UserInfo) add() {
	this.Score++
}

func (this *UserInfo) Add() {
	this.ch <- this.add
}

func (this *UserInfo) minus() {
	this.Score--
}

func (this *UserInfo) Minus() {
	this.ch <- this.minus
}

func NewUserInfo(userid, score int) *UserInfo {
	ch := make(chan OperatorFunc)
	this := &UserInfo{UserID: userid, Score: score, ch: ch}
	go this.watch()
	return this
}

func (this *UserInfo) watch() {
	for f := range this.ch {
		f()
	}
}

func main() {
	wg := sync.WaitGroup{}
	user := NewUserInfo(1, 10)

	wg.Add(2)
	go func() {
		defer wg.Done()
		//模拟业务添加积分
		for i := 0; i < 100000; i++ {
			user.Add()
		}
	}()

	go func() {
		defer wg.Done()
		//模拟业务减去积分
		for i := 0; i < 100000; i++ {
			user.Minus()
		}
	}()

	wg.Wait()
	fmt.Printf("%#v", user)
}
