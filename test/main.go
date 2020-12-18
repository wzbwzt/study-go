package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"
)

/*
继承的优缺点
优点
简单，直观，关系在编译时静态定义；
被复用的实现易于修改，派生类可以覆盖基类的实现。
缺点
无法在运行时变更从基类继承的实现；
派生类的部分实现通常定义在基类中（派生类可以拓展基类的属性和行为）；
基类的实现细节直接暴露给派生类，破坏了封装；
基类的任何变更都强制子类进行变更（除非子类重写了该方法）。
组合的优缺点
优点
可以在运行时动态对象的行为；
保持类的封装以专注于单一业务；
只通过接口来访问对象，不会破坏封装；
减少对象的依赖关系，更加灵活。
缺点
系统的行为将依赖于不同对象，而不是定义在单个类中，不便于管理；
当需要新的行为时，需要不断定义新的对象来满足需求。
*/
/*
任务类型包括按照文件过期时间清理，按照文件夹容量进行清理及将文件上传至服务器，对于这些任务而言，
具有某些共同点：如都是定时执行的，并且支持启动、停止等操作。

通过Executable接口来将任务的共同点抽象为方法，并且由结构体Task（类）实现该接口定义任务的通用行为。
每类任务也定义为一个具体的结构体（类），并且通过组合Task类来复用Task的代码，使其具有通用行为，
对于各类任务的特有行为而言，如按过期时间清理任务需要遍历文件夹筛选出满足过期条件的文件，按照文件
夹容量清理任务需要先统计文件夹的总容量，当总容量大于警戒容量时再按照修改时间对文件列表进行排序，
从过期时间最久的文件开始删除，直至文件夹容量小于安全容量，则通过实现Executable接口定义的方法Execute来定义各自的行为。

*/

type Executable interface {
	Start()
	Execute()
}

type Task struct {
	executor Executable // 实现hook函数的效果：由子类负责编写业务代码
}

func (t *Task) Start() {
	println("Task.Start()")
	// 复用父类代码
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
		//t.Execute()         //Task.Execute()
		t.executor.Execute() // 实现hook函数的效果：由子类负责编写业务代码   //CleanTask.Execute()
	}
}

func (t *Task) Execute() {
	println("Task.Execute()")
}

type CleanTask struct {
	Task
}

func (ct *CleanTask) Execute() {
	println("CleanTask.Execute()")
}

func main1() {
	cleanTask := &CleanTask{
		Task{},
	}
	cleanTask.executor = cleanTask // 实现hook函数的效果：由子类负责编写业务代码
	cleanTask.Start()
}

func main2() {
	a := int32(1)
	b := int32(3)
	fmt.Println(a / b)
	passrate, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", float64(a)/float64(b)), 64)
	fmt.Println(passrate)

}

//在使用func (t Time) AddDate(years int, months int, days int) Time{}需要注意如果是8.31往后推一个月会直接跳过9月
func main3() {
	t, _ := time.Parse(time.RFC3339, "2020-08-31T08:10:45.814Z")
	year := t.Year()
	month := t.Month()
	local := t.Location()
	date := time.Date(year, month, 1, 0, 0, 0, 0, local)
	fmt.Println(date)
	format := date.AddDate(0, 0, 0).Format("200601")
	fmt.Println(format)
	format2 := date.AddDate(0, 1, 0).Format("200601")
	fmt.Println(format2)
}

func main4() {
	timer, _ := time.Parse("200601", "202009")
	fmt.Println(timer)
	addDate := timer.AddDate(0, 1, 0)
	fmt.Println(addDate)
	date := time.Date(timer.Year(), timer.Month(), 1, 0, 0, 0, 0, timer.Location())
	startime := date.Format(time.RFC3339)
	endtime := date.AddDate(0, 1, 0).Format(time.RFC3339)
	fmt.Println(startime)
	fmt.Println(endtime)

}
func main5() {
	//var m =map[int]int64{}
	m := make(map[int]int64)
	m[1] = 123
	fmt.Println(m)
}
func main16() {
	a := []int{1, 2, 3}
	sprintf := fmt.Sprintf(`{"deleted":%v}`, a)
	fmt.Println(sprintf)
	fmt.Printf("%T", sprintf)
	type Test struct {
		Deleted []int `json:"deleted"`
	}
	var t Test
	_ = json.Unmarshal([]byte(sprintf), &t)
	fmt.Println(t)
	var t2 Test
	t2.Deleted = a
	marshal, _ := json.Marshal(t2)
	fmt.Println(string(marshal))
	t3 := struct {
		Deleted []int
	}{}
	t3.Deleted = a
	marshal3, _ := json.Marshal(t3)
	fmt.Println(string(marshal3))
}

func main9() {
	a := "ABC123123"
	fmt.Println(a)
	marshal, _ := json.Marshal(a)
	fmt.Println(string(marshal))
	var c string
	_ = json.Unmarshal(marshal, &c)
	fmt.Println(c)
	var d []string
	d = append(d, c)
	bytes, _ := json.Marshal(d)

	var tmp1 []string
	_ = json.Unmarshal(bytes, &tmp1)
	for _, v := range tmp1 {
		fmt.Println(v)
	}
}
func main12() {
	a := "sh202010261002"
	fmt.Println(a[2:])
}

func main13() {
	a := []string{"aaa", "bbb", "ccc"}
	b := []string{"bbb", "ddd", "eee", "ccc", "aaa"}
	res := RemoveFromSlice(a, b)
	fmt.Println(res)
}
func RemoveFromSlice(target, source []string) []string {
	for _, v := range target {
		for k, vv := range source {
			if v == vv {
				source = append(source[:k], source[k+1:]...)
			}
		}
	}
	return source
}

//数组去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func GenNowNumber() string {
	return time.Now().Format("20060102150405")
}

func main29() {
	// m := make(map[string]int)
	// m["go"] = 123
	// m["ptyhon"] = 12
	// for k := range m {
	// 	fmt.Println(k)
	// }
	a := []string{"aaa", "abc", "acd", "aaa", "abc"}
	c := []string{}
	res := RemoveRepeatedElement(a)
	res1 := RemoveRepeatedElement(c)
	fmt.Println(res)
	fmt.Println(res1)
	test := float64(6)
	for i := float64(1); i < test; i++ {
		fmt.Println(i)
	}
	fmt.Println(GenNowNumber())
	timer, _ := time.Parse("2006", "2020")
	fmt.Println(timer.Format(time.RFC3339))

}

//判断浮点数是不是整数
func main26() {
	var a float64
	a = 1.23
	fmt.Println(int64(a))
	//1
	if a == float64(int64(a)) {
		fmt.Println("yay")
	} else {
		fmt.Println("you fail")
	}
	//2.
	fmt.Println(math.Trunc(a))
}

func main27() {
	var codes []string
	v := "[]"
	_ = json.Unmarshal([]byte(v), &codes)
	fmt.Println(codes)
	for _, v := range codes {
		fmt.Println(123)
		fmt.Println(v)
	}
	var transferCodes []string
	transferCodes = nil
	marshal, _ := json.Marshal(transferCodes)
	fmt.Println(string(marshal))
}

func main() {
	m := make(map[*bool]int)
	t := true
	f := false
	m[nil] = 0
	m[&t] = 1
	m[&f] = 2
	fmt.Println(m[nil])
	fmt.Println(m[&t])

	m1 := make(map[*bool]*int)
	one := 0
	two := 1
	three := 2

	m1[nil] = &one
	m1[&t] = &two
	m1[&f] = &three
	fmt.Println(*m1[nil])
	fmt.Println(*m1[&t])

	data := md5.Sum([]byte("wzbwzt"))
	fmt.Printf("%T\n", data)
	fmt.Printf("%x\n", data)

}
