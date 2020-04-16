package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	fileObj, err := os.OpenFile("./nothistxt.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 6666) //第三个参数是8进制权限；windows上没用
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
	fileObj, err := os.OpenFile("./demo2.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 666)
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
	// writeDemo1()
	// writeDemo2()
	writeDemo3()
}
