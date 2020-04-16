package main

import "fmt"

//指针：go语言中不存在指针操作只需要记住：&和*

func main() {
	//1.&：取内存地址
	n := 123
	fmt.Println(&n)
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) //*int:int类型的指针
	//2.*:根据地址取值
	m := *p
	fmt.Println(m)

	/*
		make和new 的区别
		1.new 和make 都是用来申请内存的
		2.new很少用，一般用来给记住数据类型申请内存,string 、int,返回的是对应类型的指针（*string、*int）
		3.make是用来给slice 、map、chan申请内存的，make函数返回的是对应的这三个类型本身
	*/
	a := new(int)
	b := new(string)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(*a)
	*a = 12
	fmt.Println(*a)
	c := make([]string, 12)
	fmt.Println(c)
	fmt.Printf("%T", c)

}
