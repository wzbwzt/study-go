package main

import (
	"fmt"
)

//装饰器模型：允许向一个现有的对象添加新的功能，同事不会改变其结构
//分两个角度：
//函数角度
//struct角度(类)

type User struct {
	Name string
	Age  int
}

func GetUserInfo() *User {
	return &User{Name: "joel", Age: 20}
}

type userFunc func() *User

func GetUserInfoWithRole(f userFunc) userFunc {
	return func() *User {
		user := f()
		user.Name = "amdin-" + user.Name
		return user
	}
}
func main() {
	user := GetUserInfo()
	fmt.Printf("%v", user)

	withroleuser := GetUserInfoWithRole(GetUserInfo)()
	fmt.Printf("%v", withroleuser)
}
