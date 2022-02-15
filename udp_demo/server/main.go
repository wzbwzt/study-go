package main

import (
	"fmt"
	"net"
)

//udp 服务端

func main() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(127, 0, 0, 1),
		// IP:   net.IPv4(192, 168, 33, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("udp listen failed err:", err)
		return
	}
	defer udpConn.Close()
	var tmp [1024]byte
	for {
		n, udpAddr, err := udpConn.ReadFromUDP(tmp[:]) //接收数据
		if err != nil {
			fmt.Println("readform udp failed err:", err)
			return
		}
		fmt.Printf("data:%v;addr:%v;count:%d", string(tmp[:n]), udpAddr, n)
		_, err = udpConn.WriteToUDP(tmp[:n], udpAddr) //发送数据
		if err != nil {
			fmt.Println("write udp failed err:", err)
			return
		}
	}

}
