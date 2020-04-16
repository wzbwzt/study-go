package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) //纳秒
	//time.Unix()把时间戳转换为标准时间
	res := time.Unix(1585671894, 0)
	fmt.Println(res.Year())
	fmt.Println(res.Day())
	fmt.Println(res.Date())

	//时间间隔
	duration := time.Second
	fmt.Println(duration)

	//now +24hour
	fmt.Println(now.Add(time.Hour * 24))
	//定时器
	// timer := time.Tick(time.Second)
	// for i := range timer {
	// 	fmt.Println(i)
	// }

	//时间格式话  以2006 12345表示年月日时分秒
	fmt.Println(now.Format("2006-01-02 03:04:05"))
	//2020/09/09 09:09:09 AM
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))
	//毫秒
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))

	//按照对应的格式解析字符串类型的时间
	timer, err := time.Parse("2006-01-02", "2020-04-01")
	if err != nil {
		fmt.Printf("time parse failed err is %v\n", err)
		return
	}
	fmt.Println(timer.Unix())
	//Sub 两个时间相减
	d := now.Sub(timer)
	fmt.Println(d)

	//Sleep
	time.Sleep(100) //默认纳秒
	time.Sleep(3 * time.Second)
	fmt.Println("3 second pass")
	n := 100
	time.Sleep(time.Duration(n)) //强转类型才可以使用

	//指定时区
	//默认按照当前所在时区来解析时间
	fmt.Println(now)
	//默认按照当前所在时区来解析时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", "2020/04/03 14:00:20")
	if err != nil {
		fmt.Printf("time Parse failed err:%v", err)
		return
	}
	fmt.Println(timeObj)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("time location failed err :%v", err)
		return
	}
	fmt.Printf("%T\n", loc)
	timeObj2, err := time.ParseInLocation("2006/01/02 15:04:05", "2020/04/03 14:00:20", loc)
	if err != nil {
		fmt.Printf("err is %v", err)
		return
	}
	fmt.Println(timeObj2)
	td := timeObj2.Sub(now)
	fmt.Println(td)

	//tick 定时器
	timerT := time.Tick(time.Second * 5)
	for t := range timerT {
		fmt.Println(t)
	}
}
