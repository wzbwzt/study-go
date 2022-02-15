package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `json:"name" key:"value1"`
	Age  uint   `json:"age" key:"value2"`
}

func main() {
	stu1 := student{
		Name: "Alice",
		Age:  123,
	}

	t := reflect.TypeOf(stu1)
	v := reflect.ValueOf(stu1)
	fmt.Println(stu1)
	fmt.Println(t.Name(), t.Kind(), v, t)
	//通过for循环来遍历结构体中的所有信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name, field.Index, field.Type, field.Tag.Get("key"))
	}
	//通过指定字段名来获取结构体信息
	if age, ok := t.FieldByName("Age"); ok {
		fmt.Println(age.Name, age.Index, age.Type, age.Tag.Get("json"))
		fmt.Println("true")
	}

}
