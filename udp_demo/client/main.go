package main

import (
	"fmt"
	"net"
)

//udp 客户端
func main() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP: net.IPv4(127, 0, 0, 1),
		// IP:   net.IPv4(192, 168, 33, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("udp conn failed err:", err)
		return
	}
	defer udpConn.Close() //写在错误判断的后面
	var tmp = []byte("hello")
	_, err = udpConn.Write(tmp) //写发送信息
	if err != nil {
		fmt.Println("udp write failed err:", err)
		return
	}
	var data [1024]byte
	n, addr, err := udpConn.ReadFromUDP(data[:]) //读取服务端返回的信息
	if err != nil {
		fmt.Println("udp read from server failed err:", err)
		return
	}
	fmt.Printf("data:%v;addr:%v;count:%d", string(tmp[:n]), addr, n)

}
