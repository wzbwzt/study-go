package main

//导入语句
import (
	"fmt"
)

//在编程语言中标识符就是程序员定义的具有特殊意义的词，比如变量名、常量名、函数名等等。
//Go语言中标识符由字母数字和_(下划线）组成，并且只能以字母和_开头。 举几个例子：abc, _, _123, a123。

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
const (
	x = iota //0
	y        //1
	z = "zz" //zz
	k        //zz
	p = iota //4
)
const (
	a = "name"
	b = iota
	c = iota
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

//关于go的垃圾回收机制
//go垃圾回收其实就是自动回收堆中没有被执行的值时会自动回收该块内存；比如当申请了一个引用类型时，变量名会存放在
//栈中，具体值会存放在堆中，当函数执行完毕，变量名会被释放掉，这时对应的堆中的值不再有值来指向，go垃圾回收会自动将其回收;
//也可以通过指针来指向nil来释放掉;实际应用的实例：
//比如函数内部需要实例化一个结构体时，
//a:=new(mystruct):会将值存放在堆中，函数结束时会，选哟go的垃圾回收机制来自动回收；
//a:=mystruct():会直接将值存放在栈中，函数结束时，由操作系统来直接释放掉

//
//go函数外的语句比较是以关键字开头
/*
new:用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针,也可以直接分配slice,map等
make:	用来分配内存，主要用来分配引用类型，比如chan、map、slice
引用类型的值的内存分配都是在堆中，栈中使用地址来指向;当栈中的地址被释放，对应的堆中的值不再被栈中的
地址来指向，就会被go的垃圾回收机制来自动释放掉
当需要主动回收时，将栈中的指针指向nil既可



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

	// v := int8(1) //只能存8位
	//fmt.Println(v << 10) //错误的用法  int8只有8位  左移10位虽然不会报错但是只会展示8位即：00000000

	testf := float32(1.2)
	testI := int64(testf)
	fmt.Println(testI)

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

}
