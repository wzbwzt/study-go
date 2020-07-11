package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/wzbwzt/studyGo/go_micro/hello"
	"log"
)
//定义服务端约定的接口
type Hello struct {

}


var he Hello

func (h *Hello) Info(ctx context.Context,req *hello.InfoRequest, resp *hello.InfoResponse) error {
	resp.Msg="hello "+req.Name
	return nil

}
//proto生成go文件：protoc -I . --micro_out=. --go_out=. ./hello.proto

func main(){
	//1.得到微服务实例
	se := micro.NewService(
		//设置微服务的名字
		micro.Name("hello"),
	)
	//2.初始化
	se.Init()
	//3.服务注册
	err := hello.RegisterHelloInfoHandler(se.Server(), &he)
	if err != nil {
		log.Fatal("registry failed err:",err)
	}
	//4.启动微服务
	err = se.Run()
	if err != nil {
		log.Fatal("server run failed; err:",err)
	}

}