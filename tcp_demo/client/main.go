package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp  客户端
//1.与server端建立连接
//2.发送信息
//3.
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed err :", err)
		return
	}

	reader := bufio.NewReader(os.Stdin) //获取终端输入 os.Args 也可以获取终端输入
	for {
		fmt.Print("请写入发送内容：")
		msg, err := reader.ReadString('\n')
		if msg == "exit" {
			break
		}
		msg = strings.TrimSpace(msg)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("conn write to server failed,err:", err)
			return
		}
	}
	conn.Close()
}
