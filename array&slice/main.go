package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
type slice struct {
  array unsafe.Pointer
  len   int
  cap   int
}
​
type Pointer *ArbitraryType

slice实际上是一个结构体类型，包含3个字段，分别是

array: 是指针，指向一个数组，切片的数据实际都存储在这个数组里。
len: 切片的长度。
cap: 切片的容量，表示切片当前最多可以存储多少个元素，如果超过了现有容量会自动扩容。
*/
func main() {
	//数组
	fmt.Println("数组")
	/*
		存放元素的容器
		必须指定存放的元素的类型和容器（长度）
		数组的长度是数组类型的一部分
	*/
	var a1 [3]bool //[true false true]
	var a2 [2]bool
	//a1 和a2是不可以比较的因为类型 不一样
	fmt.Println(a1, a2)
	//数组的初始化：
	//如果不初始化：默认元素都是零值（bool是false;整型和浮点型都是0;字符串是‘’）
	//初始化方式：
	//方式1
	a11 := [3]bool{true, true, true}
	fmt.Println(a11)
	//方式2:根据初始化值自动判断数组的长度
	a12 := [...]int{1, 234, 234, 123, 1, 45, 66, 88, 7896}
	fmt.Println(a12)
	//方式3：根据索引来初始化
	a13 := [6]int{0: 1, 5: 8}
	fmt.Println(a13)
	//数组的遍历
	name := [...]string{"Tom", "Dog", "Cat"}
	//方式1
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
	//方式2
	for i, v := range name {
		fmt.Println(i, v)
	}

	//多维数组
	//[[1 2] [3 4] [5 6]]
	var a21 [3][2]int
	a21 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a21)
	//多维数组的遍历
	for _, v1 := range a21 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	//数组是值类型（相当于复制黏贴）
	a221 := [3]int{1, 2, 3} //[1 2 3]
	a222 := a221            //[1 2 3] 相当于复制黏贴
	a222[0] = 100
	fmt.Println(a221, a222)

	b1 := 12
	b2 := b1
	b2 = 122
	fmt.Println(b1, b2)

	//切片  slice:
	//1.切片是指向了一个底层数组
	//2.切片的长度是他元素的个数
	//3.切片的容量是底层数组从切片第一个元素到最后一个元素的容量
	//4.切片的本质就是一个框，框住了一块连续的内存（数据类型必须一致）；属于引用类型。真正的
	//数据都是保存在底层数组里的
	fmt.Println("切片：slice")
	var sli1 []int    //定义一个存放int类型元素的切片
	var sli2 []string ////定义一个存放string类型元素的切片
	fmt.Println(sli1, sli2)
	fmt.Println(sli1 == nil) //是否为空

	//初始化
	sli1 = []int{1, 2, 3}
	sli2 = []string{
		"tiezhu",
		"gaiya",
		"tiedan",
	}
	fmt.Println(sli1, sli2)
	fmt.Println(sli1 == nil)
	//长度和容量
	fmt.Printf("len(sli1):%d cap(sli1):%d\n", len(sli1), cap(sli2))

	//由数组得到切片
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	arr2 := arr1[0:6] //左闭右开
	arr3 := arr1[3:]
	fmt.Println(arr2, arr3)
	//切片的容量指底层数组从切片第一个元素到最后一个元素的容量
	fmt.Printf("len(arr2):%d cap(arr2):%d\n", len(arr2), cap(arr2))
	fmt.Printf("len(arr3):%d cap(arr3):%d\n", len(arr3), cap(arr3))
	//切片再切割
	arrarr1 := arr2[2:]
	fmt.Println(arr2)
	fmt.Println(arrarr1)
	fmt.Printf("len(arrarr1):%d cap(arrarr1):%d\n", len(arrarr1), cap(arrarr1))
	//切片是引用类型  都是指向底层的一个数组
	fmt.Println(arr2)
	arr2[2] = 300000
	fmt.Println(arr2)
	fmt.Println(arrarr1)

	//make（）函数创建切片
	ms1 := make([]int, 5, 10) //cap不指定的化就会默认同长度
	fmt.Printf("ms1:%v  len:%d  cap:%d\n", ms1, len(ms1), cap(ms1))

	// 一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。
	//但是我们不能说一个长度和容量都是0的切片一定是nil
	//所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断

	//append()为切片追加元素   可以自动初始化切片
	app := []int{1, 2, 3}
	fmt.Printf("(追加前)app=%v;len=%d;cap=%d\n", app, len(app), cap(app))
	app = append(app, 6) //调用append函数必须用原来的切片变量接受返回值
	//append函数追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	fmt.Printf("(追加后1.0)app=%v;len=%d;cap=%d\n", app, len(app), cap(app))
	app = append(app, 6, 8)
	fmt.Printf("(追加后1.1)app=%v;len=%d;cap=%d\n", app, len(app), cap(app))
	app2 := []int{66, 77, 88}
	app = append(app, app2...) //...表示拆开切片
	fmt.Printf("(追加后1.2)app=%v;len=%d;cap=%d\n", app, len(app), cap(app))

	//copy()复制切片
	//切片是引用类型，但是copy()函数是将切片值赋值到另外一个内存中，所以不受底层数组的改变而改变
	cop := []int{1, 2, 3}
	cop2 := cop //赋值
	cop3 := make([]int, 3, 3)
	copy(cop3, cop)
	fmt.Println(cop, cop2, cop3)
	cop[0] = 100
	fmt.Println(cop, cop2, cop3)
	fmt.Println("--------------------")
	//删除切片元素
	//删除del切边中的3元素
	del := [...]int{1, 2, 3, 4, 5, 6}
	del1 := del[:]
	fmt.Println(del) //[1,2,3,4,5,6]
	fmt.Printf("%T\n", del1)
	fmt.Printf("%T\n", del)
	//1.切片不保存具体的值2.切片对应一个底层数组3.底层数组都是占用一块连续的内存
	del1 = append(del1[:2], del1[3:]...) //修改了底层数组
	fmt.Println(del)                     //[1,2,4,5,6,6]

	// []int->int
	sint := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T %v\n", sintToInt(sint), sintToInt(sint))

	splitSlice()

	appendSlice()

	audition()

	copyForSlice()
}

/**
类型的转换
*/

// []int->int
func sintToInt(sint []int) int {
	sstr := []string{}
	for _, v := range sint { //将[]int->[]string
		str := strconv.Itoa(v)
		sstr = append(sstr, str)
	}
	//将[]string->string->int
	n, err := strconv.Atoi(strings.Join(sstr, ""))
	if err != nil {
		fmt.Println("[]int->int failed err:", err)
	}
	return n

}

//字节
func f1() {
	n := "0123abc"
	fmt.Printf("%#v;%T\n", '0', '0')
	fmt.Printf("A:%#v;%T\n", 'A', 'A')
	for _, v := range n {
		fmt.Println(v - '0')
		fmt.Printf("%v-%T\n", v, v)
	}
}

//:分割操作符机制
/*
	1. :可以对数组或者slice做数据截取，:得到的结果是一个新slice。
	2. 新slice结构体里的array指针指向原数组或者原slice的底层数组，
	新切片的长度是：右边的数值减去左边的数值，
	新切片的容量是原切片的容量减去:左边的数值。
	3. :的左边如果没有写数字，默认是0，右边没有写数字，默认是被分割的数组或被分割的切片的长度。
	4. :分割操作符右边的数值有上限，上限有2种情况
	如果分割的是数组，那上限是是被分割的数组的长度。
	如果分割的是切片，那上限是被分割的切片的容量。
	注意，这个和下标操作不一样，如果使用下标索引访问切片，下标索引的最大值是(切片的长度-1)，而不是切片的容量。
*/
func splitSlice() {
	a := make([]int, 0, 4) // a的长度是0，容量是4
	b := a[:]              // 等价于 b := a[0:0], b的长度是0(0-0)，容量是4(4-0)
	c := a[:1]             // 等价于 c := a[0:1], b的长度是1(1-0)，容量是4(4-0)
	// d := a[1:]                     // 等价于 d:=a[1:0] 编译报错 panic: runtime error: slice bounds out of range
	e := a[1:4]                    // e的长度3(4-1)，容量3(4-1)
	fmt.Println(a, len(a), cap(a)) //[] 0 4
	fmt.Println(b, len(b), cap(b)) //[] 0 4
	fmt.Println(c, len(c), cap(c)) //[0] 1 4
	// fmt.Println(d, len(d), cap(d))
	fmt.Println(e, len(e), cap(e)) //[0 0 0] 3 4

	return
}

//append 机制
/*
- append函数返回的是一个切片，append在原切片的末尾添加新元素，这个末尾是切片长度的末尾，不是切片容量的末尾
- 如果原切片的容量足以包含新增加的元素，那append函数返回的切片结构里3个字段的值是：
array指针字段的值不变，和原切片的array指针的值相同，也就是append是在原切片的底层数组返回的切片还是指向原切片的底层数组
len长度字段的值做相应增加，增加了N个元素，长度就增加N
cap容量不变

- 如果原切片的容量不够存储append新增加的元素，Go会先分配一块容量更大的新内存，然后把原切片里的所有元素拷贝过来，最后在新的内存里添加新元素。append函数返回的切片结构里的3个字段的值是：
array指针字段的值变了，不再指向原切片的底层数组了，会指向一块新的内存空间
len长度字段的值做相应增加，增加了N个元素，长度就增加N
cap容量会增加
*/
func appendSlice() {
	fmt.Println("append 机制")
	a := make([]int, 0, 4)
	fmt.Printf("%#v\n", a)
	b := append(a, 1) // b=[1], a指向的底层数组的首元素为1，但是a的长度和容量不变

	fmt.Println(a, b)              //[] [1]
	fmt.Println(b, len(b), cap(b)) //[1] 1 4

	// c := append(a, 2)    // a的长度还是0，c=[2], a指向的底层数组的首元素变为2
	// fmt.Println(a, b, c) // [] [2] [2]

	arraylist := [6]int{0, 1, 2, 3, 4, 5}
	slice1 := arraylist[:2]
	fmt.Println(slice1, len(slice1), cap(slice1)) //[0 1] 2 6
	slice2 := append(slice1, 8)
	fmt.Println(slice2, len(slice2), cap(slice2)) //[0 1 8] 3 6

	fmt.Println(arraylist) //[0 1 8 3 4 5]

}

//slice 扩容机制
/*
slice的扩容机制随着Go的版本迭代，是有变化的。目前网上大部分的说法是下面这个：

当原 slice 容量小于 1024 的时候，新 slice 容量变成原来的 2 倍；原 slice 容量超过 1024，新 slice 容量变成原来的1.25倍。
这里明确告诉大家，这个结论是错误的。

slice扩容的源码实现在src/runtime/slice.go里的growslice函数，
源码地址：https://github.com/golang/go/blob/master/src/runtime/slice.go。
*/
func expansionSlice() {

}

//copy机制
/*
注意：原切片和目标切片的内存空间可能会有重合，copy后可能会改变原切片的值
*/
func copyForSlice() {
	a := []int{1, 2, 3}
	b := a[1:]        // [2 3]
	copy(a, b)        // a和b内存空间有重叠
	fmt.Println(a, b) // [2 3 3] [3 3]
}

//面试题目
func audition() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	x = append(x, y...)
	x = append(x, y...)
	fmt.Println(a, x) //[0 2 3 3] [0 2 3 3 3]
}
