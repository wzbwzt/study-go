package main

import (
	"fmt"
	"strings"
	"sync"
)

//交替打印数字和字符
func f1() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		for {
			i := 1
			select {
			case <-number:
				fmt.Println(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wg.Add(1)
	go func(w *sync.WaitGroup) {
		str := "ASKJFADSFJADFJLJFLF"
		i := 0
		for {
			if i >= strings.Count(str, "")-1 {
				w.Done()
				break
			}
			select {
			case <-letter:
				fmt.Println(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}

	}(&wg)

	number <- true
	wg.Wait()
}
