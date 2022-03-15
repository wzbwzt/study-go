package main

import "fmt"

//自定义结构体
type dog struct {
	name string
	age  int
}

//构造函数
func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

/*
方法：method
方法是作用于特定类型的函数
接收者表示的是调用该方法的具体类型变量，多用类型名的首字母来作为变量
接收者必须是自己包里定义的类型，不能是内置的类型如int  如果需要可以自定义int:(type myint int)
*/
//指针类型接收者
func (d *dog) wang() { //因为go函数传的值永远是拷贝的值，所以当需要修改接收者中的值的时候，需要传指针
	d.age++ //语法糖果：相当于（*d）.age++
	fmt.Printf("%s wangwang...\n", d.name)
}

//值接收者
func (d dog) wang2() {
	fmt.Printf("%s wangwang...\n", d.name)
}

func main() {
	d1 := newDog("gaiya", 12)
	d1.wang() //默认传的就是结构体的指针
	fmt.Println(d1.age)
}
