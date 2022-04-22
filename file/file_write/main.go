package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/*
func (f *File) Write(b []byte) (n int, err error)
Write 调用成功并不能保证数据已经写入磁盘，因为内核会缓存磁盘的 I/O 操作。
如果希望立刻将数据写入磁盘（一般场景不建议这么做，因为会影响性能），有两种办法：
1. 打开文件时指定 `os.O_SYNC`；
2. 调用 `File.Sync()` 方法。
说明：File.Sync() 底层调用的是 fsync 系统调用，这会将数据和元数据都刷到磁盘；如果只想刷数据到磁盘
（比如，文件大小没变，只是变了文件数据），需要自己封装，调用 fdatasync 系统调用。（syscall.Fdatasync）
*/

func writeDemo1() {
	fileObj, err := os.OpenFile("./nothistxt.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666) //第三个参数是8进制权限；windows上没用
	if err != nil {
		fmt.Printf("file open failed  the err is %v", err)
		return
	}
	//write
	fileObj.Write([]byte("heiheihei---\nos.O_TRUNC 指每次打开先清空再重写入\n"))
	//WriteString
	fileObj.WriteString("hahaha+++++")
	fileObj.Close()
}

func writeDemo2() {
	fileObj, err := os.OpenFile("./demo2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file failed the err is %v", err)
		return
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("bufio write的方法\n") //写入到缓存
	wr.Flush()                         //将缓存中的文件写入到文件
}

func writeDemo3() {
	str := "ioutil.WriteFile() write function\n"
	err := ioutil.WriteFile("./demo3.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file faild the err is %v", err)
		return
	}
}
func main() {
	writeDemo1()
	// writeDemo2()
	// writeDemo3()
}
