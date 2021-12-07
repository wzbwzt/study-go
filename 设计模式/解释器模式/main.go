package main

import (
	"fmt"
	"regexp"
	"strconv"
)

//解释器模式：定义一个"语法或者文法",并定义一个语法解释器来处理这个方法
//场景：
//规则引擎
//验证判断
//统计模板，自定义语法
//sql解析
//业务中的简单的过滤器

type User struct {
	ID   int64
	Name string
	Age  int
}

type UserFilter struct {
	Expr IFilter
}

func NewUserFilter(rule string) *UserFilter {
	spres := regexp.MustCompile("\\s+").Split(rule, -1)
	if len(spres) < 3 {
		panic("error rule")
	}
	switch spres[0] {
	case "age":
		value, err := strconv.Atoi(spres[2])
		if err != nil {
			panic("err value")
		}
		return &UserFilter{Expr: &AgeFilter{Operator: spres[1], Value: value}}
	}

	return &UserFilter{}
}

func (this *UserFilter) Filter(users []*User) []*User {
	var res []*User
	for _, u := range users {
		if this.Expr.Filter(u) {
			res = append(res, u)
		}
	}
	return res
}

//##############################################################################
type IFilter interface {
	Filter(*User) bool
}

//age过滤器
type AgeFilter struct {
	Operator string //eg. > < >= ...
	Value    int
}

func (this *AgeFilter) Filter(u *User) bool {
	switch this.Operator {
	case "<":
		return u.Age < this.Value
	case "<=":
		return u.Age <= this.Value
	case ">":
		return u.Age > this.Value
	case ">=":
		return u.Age >= this.Value
	default:
		panic("error operator")
	}
}

//##############################################################################
func main() {
	users := []*User{{ID: 100, Name: "joel", Age: 20}, {ID: 101, Name: "Alex", Age: 30}, {ID: 102, Name: "Jack", Age: 38}}
	fmt.Printf("%#v", NewUserFilter("age > 50").Filter(users))
}
