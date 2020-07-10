package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

//用户查询测试示例


type User struct {
	Name string
	Age int
}

func Query(uid int)(User,error){
	user:=make(map[int]User)
	user[0]=User{"zs",12}
	user[1]=User{"li",13}
	user[2]=User{"ww",14}
	if v,ok:= user[uid];ok{
		return v,nil
	}
	return User{},fmt.Errorf("user id :%v not exited",uid)
}

func TestRPC (t *testing.T) {
	//需要对interface{}可能产生的类型进行注册
	gob.Register(User{})
	//创建服务端
	server := NewServer("127.0.0.1:8000")
	server.Register("Query",Query)
	//服务端等待调用
	go server.Run()
	//客户端获取连接
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	client := NewClient(conn)
	//先声明一个函数原型
	var fPtr func(int)(User,error)
	client.callRPC("Query",&fPtr)
	//得到查询结果
	rUser, err := Query(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(rUser)
}