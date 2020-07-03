package main

//导入语句
import (
	"fmt"
	"strconv"
	"strings"
)

//函数外只能放置标识符的声明（变量/常量/函数/类型）
// var name string
// var age int
// var isOk bool

//Go语言中如果标识符的首字母是大写的话 就表示是对外部包是可见的
//此外大写的标识符的上面必须要些注释，变量名+注释（空格隔开）

//批量声明变量        //go语言中非全局变量声明必须使用   不使用就无法编译过去
var (
	name string
	age  int
	isOk bool
)

//批量申明常量
const (
	statusOK = 200
	notFound = 404
)

//n2，n3...延续第一个的值
const (
	n1 = 666
	n2
	n3
)

//iota ：当const出现就会被置为0  const中每新增一行常量声明iota就会计数一次
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	_         //1
	b2        //2
)

//插队
const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
	c5        //4
)

//多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1=1;d2=2
	d3, d4 = iota + 1, iota + 2 //d3=2;d4=3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //‘<<’左移  1左移10位就是2的10次方=1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

//Go语言中字符串必须是用“”双引号
//单引号包裹的是字符  例如'h'/'汉'这些单独的字母/汉字/符号
//多行字符串打印
var ss string = `
	世情薄
	人情恶
	雨送黄昏花易落
`

//go函数外的语句比较是以关键字开头
/*
new:用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
make:	用来分配内存，主要用来分配引用类型，比如chan、map、slice

数据类型
整型
无符号整型：uint8、uint16、uint32、uint64
带符号整型：int8、int16、int32、int64
uint和int:具体是32位还是64位看操作系统
uintptr:表示指针

浮点型
float64 和float32  默认是float64

复数
complex128 和 complex64   科学运算用到

布尔型
不能和其他的类型转换

byte 和 rune 类型
都属于类型别名
byte是uint8
rune是int32

go语言中字符串都是UTF8编码，UTF8编码中一个常用汉字一般占用3个

go语言中无法直接定义二进制数；
定义八进制数前面加0
定义十六进制数前面加0x

go语言中函数传的参数永远是复制黏贴的值

*/
//程序的入口函数
func main() {
	fmt.Println("hello wzb")
	//声明的同时赋值
	var s1 string = "tiezhu"
	//类型推导（根据值判断该变量的类型）
	var s2 = "gaiya"
	//简短变量声明  只能在函数内部使用
	s3 := "tiedan" //同一作用域中不能重复申明同名的变量
	//匿名变量
	_ = "wuming"
	fmt.Println(s1, s2, s3)

	//运算符  -位运算（针对的是二进制）
	//例如：5的二进制是101；2的二进制是10
	//& 按位与（两位均为1才为1）
	fmt.Println(5 & 2) //000
	//| 按位或（两位有个为1就为1）
	fmt.Println(5 | 2) //111
	//^ 按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) //111
	//<< 将二进制位左移指定位数
	fmt.Println(5 << 1)  //1010
	fmt.Println(1 << 10) //10000000000=1024
	//>> 将二进制位右移指定位数
	fmt.Println(5 >> 2) //1

	//v := int8(1) //只能存8位
	//fmt.Println(v << 10) //错误的用法  int8只有8位  左移10位虽然不会报错但是只会展示8位即：00000000

	//赋值运算符  用来给变量赋值
	val := 10
	val += 2 //val=val+2
	val -= 2
	val *= 2
	val /= 2
	val %= 3
	val <<= 2 //val=val<<2
	val &= 2  //val=val&2
	val |= 2
	val ^= 2
	val >>= 2
	fmt.Println(val)
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

	//切片  slice:1.切片是指向了一个底层数组2.切片的长度是他元素的个数
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
