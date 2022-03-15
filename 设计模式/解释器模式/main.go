package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//解释器模式：定义一个"语法或者文法",并定义一个语法解释器来处理这个方法
//场景：
//规则引擎
//验证判断
//统计模板，自定义语法
//sql解析
//业务中的简单的过滤器

//##############################################################################
type IFilter interface {
	Filter(*User) bool
}
type User struct {
	ID   int64
	Name string
	Age  int
}

type UserFilter struct {
	Expr IFilter
}

func NewUserFilter(rule string) *UserFilter {
	list := regexp.MustCompile("\\s+").Split(rule, -1)
	if len(list) < 3 {
		panic("error rule")
	}
	field := list[0]
	operator := list[1]
	value := list[2]
	if IsCompare(operator) {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		return &UserFilter{
			Expr: &CompareFilter{
				operator: operator,
				value:    intValue,
				field:    field,
			},
		}
	} else {
		if operator == "like" {
			return &UserFilter{Expr: &LikeFilter{
				value: value,
				field: field,
			}}
		}
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

// //age过滤器
// type AgeFilter struct {
// 	Operator string //eg. > < >= ...
// 	Value    int
// }

// func (this *AgeFilter) Filter(u *User) bool {
// 	switch this.Operator {
// 	case "<":
// 		return u.Age < this.Value
// 	case "<=":
// 		return u.Age <= this.Value
// 	case ">":
// 		return u.Age > this.Value
// 	case ">=":
// 		return u.Age >= this.Value
// 	default:
// 		panic("error operator")
// 	}
// }

//##############################################################################

//##############################################################################
//封装比较符
type CompareFilter struct {
	operator string //操作符
	value    int    //值
	field    string //对象属性
}

var compareoperator = []string{"<", "<=", ">", ">=", "="}

func IsCompare(operator string) bool {
	for _, v := range compareoperator {
		if operator == v {
			return true
		}
	}
	return false
}

func (this *CompareFilter) Filter(user *User) bool {
	u := reflect.ValueOf(user)
	if u.Kind() == reflect.Ptr {
		u = u.Elem()
	}
	f := u.FieldByName(this.field)
	if !f.IsValid() || f.Kind() != reflect.Int {
		return false
	}
	switch this.operator {
	case ">":
		return int(f.Int()) > this.value
	case ">=":
		return int(f.Int()) >= this.value
	case "<":
		return int(f.Int()) < this.value
	case "<=":
		return int(f.Int()) <= this.value
	case "=":
		return int(f.Int()) == this.value
	}
	return false
}

//##############################################################################
//封装like
type LikeFilter struct {
	value string //值
	field string //对象属性
}

func (this *LikeFilter) Filter(user *User) bool {
	u := reflect.ValueOf(user)
	if u.Kind() == reflect.Ptr {
		u = u.Elem()
	}
	f := u.FieldByName(this.field)
	if !f.IsValid() || f.Kind() != reflect.String {
		return false
	}

	return strings.Contains(f.String(), this.value)
}

//##############################################################################

func main() {
	users := []*User{{ID: 100, Name: "joel", Age: 20}, {ID: 101, Name: "Alex", Age: 30}, {ID: 102, Name: "Jack", Age: 38}}

	fmt.Printf("%#v", NewUserFilter("Age > 20").Filter(users))
	fmt.Println()
	fmt.Printf("%#v", NewUserFilter("Name like joe").Filter(users))
}
