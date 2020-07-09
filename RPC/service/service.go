package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*-----------------------------------------
RPC：
1.远程过程调用（remote procedure call）是一个计算机通信协议
2.该协议允许运行于一台计算机的程序调用另一台计算机的子程序，而程序员无需额外的为这个交互作用编程
3.RPC主要是使用在微服务中的不同微服务之间的信息传递
golang中使用RPC：
1.golang官方提供的net/rpc库，使用的是encodeing/gob进行编解码；支持Tcp、http数据传输方式，由于
其他语言不支持gob编解码的方式，所以golang中的RPC只支持golang开发的服务器与客户端之间的交互
2.官方还提供了net/rpc/jsonrpc库实现RPC方法，使用json进行数据的编解码，所以支持跨语言的调用；但是
jsonrpc是基于tcp协议实现的，咋不支持http传输方式
golang中的rpc必须符合的4个条件
1.结构体首字母大写
2.函数首字母大写
3.函数传输的参数第一个是接收的参数，第二个是参数是返回给客户端的参数且必须是指针类型
4.必须是返回一个error
-----------------------------------------*/

//以计算矩形面积和周长为例

type Rect struct {

}


type Params struct {
	Width int
	Height int
}
//矩形面积
func (r *Rect)Area(params Params,resp *int)(error){
	*resp = params.Width*params.Height
	return nil
}
//矩形周长
func (r *Rect)Perimeter(params Params,resp *int)(error){
	*resp=(params.Width+params.Height)*2
	return nil
}
//net/rpc
//func main(){
//	//1.注册服务
//	rect:=new(Rect)
//	err := rpc.Register(rect)
//	if err != nil {
//		return
//	}
//	//2.把服务处理绑定到http协议中
//	rpc.HandleHTTP()
//	//3.监听服务等待客户端调用
//	err = http.ListenAndServe(":8002", nil)
//	if err != nil {
//		return
//	}
//}

//net/rpc/jsonrpc
func main(){
	//注册服务
	rect:=new(Rect)
	err := rpc.Register(rect)
	if err != nil {
		log.Fatal(err)
	}
	//监听服务
	listen, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		log.Fatal(err)
	}
	//循环监听
	for  {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn){
			fmt.Println("new a client")
			jsonrpc.ServeConn(conn)

		}(conn)
	}
	
}