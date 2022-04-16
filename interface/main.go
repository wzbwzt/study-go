//go:build ignore
// +build ignore

package main

import "fmt"

//------------------------
//接口的断言
//------------------------

//func main() {
//	var inter interface{} //定义一个空接口
//	inter = 100
//	//接口的断言
//	//1.x.(T)
//	v1, ok := inter.(int8)
//	if !ok {
//		fmt.Println("bad choice!", v1)
//	} else {
//		fmt.Println("good choice", v1)
//	}
//	//switch判断
//	switch v2 := inter.(type) {
//	case int8:
//		fmt.Println("int8", v2)
//	case int16:
//		fmt.Println("int16", v2)
//	case string:
//		fmt.Println("string", v2)
//	case int:
//		fmt.Println("int", v2)
//	default:
//		fmt.Println("no choice", v2)
//	}
//}

//------------------------
//接口的多态性
//------------------------
type Phoner interface {
	call()
}

type AppPhone struct {
}

type HuaweiPhone struct {
}

func (a *AppPhone) call() {
	fmt.Println("i am appPhone")
}
func (a *HuaweiPhone) call() {
	fmt.Println("i am huaweiPhone")
}

//框架层  基于抽象的接口来封装的 （接口本生就是指针）
func callPhone(phone Phoner) { //传入的类型是相同的，传入不同的子类，调对应的方法； 可以兼容后期的扩展，调取未来的方法
	phone.call()
}

func main() {
	callPhone(&AppPhone{})
	callPhone(&HuaweiPhone{})

}
