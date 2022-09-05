package main

import (
	"github.com/wzbwzt/studyGo/init/pk1"
	"github.com/wzbwzt/studyGo/init/pk2"
	_ "github.com/wzbwzt/studyGo/init/pk3"
)

/*
包级别变量的初始化先于包内init函数的执行。

一个包下可以有多个init函数，每个文件也可以有多个init 函数。

多个 init 函数按照它们的文件名顺序逐个初始化。

应用初始化时初始化工作的顺序是，从被导入的最深层包开始进行初始化，层层递出最后到main包。

不管包被导入多少次，包内的init函数只会执行一次。

应用在所有初始化工作完成后才会执行main函数。
*/
func main() {
	pk1.Get()
	pk2.Get()
}
