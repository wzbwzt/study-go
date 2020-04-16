package main

//导入的包名需要使用双引号包裹起来。
// 包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔。
// Go语言中禁止循环导入包。
//如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包:如import _ "包的路径"

import (
	"fmt"

	calcs "github.com/wzbwzt/studyGo/calc"
)

func init() {
	fmt.Println("自动执行的函数")
}

func main() {
	ret := calcs.Add(1, 2)
	fmt.Println(ret)
}
