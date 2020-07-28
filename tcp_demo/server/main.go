package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//tcp服务端
//1.本地端口启动服务
//2.等待别人请求连接
//3.与客户端通信（为每一个客户端开辟一个独立的并发与其IO）
func main() {
	//1.
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Listen failed err:", err)
		return
	}

	for {
		//2.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed err:", err)
			return
		}
		//3.
		go processConn(conn)
	}

}

func processConn(conn net.Conn) {
	//3.
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin) //获取终端输入
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read msg failed err:", err)
			return
		}
		fmt.Println("收到消息：",string(tmp[:n]))
		fmt.Print("请回复：")
		msg, err := reader.ReadString('\n')
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("conn write to client failed,err:", err)
			return
		}
	}
}
