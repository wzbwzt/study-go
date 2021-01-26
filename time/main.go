package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main1() {
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
	//timer := time.Tick(time.Second)
	//for i := range timer {
	//	fmt.Println(i)
	//}

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

	//判断两个时间是否相等；这种方法还会比较地点和时区信息。
	fmt.Println(now.Equal(timer))
	//判断时间（now）是否在（timer）之前
	fmt.Println(now.Before(timer))
	//判断时间（now）是否在（timer）之后
	fmt.Println(now.After(timer))

	//Sleep
	time.Sleep(100) //默认纳秒
	time.Sleep(3 * time.Second)
	fmt.Println("3 second pass")
	n := 100
	time.Sleep(time.Duration(n)) //强转类型才可以使用

	//同样起到sleep的作用；time.After本质是通道
	<-time.After(time.Second * 2)

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
	timerT := time.Tick(time.Second * 2)
	for t := range timerT {
		fmt.Println(t)
	}

}
func main2() {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	println(currentYear, currentMonth)
	currentLocation := now.Location()
	fmt.Println(currentLocation)
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, -12, 0)

	fmt.Println(firstOfMonth)
	fmt.Println(lastOfMonth)
	fmt.Println(lastOfMonth)

	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Chongqing")
	zeroday, _ := time.ParseInLocation(layout, "1949-10-01 00:00:00", loc)
	fmt.Println(zeroday)
	str := now.Format(time.RFC3339)
	fmt.Println(str)
	if str > "0" {
		println(123)
	}

	str_test := "12-24-20"
	fmt.Println(str_test)
	timer_test, err := time.Parse("01-02-06", str_test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timer_test)
}

func main() {

	year, month, _ := time.Now().AddDate(0, -1, 0).Date()
	fmt.Println(year)
	fmt.Println(int(month))
	fmt.Println(month.String())
	fmt.Println(int(month))

	a := float64(0.1)
	fmt.Println(a)
	fmt.Println(float32(a))

	month_timer := GetOneMonthDays("2")
	for _, v := range month_timer {
		// _, m, d := v.Date()
		// fmt.Println(int(m))
		// fmt.Println(d)
		fmt.Println(v)

	}
	start, end := GetTimeRange(time.Now(), time.Now())
	fmt.Println(start)
	fmt.Println(end)

}

//计算指定年月的天数
func Count1(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
			fmt.Fprintln(os.Stdout, "The month has 31 days")
		}
	} else {
		//判定公历闰年应遵循的一般规律为:四年一闰，百年不闰，四百年再闰.
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}

//计算指定年月的天数
func Count2(year int, month int) (days int) {
	YMonth := strconv.Itoa(year) + strconv.Itoa(month)
	time.Parse("200601", YMonth)
	return
}

//获取一个月中的每一天
func GetOneMonthDays(month string) (result []time.Time) {
	if len(month) == 1 {
		month = "0" + month
	}
	startDate, _ := time.Parse("01", month)
	endDate := startDate.AddDate(0, 1, 0)
	for {
		if startDate.After(endDate) || startDate.Equal(endDate) {
			break
		} else {
			result = append(result, startDate)
		}
		startDate = startDate.AddDate(0, 0, 1)
	}
	return result
}

// 归一化时间范围，从起始天的00:00:00到最后一天的23:59:59
func GetTimeRange(s, e time.Time) (start, end time.Time) {
	loc := s.Location()
	yy, mm, dd := s.Date()
	start = time.Date(yy, mm, dd, 0, 0, 0, 0, loc)
	yy, mm, dd = e.Date()
	end = time.Date(yy, mm, dd, 23, 59, 59, 0, loc)
	return
}
