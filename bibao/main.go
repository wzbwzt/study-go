package main

import (
	"fmt"
)

//闭包
//闭包是一个函数。这个函数包含外部作用域的一个变量
//底层原理：1.函数可以作为返回值2.函数内部查找变量的顺序，现在自己内部找 找不到再去外层找

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数f2进行封装  以达到f1的传入参数要求（f1传入的是一个没有传参和返回值的一个函数）

func bb(f func(int, int), m, n int) func() {
	tmp := func() {
		f(m, n)
	}
	return tmp
}
type  Slice []int


func main1() {
	ret := bb(f2, 1, 2) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	f1(ret)
}

//闭包延迟求值
//错误
func testF() []func() {
	var funs []func()
	for i:=0;i<2;i++  {
		funs = append(funs,func() {
			println(&i,i)      //结果拿的始终是i=2,和其地址；要想结果不一样；需要用一个变量来接受
		})
	}
	return funs
}
//正确
func testT() []func() {
	var funs []func()
	for i:=0;i<2;i++  {
		x:=i
		funs = append(funs,func() {
			println(&x,x)
		})
	}
	return funs
}

func main(){
	funs:=testT()
	for _,f:=range funs{
		f()
	}
}





