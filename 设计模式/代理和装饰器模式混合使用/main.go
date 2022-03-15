package main

import "fmt"

type LoginFunc func(string, string)

type UserService struct{}

func (this *UserService) Login(account, passwd string) {
	println("login success")
}

//==============================================================================
type UserServiceProxy struct {
	svc *UserService
}

//装饰器
type LogDecorator func(LoginFunc) LoginFunc

func LogToRedis(f LoginFunc) LoginFunc {
	return func(s1, s2 string) {
		fmt.Println("log to redis")
		f(s1, s2)
	}
}

func LogToMysql(f LoginFunc) LoginFunc {
	return func(s1, s2 string) {
		fmt.Println("log to mysql")
		f(s1, s2)
	}
}

//代理服务
func (this *UserServiceProxy) Login(logdecorator LogDecorator) LoginFunc {
	//添加的业务功能,可能是redis或者mysql
	return logdecorator(this.svc.Login)
}
func NewUserSeriviceProxy(svc *UserService) *UserServiceProxy {
	return &UserServiceProxy{svc: svc}
}

func main() {
	user := &UserService{}
	proxyUser := NewUserSeriviceProxy(user)
	proxyUser.Login(LogToRedis)("joel", "123123")

}
