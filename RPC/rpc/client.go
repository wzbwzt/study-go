package rpc

import (
	"log"
	"net"
	"reflect"
)

//rpc客户端

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn)*Client{
	return &Client{
		conn: conn,
	}
}

//实现通用的RPC客户端
//绑定RPC访问的方法
//传入的访问的函数名
//函数的具体实现在服务端；客户端只有函数原型
//使用makefunc()实现原型到函数的调用
//fPtr 指向函数原型 //xxx.callRPC("queryUser",&query)
func (c *Client)callRPC(rpcName string,fPtr interface{}){
	//通过反射获取fPtr未初始化的函数原型
	fn := reflect.ValueOf(fPtr).Elem()
	//另一个函数作用是对第一个函数参数进行操作
	//完成与server的交互
	f:=func (args []reflect.Value)[]reflect.Value{
		//处理输入的参数
		inArgs:=make([]interface{},0,len(args))
		for _,arg:=range args{
			inArgs=append(inArgs,arg)
		}
		//创建连接
		cliSession:=NewSession(c.conn)
		//编码数据
		cdata := RPCdata{FuncName: rpcName, Args: inArgs}
		encode, err := Encode(cdata)
		if err != nil {
			log.Fatal("client encode failed;err:",err)
		}
		err = cliSession.Write(encode)
		if err != nil {
			log.Fatal("client write to conn failed;err:",err)
		}
		//读取响应数据
		read, err := cliSession.Read()
		if err != nil {
			log.Fatal("client read from conn failed;err:",err)
		}
		decode, err := Decode(read)
		if err != nil {
			log.Fatal("client decode failed;err:",err)
		}
		outArgs:=make([]reflect.Value,0,len(decode.Args))
		for i,arg:=range decode.Args{
			//必须进行nil转换
			if arg==nil{
				outArgs=append(outArgs,reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs=append(outArgs,arg.(reflect.Value))
		}
		return outArgs
	}

	//参数1：一个未初始化函数的方法值；类型是reflect.Type
	//参数2：另一个函数，作用是对第一个函数参数操作
	//返回reflect.value类型
	//makefunc使用传入的函数原型，创建一个绑定参数2的新函数
	v := reflect.MakeFunc(fn.Type(), f)
	//为函数fPtr赋值
	fn.Set(v)
}

