package main

import (
	"math/rand"
	"sync"
	"time"
)

//无键入统计协程运算总时间

type MyWg struct {
	sync.WaitGroup
	t int
}

func main() {
	start := time.Now()
	var mWg MyWg

	for i := 0; i < 5; i++ {
		mWg.Add(1)
		go func() {
			defer mWg.Done(time.Now())
			dojob()
		}()
	}
	mWg.Wait()
	println("主协程执行了", time.Since(start).Milliseconds()/1000, "S")
	println("协程一共执行了", mWg.t/1000, "s")
}

func dojob() {
	t := time.Duration(rand.Int31n(3))
	time.Sleep(time.Second * t)
	println("运行：", t, "S")
}

func (m *MyWg) Done(t time.Time) {
	defer m.WaitGroup.Done()
	defer func() {
		m.t += int(time.Since(t).Milliseconds())
	}()

}
