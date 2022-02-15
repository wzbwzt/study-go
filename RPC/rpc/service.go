package rpc

import (
	"log"
	"net"
	"reflect"
)

//服务端的实现；主要解决的问题是：
//client调用时只传过来函数名；需要维护函数名与函数之间的map映射
//一般约定函数第二个返回值是error类型

type Server struct {
	addr string
	//函数名到函数反射值的map
	funcs map[string]reflect.Value
}

//创建服务端对象
func NewServer(addr string)*Server{
	return &Server{
		addr:addr,
		funcs:make(map[string]reflect.Value),
	}
}
//服务端绑定注册方法
//将函数名与函数真正实现对应起来
//第一个参数是传入的函数名，第二个是传入的真正的函数
func (s *Server)Register(rpcName string ,f interface{}){
	_ ,ok:= s.funcs[rpcName]
	if ok{
		return
	}
	valueOf := reflect.ValueOf(f)
	s.funcs[rpcName]=valueOf
}

//服务端等待调用
func (s *Server)Run(){
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal(err)
	}
	for  {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//从链接中读取数据
		newSession := NewSession(conn)
		readData, err := newSession.Read()
		if err != nil {
			log.Fatal("session read failed:",err)
		}
		//数据解码
		decodeData, err := Decode(readData)
		if err != nil {
			log.Fatal("decode failed:",err)
		}
		readFuncName:=decodeData.FuncName
		value,ok := s.funcs[readFuncName]
		if !ok{
			log.Fatal("func not exited")
		}
		//解析遍历客户端传来的参数，放在一个数组中
		rpcArgs := decodeData.Args
		args:=make([]reflect.Value,0,len(rpcArgs))
		for _,rpcArg:=range rpcArgs{
			args=append(args,reflect.ValueOf(rpcArg))
		}
		//反射调用方法
		out := value.Call(args)
		outArgs:=make([]interface{},0,len(out))
		for _,outArg:=range out{
			outArgs=append(outArgs,outArg)
		}
		//包装数据，返回客户端
		respRPCdata:=RPCdata{FuncName: readFuncName,Args: outArgs}
		//编码
		encode, err := Encode(respRPCdata)
		if err != nil {
			log.Fatal("encode failed err:",err)
		}
		//写出数据
		err = newSession.Write(encode)
		if err != nil {
			log.Fatal("session write failed err:",err)
		}
	}
}
