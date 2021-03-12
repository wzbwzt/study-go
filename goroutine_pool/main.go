package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

//goroutine pool的设计思路：
//参照博客：https://www.jianshu.com/p/fa6d82934cb8
//goroutine pool封装的代码：https://github.com/panjf2000/ants
//
//启动服务之时先初始化一个 Goroutine Pool 池，这个Pool维护了一个类似栈的LIFO队列 ，里面存放负责处理任务的Worker，然后在client端提交task
//到Pool中之后，在Pool内部，接收task之后的核心操作是：
//1.检查当前Worker队列中是否有空闲的Worker，如果有，取出执行当前的task；
//2.没有空闲Worker，判断当前在运行的Worker是否已超过该Pool的容量，是 — 阻塞等待直至有Worker被放回Pool；否 — 新开一个Worker（goroutine）
//处理；
//3.goroutine复用，每个Worker执行完任务之后，放回Pool的队列中等待。因为复用，所以规避了无脑启动大规模goroutine的弊端，可以节省大量的内存。
var (
	ErrPoolSizeInvalid = errors.New("pool size invalid")
	ErrPoolClosed      = errors.New("pool has cloesd")
)

// NewPool generates a instance of ants pool
func NewPool(size, expiry int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrPoolSizeInvalid
	}
	p := &Pool{
		capacity:       int32(size),
		freeSignal:     make(chan sig, math.MaxInt32),
		release:        make(chan sig, 1),
		expiryDuration: time.Duration(expiry) * time.Second,
	}
	// 启动定期清理过期worker任务，独立goroutine运行，
	// 进一步节省系统资源
	p.monitorAndClear()
	return p, nil
}

func main() {
	p, err := NewPool(1000, 100)
	if err != nil {
		panic(err)
	}
	p.Submit(func() error {
		i := 0
		i++
		fmt.Println(i)
		return nil
	})
}
