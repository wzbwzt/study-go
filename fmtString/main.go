package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

	f := 12.34
	fmt.Printf("浮点数:%T\n", f)
	fmt.Printf("%b\n", f)               //无小数部分、二进制指数的科学计数法，如-123456p-78
	fmt.Printf("%e\n", f)               //科学计数法，如-1234.456e+78
	fmt.Printf("%E\n", f)               //科学计数法，如-1234.456E+78
	fmt.Printf("%f\n", f)               //有小数部分但无指数部分，如123.456
	fmt.Printf("%F\n", f)               //同%f
	fmt.Printf("%g\n", f)               //根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	fmt.Printf("%G\n", f)               //根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	fmt.Println(fmt.Sprintf("%.1f", f)) //保留指定小数位

	fmt.Printf("%08d\n", 12345)    //8位，不足前面凑0补齐 输出 00012345
	fmt.Printf("%0*d\n", 8, 12345) //同上  输出 00012345

	s := "hello/wzb"
	//打印字符串
	fmt.Printf("string:%s\n", s)
	fmt.Printf("value:%v\n", s)
	o := struct{ name string }{"teizhu"}
	fmt.Printf("%+v\n", o) //输出结构体时会添加字段名
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o) //值的Go语法表示
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

	//Fprint
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./test2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "阿无的吴"
	//向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

	ss1 := fmt.Sprint("沙河小王子")
	name1 := "沙河小王子"
	age := 18
	ss2 := fmt.Sprintf("name:%s,age:%d", name1, age)
	ss3 := fmt.Sprintln("沙河小王子")
	fmt.Println(ss1, ss2, ss3)

	//Errorf
	err3 := fmt.Errorf("这是一个错误") //定义一个error
	fmt.Println(err3)
	e := errors.New("原始错误e")
	fmt.Println(e)
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(w)
	var (
		name3   string
		age3    int
		married bool
	)
	fmt.Scan(&name3, &age3, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name3, age3, married)

}
func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

//前/后置补零
func buZero() {
	a := 12345
	fmt.Println(a) // 输出 12345
	//前置补0
	fmt.Printf("%08d\n", a)    //9位，不足前面凑0补齐 输出 00012345
	fmt.Printf("%0*d\n", 8, a) //同上  输出 00012345

	in := 12345
	fmt.Println(in) // 输出 12345
	// 需要输出 12300 后面两位置0

	// 小于100则不处理
	if in > 100 {
		in = in / 100 * 100
	}
	fmt.Println(in) // 输出 12300
}
