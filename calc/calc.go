package calcs //包名通常和目录名一致；不能包含‘-’ ;一个文件夹就是一个包；文件夹中放的都是.go文件

import "fmt"

//包中的标识符（变量名/函数名/结构体/结构等），如果首字符小写的话，表示私有（只有在这个包里才可以调用）
//要想在别的包里调用的话首字符要大写

//在Go语言程序执行时导入包语句会自动触发包内部init()函数的调用。需要注意的是：
//init()函数没有参数也没有返回值。 init()函数在程序运行时自动被调用执行，不能在代码中主动调用它
func init() {
	fmt.Println("改包(calcs)被导入时自动执行的函数")
}

func Add(x int, y int) int {
	return x + y
}
