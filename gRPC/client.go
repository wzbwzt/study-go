package main

import (
	"context"
	"fmt"
	"github.com/wzbwzt/studyGo/gRPC/proto"
	"google.golang.org/grpc"
	"log"
)

func main(){
	//1.创建与gRPC服务端的连接 //grpc.WithInsecure()表示安全的连接
	conn, err := grpc.Dial("127.0.0.1:8002", grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc.dial failed;err:",err)
	}
	defer conn.Close()
	//2.实例化gRPC客户端
	client := proto.NewUserInfoServiceClient(conn)
	//3.组装参数
	req:=new(proto.UserRequest)
	req.Name="zs"
	//4.调用接口
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatal("client.GetUserInfo failed err:",err)
	}
	fmt.Println(response)
}
