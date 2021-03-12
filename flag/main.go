package main

import (
	"flag"
	"fmt"
	"time"
)

//flag  命令行标位符
func main1() {
	//定义命令行参数方式1// flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//定义命令行参数方式2 //flag.Type(flag名, 默认值, 帮助信息)
	//需要注意的是，此时name、age、married、delay均为对应类型的指针。
	//name := flag.String("name", "张三", "姓名")
	//age := flag.Int("age", 18, "年龄")
	//married := flag.Bool("married", false, "婚否")
	//delay := flag.Duration("d", 0, "时间间隔")

	//解析命令行参数（必须的步骤）
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	// fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}

func main() {
	s := flag.String("name", "wzb", "xingming")
	flag.Parse()
	fmt.Println(*s)
}
