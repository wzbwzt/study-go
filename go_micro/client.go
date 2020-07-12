package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"go_micro_demo/hello"
	"log"
)

func main(){
	//1.定义服务，可以传入其他参数
	service := micro.NewService(
		micro.Name("hello.client"),
	)
	//2.初始化
	service.Init()
	//3.创建客户端
	client := hello.NewHelloInfoService("hello.service", service.Client())
	//4.调用服务
	resp, err := client.Info(context.Background(), &hello.InfoRequest{Name: "阿无的吴"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Msg)

}
