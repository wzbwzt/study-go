package main

import "fmt"

//代理模式指对方法的请求通过代理函数来执行，这样可以在代理函数中条件一些额外的业务操作

type UserService struct{}

func (this *UserService) Login(account, passwd string) {
	println("login success")
}

//代理服务对象
type UserServiceProxy struct {
	svc *UserService
}

func (this *UserServiceProxy) Login(account, passwd string) {
	//添加的业务功能
	fmt.Println("记录日志")

	this.svc.Login(account, passwd)
}
func NewUserSeriviceProxy(svc *UserService) *UserServiceProxy {
	return &UserServiceProxy{svc: svc}
}

func main() {
	user := &UserService{}
	proxyUser := NewUserSeriviceProxy(user)
	proxyUser.Login("joel", "123123")

}
