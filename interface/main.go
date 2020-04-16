package main

import "fmt"

//接口的断言

func main() {
	var inter interface{} //定义一个空接口
	inter = 100
	//接口的断言
	//1.x.(T)
	v1, ok := inter.(int8)
	if !ok {
		fmt.Println("bad choice!", v1)
	} else {
		fmt.Println("good choice", v1)
	}
	//switch判断
	switch v2 := inter.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case string:
		fmt.Println("string", v2)
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("no choice", v2)

	}
}
