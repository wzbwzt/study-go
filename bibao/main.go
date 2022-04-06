//go:build ignore
// +build ignore

package main

import (
	"fmt"
)

/*
闭包定义：
对闭包来说，函数在该语言中得是一等公民。一般来说，一个函数返回另外一个函数，
这个被返回的函数可以引用外层函数的局部变量，这形成了一个闭包。
通常，闭包通过一个结构体来实现，它存储一个函数（通常是其入口地址）和一个关联的上下文环境（相当于一个符号查找表）。
但 Go 语言中，匿名函数就是一个闭包，它可以直接引用外部函数的局部变量，因为 Go 规范和 FAQ 都这么说了。

*/

/*
所谓一等公民，是指支持所有操作的实体， 这些操作通常包括作为参数传递，从函数返回，修改并分配给变量等。

比如 int 类型，它支持作为参数传递，可以从函数返回，也可以赋值给变量，因此它是一等公民。
类似的，函数是一等公民，意味着可以把函数赋值给变量或存储在数据结构中，也可以把函数作为其它函数的参数或者返回值。
具体参照文章：https://mp.weixin.qq.com/s/H3iuhkvQWonZbi7AzmokSA (函数是一等公民，这到底在说什么？)
*/

//闭包
//闭包是一个函数。这个函数包含外部作用域的一个变量
//底层原理：1.函数可以作为返回值2.函数内部查找变量的顺序，先在自己内部找 找不到再去外层找

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

type Slice []int

func main1() {
	ret := bb(f2, 1, 2) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	f1(ret)
}

//闭包延迟求值
//错误
func testF() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i) //结果拿的始终是i=2,和其地址；要想结果不一样；需要用一个变量来接受
		})
	}
	return funs
}

//正确
func testT() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}

func main() {
	funs := testT()
	for _, f := range funs {
		f()
	}
}
