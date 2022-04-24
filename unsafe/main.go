package main

import (
	"fmt"
	"unsafe"
)

/*
unsafe库徘徊在“类型安全”边缘，由于它们绕过了 Golang 的内存安全原则，一般被认为使用该库是不安全的。
但是，在许多情况下， unsafe库的作用又是不可替代的，灵活地使用它们可以实现对内存的直接读写操作。
在reflect库、syscall库以及其他许多需要操作内存的开源项目中都有对它的引用。

unsafe库源码极少，只有两个类型的定义和三个方法的声明。
*/
func main() {
	//Pointer 类型
	//这个类型比较重要，它是实现定位欲读写的内存的基础。官方文档对该类型有四个重要描述：
	//（1）任何类型的指针都可以被转化为 Pointer
	//（2）Pointer 可以被转化为任何类型的指针
	//（3）uintptr 可以被转化为 Pointer
	//（4）Pointer 可以被转化为 uintptr
	i := 100
	fmt.Println(i) // 100
	p := (*int)(unsafe.Pointer(&i))
	fmt.Println(*p) // 100
	*p = 0
	fmt.Println(i)  // 0
	fmt.Println(*p) // 0

	//Sizeof 函数
	//该函数的定义如下：
	//Sizeof 函数返回变量 v 占用的内存空间的字节数，该字节数不是按照变量 v 实际占用的内存计算，
	//而是按照 v 的“ top level ”内存计算。比如，在 64 位系统中，如果变量 v 是 int 类型，
	//会返回 16，因为 v 的“ top level ”内存就是它的值使用的内存；如果变量 v 是 string 类型，
	//会返回 16，因为 v 的“ top level ”内存不是存放着实际的字符串，而是该字符串的地址；
	//如果变量 v 是 slice 类型，会返回 24，这是因为 slice 的描述符就占了 24 个字节。
	psize := unsafe.Sizeof(p)
	fmt.Println("psize:", psize)

	//该函数的定义如下：
	//func Offsetof(v ArbitraryType) uintptr
	//函数返回由 v 所指示的某结构体中的字段在该结构体中的位置偏移字节数，注意，
	//v 的表达方式必须是“ struct.filed ”形式。 举例说明，在 64 为系统中运行以下代码：
	type Datas struct {
		c0 byte
		c1 int
		c2 string
		c3 int
	}
	var d Datas
	fmt.Println(unsafe.Offsetof(d.c0)) // 0
	fmt.Println(unsafe.Offsetof(d.c1)) // 8
	fmt.Println(unsafe.Offsetof(d.c2)) // 16
	fmt.Println(unsafe.Offsetof(d.c3)) // 32

	//如果知道的结构体的起始地址和字段的偏移值，就可以直接读写内存：
	d.c3 = 13
	pp := unsafe.Pointer(&d)
	offset := unsafe.Offsetof(d.c3)
	q := (*int)(unsafe.Pointer(uintptr(pp) + offset))
	fmt.Println(*q) // 13
	*pp = 1023

	fmt.Println(d.c3) // 1013

}
