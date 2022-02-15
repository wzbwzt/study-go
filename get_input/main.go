package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取终端输出时  获取到空格就会停止
func useScan() {
	var s string
	fmt.Print("请输入：")
	fmt.Scanln(&s)
	fmt.Printf("你输入的是：%v", s)
}

func useBufio() {
	res := bufio.NewReader(os.Stdin)
	fmt.Print("请输入（bufio）:")
	s, _ := res.ReadString('\n')
	fmt.Printf("输入的内容是%s", s)

}

func main() {
	// useScan()
	useBufio()
}
