package main

import "fmt"

/*
工厂模式，先创建一个接口，多态两个实例
再创建一个接口，继承另一个工厂
*/

type Interface1 interface {
	Get() string
	Set(int, string) Interface1
	Remove(int)
}

type admin struct {
	name string
	age  int
}

func (a admin) Get() string {
	return fmt.Sprintf("name:%s,age:%d", a.name, a.age)
}
func (a admin) Set(i int, str string) Interface1 {
	a.name = str
	a.age = i
	return a
}
func (a admin) Remove(id int) {
	fmt.Println("管理员用户删除成功")
}

type member struct {
	name string
	age  int
}

func (m member) Get() string {
	return fmt.Sprintf("name:%s,age:%d", m.name, m.age)
}
func (m member) Set(i int, str string) Interface1 {
	m.name = str
	m.age = i
	return m
}
func (m member) Remove(id int) {
	fmt.Println("普通用户删除成功")
}

type Users interface {
	Admin() Interface1
	Member() Interface1
	Create()
	Alter()
}

type BlogUsers struct{}

func (b BlogUsers) Admin() Interface1 {
	return &admin{}
}

func (b BlogUsers) Member() Interface1 {
	return &member{}
}

func (b BlogUsers) Create() { fmt.Println("创建一个用户") }
func (b BlogUsers) Alter()  { fmt.Println("修改一个用户") }

func main() {
	var bloguser Users = new(BlogUsers)
	admin := bloguser.Admin().Set(21, "孙海铭")
	fmt.Println(admin.Get())

}
