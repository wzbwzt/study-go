package main

import "fmt"

/*
map 的 key 为什么是无序的？

map 在扩容后，会发生 key 的搬迁，原来落在同一个 bucket 中的 key，搬迁后，
有些 key 就要远走高飞了（bucket 序号加上了 2^B）。而遍历的过程，就是按顺序遍历 bucket，
同时按顺序遍历 bucket 中的 key。搬迁后，key 的位置发生了重大的变化，有些 key 飞上高枝，
有些 key 则原地不动。这样，遍历 map 的结果就不可能按原来的顺序了。
当然，如果我就一个 hard code 的 map，我也不会向 map 进行插入删除的操作，
按理说每次遍历这样的 map 都会返回一个固定顺序的 key/value 序列吧。的确是这样，
但是 Go 杜绝了这种做法，因为这样会给新手程序员带来误解，以为这是一定会发生的事情，
在某些情况下，可能会酿成大错。
当然，Go 做得更绝，当我们在遍历 map 时，并不是固定地从 0 号 bucket 开始遍历，
每次都是从一个随机值序号的 bucket 开始遍历，并且是从这个 bucket 的一个随机序号的 cell 开始遍历。
这样，即使你是一个写死的 map，仅仅只是遍历它，也不太可能会返回一个固定序列的 key/value 对了。

`多说一句，“迭代 map 的结果是无序的”这个特性是从 go 1.0 开始加入的。`
*/

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
