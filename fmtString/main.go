package main

import (
	"fmt"
	"strings"
)

//fmt占位符
//fmt.Printf()格式化输出
//fmt.Print()连续不换行输出
//fmt.Println()换行输出
func main() {
	//打印内存地址
	ss := []int{1, 2, 3}
	fmt.Printf("%p\n", &ss) //使用%p  传的值得是内存地址（指针）
	fmt.Printf("%p\n", &ss[0])
	fmt.Printf("%p\n", &ss[1])
	fmt.Printf("%p\n", &ss[2])

	n := 100
	//打印类型
	fmt.Printf("%T\n", n)
	//打印值
	fmt.Printf("%v\n", n)
	//打印二进制
	fmt.Printf("%b\n", n)
	//打印十进制
	fmt.Printf("%d\n", n)
	//打印八进制
	fmt.Printf("%o\n", n)
	//打印十六进制
	fmt.Printf("%x\n", n)
	s := "hello/wzb"

	//打印字符串
	fmt.Printf("string:%s\n", s)
	fmt.Printf("value:%v\n", s)
	fmt.Printf("value:%#v\n", s)
	sc := 'c'
	fmt.Printf("字符:%c\n", sc)    //打印字符
	fmt.Println(len(s))          //字符串长度
	ret := strings.Split(s, "/") //字符串分割
	fmt.Println(ret)
	fmt.Printf("%T\n", ret)
	fmt.Println(strings.Contains(s, "wzb")) //包含
	fmt.Println(strings.HasPrefix(s, "h"))  //前缀
	fmt.Println(strings.HasSuffix(s, "h"))  //后缀

	sIndex := "adfadfwett"
	fmt.Println(strings.Index(sIndex, "a")) //查找index
	fmt.Println(strings.Join(ret, "+"))     //拼接

	//字符串中拿出具体的字符
	strRune := "hello铁柱你好안녕하세요."
	for _, c := range strRune {
		fmt.Printf("%c\n", c)
	}
	//字符串修改
	s1 := "白萝卜"
	s2 := []rune(s1) //1.将字符串转换为一个rune切片  rune的类型即为int32
	s2[0] = '红'      //2。修改的话必须用单引号 因为装换位rune后所有的选择都是字符
	fmt.Println(s2)
	fmt.Println(string(s2)) //把切片强制转换为字符串
}
