package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"go_micro_demo/hello"
	"log"
)

//定义服务端约定的接口
type Hello struct {
}

var he Hello

func (h *Hello) Info(ctx context.Context, req *hello.InfoRequest, resp *hello.InfoResponse) error {
	resp.Msg = "hello " + req.Name
	return nil

}

//proto生成go文件：protoc -I . --micro_out=. --go_out=. ./hello.proto

//cmd "micro web"  开启web服务可以使用127.0.0.1:8082 浏览器访问
func main() {
	registry := consul.NewRegistry(
		registry.Addrs("192.168.241.129:8500"),
		)
	//1.得到微服务实例
	se := micro.NewService(
		//设置微服务的名字
		micro.Name("hello.service"),
		//定义micro的版本
		micro.Version("latest"),
		//注册组件改为consul,默认mdns
		micro.Registry(registry),
		//指定端口
		micro.Address(":8002"),

	)
	//2.初始化
	se.Init()
	//3.服务注册
	err := hello.RegisterHelloInfoHandler(se.Server(), &he)
	if err != nil {
		log.Fatal("registry failed err:", err)
	}
	//4.启动微服务
	err = se.Run()
	if err != nil {
		log.Fatal("server run failed; err:", err)
	}

}
