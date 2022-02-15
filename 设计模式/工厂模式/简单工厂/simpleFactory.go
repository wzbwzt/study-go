package main

//简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类（接口）
//一个工厂
type UserType int

const (
	ADMIN UserType = iota
	MEMBER
)

type admin struct {
	name string
	age  int
}

type member struct {
	name string
	age  int
}

type UserCreateFunc func(name string, age int) interface{}

//main
func main() {
	amdin := CreateUser(ADMIN)("wzb", 18).(*admin)
	println(amdin.name)

	member := CreateUser(MEMBER)("awu", 18).(*member)
	println(member.name)

}

func NewAdmin() UserCreateFunc {
	return func(name string, age int) interface{} {
		return &admin{name: name, age: age}
	}
}

func NewMember() UserCreateFunc {
	return func(name string, age int) interface{} {
		return &member{name: name, age: age}
	}
}

func CreateUser(usertype UserType) UserCreateFunc {
	switch usertype {
	case ADMIN:
		return NewAdmin()
	case MEMBER:
		return NewMember()
	}
	return nil
}
