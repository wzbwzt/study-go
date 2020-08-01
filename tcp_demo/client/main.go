package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
	defer conn.Close()
	//子协程从读取用户输入，往conn中写
	go func(){
		//fmt.Print("input:")
		for  {
			//fmt.Print("input:")
			newReader := bufio.NewReader(os.Stdin)
			msg, err := newReader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			_, err = conn.Write([]byte(msg))
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	//从conn中读取数据输出到控制台
	var tmp [128]byte
	for  {
		n, err := conn.Read(tmp[:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("收到消息：",string(tmp[:n]))
	}
}
