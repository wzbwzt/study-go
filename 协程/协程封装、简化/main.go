package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//简化协程的封装和使用

type MyDogFunc func() interface{}

type MyDog struct {
	funcs []MyDogFunc
	data  chan interface{}
	wg    *sync.WaitGroup
}

func NewMyDog() *MyDog {
	return &MyDog{data: make(chan interface{}), wg: &sync.WaitGroup{}}
}

func (m *MyDog) Set(f ...MyDogFunc) {
	m.funcs = append(m.funcs, f...)
}

func (m *MyDog) do() {
	for _, f := range m.funcs {
		m.wg.Add(1)
		go func(fu MyDogFunc) {
			defer m.wg.Done()

			out := fu()
			m.data <- out
		}(f)
	}
}

//拿到协程执行结果
func (m *MyDog) Range(f func(interface{})) {
	m.do()

	go func() {
		defer close(m.data)

		m.wg.Wait()
	}()

	for value := range m.data {
		f(value)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dog := NewMyDog()
	dog.Set(func() interface{} {
		time.Sleep(time.Second * 3)
		return rand.Int31n(10)
	})
	dog.Set(func() interface{} {
		return string(rand.Int63n(60) + 10)
	})
	dog.Range(func(i interface{}) {
		fmt.Println(i)
	})
}
