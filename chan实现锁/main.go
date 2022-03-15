package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name  string
	Score int
	*Mutex
}

//##############################################################################
type Mutex struct {
	sign chan struct{}
}

func (this *Mutex) Lock() {
	this.sign <- struct{}{}
}

func (this *Mutex) UnLock() {
	select {
	case <-this.sign:
	default:
		panic("chan has close")
	}
}

//##############################################################################

func (this *User) Add() {
	this.Mutex.Lock()
	defer this.Mutex.UnLock()

	this.Score++
}

func (this *User) Minus() {
	this.Mutex.Lock()
	defer this.Mutex.UnLock()

	this.Score--
}

func NewUser(name string, score int) *User {
	sign := make(chan struct{}, 1)
	return &User{Name: name, Score: score, Mutex: &Mutex{sign: sign}}

}

func main() {

	u := NewUser("joel", 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			u.Add()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			u.Minus()
		}()
	}

	wg.Wait()
	fmt.Println(u)
}
