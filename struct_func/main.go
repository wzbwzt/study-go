package main

import "fmt"

type person struct {
	name string
	age  int
}

/*
构造函数
定义：返回一个结构体变量的函数
命名一般是以new开头的
返回的是结构体还是结构体指针：
当结构体大的时候返回的是结构体问题不大  但是当结构体大的时候，尽量使用结构体指针来减少内存的消耗
*/

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("铁柱", 3000)
	fmt.Printf("%p\n", p1)
	fmt.Println(*p1)
}
