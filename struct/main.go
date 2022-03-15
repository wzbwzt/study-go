package main

//结构体:值类型；占用一块连续的内存空间

import "fmt"

//自定义类型 (区别于类型的别名设置)打印出的类型是main.person
/*
匿名结构体：struct{}
实例化的话直接：struct{}{}
因为不占内存；所以匿名空的结构体一般用来做通知使用

*/

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

//类型的别名设置（打印出的类型还是int）
type myint = int

func main() {
	var person1 person
	person1.name = "铁柱"
	person1.age = 3000
	person1.gender = "男"
	person1.hobby = []string{
		"更好",
		"优秀",
		"up",
	}
	fmt.Println(person1)
	fmt.Printf("%T\n", person1)

	//匿名结构体 （多用于临时场景）
	var s struct {
		name string
		age  int
	}
	s.name = "铁蛋"
	s.age = 1000
	fmt.Printf("%T %v\n", s, s)
	//结构体指针1
	var p1 = new(person)
	p1.name = "gaiya" //语法糖果  相当于 （*p1）.name="gaiya"
	fmt.Printf("%T %p\n", p1, p1)
	//结构体指针2
	//key-value 初始化
	var p2 = person{
		name: "tiezhu",
		age:  2000,
	}
	fmt.Printf("%#v\n", p2)
	//使用值列表的形式初始化  但是字段顺序必须要和声明的结构体字段顺序一致  且字段不能少必须全部写上
	p3 := person{
		"tiedan",
		1000,
		"男",
		[]string{
			"up", "fight",
		},
	}
	fmt.Println(p3)
}
