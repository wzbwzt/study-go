package main

import "fmt"

//结构体嵌套

type addres struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	addr addres
}
type company struct {
	name   string
	addres //匿名结构体嵌套
}

func main() {
	p1 := person{
		name: "tiezhu",
		age:  2000,
		addr: addres{
			province: "安徽",
			city:     "合肥",
		},
	}
	c1 := company{
		name: "公司",
		addres: addres{
			province: "浙江",
			city:     "杭州",
		},
	}
	fmt.Println(p1, p1.addr.city)
	fmt.Println(c1, c1.city) //只适用于匿名结构体的嵌套  会首先从自己的结构体中找字段 没有的话就去匿名结构体中寻找

}
