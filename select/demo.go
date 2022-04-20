//go:build ignore
// +build ignore

package main

import (
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		var i = 1
		for {
			i++
			ch <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for {
		select {
		case x := <-ch:
			println(x)
		case <-time.After(3 * time.Second): //无法走到这个分支，会发生内存泄漏：原因详见Notion-Journal-Go
			println(time.Now().Unix())
			return
		}
	}
}
