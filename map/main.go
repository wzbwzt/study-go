package main

import "fmt"

type s struct {
	name string
}

//map因为是引用类型 所以使用一定要先初始化
func main() {
	//---------------------------
	//重点：map中的结构体无法直接寻址，必须取地
	//---------------------------
	ma := map[string]*s{"x": &s{"one"}}
	ma["x"].name = "two"
	fmt.Println(ma)

	var m1 map[string]int
	fmt.Println(m1)               //没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) //要估算好该map容量，避免在程序运行中再动态扩容
	m1["tiezhu"] = 12
	m1["tiedan"] = 3000
	fmt.Println(m1)
	fmt.Println(m1["tiezhu"])

	fmt.Println(m1["gaiya"]) //如果不存在这个Key就拿到对应类型的零值

	//约定俗成用Ok来接受返回的布尔型数据
	value, ok := m1["gaiya"]
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(value)
	}
	//遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//删除
	delete(m1, "tiezhu")
	fmt.Println(m1)

	//map和slice 组合
	//元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	//需要对内部的map做初始化
	s1[0] = make(map[int]string, 1)
	fmt.Println(s1)
	s1[0][01] = "tiezhu"
	fmt.Println(s1)

	//值为切片类型的map
	sm1 := make(map[string][]int, 2)
	sm1["tiezhu"] = []int{1, 2, 3}
	fmt.Println(sm1)

}
