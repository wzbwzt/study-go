package main

import "fmt"

//用结构体模拟实现其他语言中的“继承”
type animal struct {
	name string
}

//给animal实现一个move方法
func (a animal) move() {
	fmt.Printf("%s:在动。。。\n", a.name)
}

type dog struct {
	feet   uint8
	animal //animal拥有的方法，dog也会拥有
}

//给dog实现一个jiao方法
func (d dog) jiao() {
	fmt.Printf("%s的叫声是“汪汪。。。”\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "gaiya",
		},
	}
	fmt.Println(d1)
	d1.jiao()
	d1.move()
}
