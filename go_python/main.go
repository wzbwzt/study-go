package main

import (
	"fmt"

	"github.com/sbinet/go-python"
)

//Golang是静态语言，性能很好，当它不那么灵活，不好在运行时动态运行代码。Python是动态语言，非常灵活，
//但是性能很差。但是如今有了Go-Python，鱼和熊掌也可以兼得

//go中调用python
func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	gostr := "foo"
	pystr := python.PyString_FromString(gostr)
	str := python.PyString_AsString(pystr)
	fmt.Println("hello [", str, "]")
}
