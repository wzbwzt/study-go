package main

import (
	"fmt"
	"strconv"
	"sync"
)

//go语言中内置的map不是并发安全的
//所以sync包中提供了一个开箱即用的并发安全版map  sync.Map
//开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用

var m sync.Map
var wg sync.WaitGroup

func main() {
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         //往map中存取键值对
			value, _ := m.Load(key) //根据键从map中取值
			fmt.Println(key, value)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

}


