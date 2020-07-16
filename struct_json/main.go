package main

import (
	"encoding/json"
	"fmt"
)

//结构体和json转换
//1.序列化：Go语言中结构体变量转换为json格式的字符串
//2.反序列化：json格式的字符串转换为Go语言中可以识别的结构体变量

type person struct {
	Name string `json:"name"` //因为结构体会传入到Marshal函数中所以字段必须是大写的才可以被别的包调用
	Age  int    `json:"age"`  //``是为了指定转换为对应格式的字段名。json："age"标识转json时该字段显示为age
}

func main1() {
	p1 := person{
		Name: "Bradley",
		Age:  18,
	}
	fmt.Printf("%v\n", p1)
	//序列化
	v, error := json.Marshal(p1)
	if error != nil {
		fmt.Printf("error is %v\n", error)
		return
	}
	fmt.Printf("%v\n", string(v))
	//反序列化
	str := `{"name":"coope","age":20}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //因为是传入函数且是修改  所以需要传入的是指针
	fmt.Printf("%#v\n", p2)
}

type Person struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person) SetDreams1(dreams []string) {
	p.dreams = dreams
}
func (p *Person) SetDreams2(dreams []string) {
	p.dreams = make([]string, len(dreams))
	p.dreams = dreams
}
func SetDreams(dreams []string) {
	dreams[0] = "dreams"
}
func main(){
	//p1 := Person{name: "小王子", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	//p1.SetDreams1(data)
	SetDreams(data)
	//data[1]="bu"
	fmt.Println(data)
}