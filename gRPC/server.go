package main

import (
	"context"
	"github.com/wzbwzt/studyGo/gRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

//定义服务端约定的接口
type UserInfoService struct {

}
var us UserInfoService

func (u *UserInfoService)GetUserInfo(ctx context.Context,req  *proto.UserRequest) (resp *proto.UserResponse,err error){
	name:=req.Name
	//从数据库查找数据 这边模拟数据库的查找
	if name == "zs" {
		return &proto.UserResponse{
			Id: 1,
			Name: "zs",
			Age: 18,
			Hobby: []string{"cook","code","run"} ,
		},nil
	}
	return nil,nil
}



func main(){
	//1.监听
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("net.listen failed err:",err)
	}
	log.Println("127.0.0.1:8000 has listening")
	//2.实例化grpc服务端
	serv:= grpc.NewServer()
	//3.在gRPC中注册该服务
	//第二个参数要接口类型的变量
	proto.RegisterUserInfoServiceServer(serv,&us)
	//4.启动gRPC服务端
	serv.Serve(listen)
}
