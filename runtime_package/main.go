package main

import (
	"fmt"
	"path"
	"runtime"
)

//runtime.Caller()传的参数是 int;表示调用层级；如果是0表示在
//在文件中原始调用位置的具体行数及其对应的func,此时的line也是调用的位置；
//如果是1，则表示调用的上一层级所在func 及所在行数；

func f2() {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Call() failed\n")
		return
	}

	fmt.Println(file)
	fileBase := path.Base(file)
	fmt.Println(fileBase)
	fmt.Println(line)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
}

func f1() {
	// pc, file, line, ok := runtime.Caller(1)
	// if !ok {
	// 	fmt.Println("runtime.Call() failed\n")
	// 	return
	// }
	// fmt.Println(file)
	// fmt.Println(line) //35
	// funcName := runtime.FuncForPC(pc).Name()
	// fmt.Println(funcName) //main.main
	f2()
}

func main() {
	// pc, file, line, ok := runtime.Caller(0)
	// if !ok {
	// 	fmt.Println("runtime.Call() failed\n")
	// 	return
	// }
	// fmt.Println(file)
	// fmt.Println(line)
	// funcName := runtime.FuncForPC(pc).Name()
	// fmt.Println(funcName)

	// f1()
	f1()
}
