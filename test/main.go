// +build ignore
package main

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/labstack/gommon/log"
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

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- elem
			fmt.Println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()

	}()

	return ch
}

func main188() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}
	v := <-th.Iter()
	fmt.Printf("%s%v", "ch", v)

}

func main199() {
	channel := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			channel <- i
		}()
	}

	for i := 0; i < 10; i++ {
		num := <-channel
		fmt.Println("num:", num)
	}

	WorkTimeStart := "22:00:01"
	WorkTimeEnd := "08:00:30"
	ZeroTime := "24:00:00"
	workStart, err := time.Parse("15:04:05", WorkTimeStart)
	workEnd, err := time.Parse("15:04:05", WorkTimeEnd)
	zeroTime, err := time.Parse("15:04:05", ZeroTime)
	if err != nil {
		fmt.Println(err)
	}
	if workStart.Before(workEnd) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if zeroTime.After(workStart) {
		fmt.Println("zer0")
	}

}

func main1() {
	cleanTask := &CleanTask{
		Task{},
	}
	cleanTask.executor = cleanTask // 实现hook函数的效果：由子类负责编写业务代码
	cleanTask.Start()
}

func main18() {
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

func main28() {
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

func main88() {
	a := []int{}
	if a == nil {
		fmt.Println("is nil")
	}
	if len(a) == 0 {
		fmt.Println("is 0")
	}
	fmt.Println(a)

	str := "03排B01"
	pos := strings.Index(str, "排")
	fmt.Println(pos)
	strs := []rune(str)
	regionName := string(strs[:pos+1])
	shelveName := string(strs[pos+1:])
	fmt.Println(regionName)
	fmt.Println(shelveName)

	strr := []rune(str)
	fmt.Printf("str len is:%v", len(strr))
	fmt.Printf("strs  is:%v", strr)
	for _, v := range strr {
		fmt.Println(string(v))
	}

	fmt.Println("==========================")
	for _, v := range str {
		fmt.Println(string(v))
	}

	s := "device/sdk/CMakeLists.txt"
	ss := "CMakeListstxt"
	fmt.Println(path.Ext(s))
	fmt.Println(path.Ext(ss))

	var Map_img map[string]struct{} = map[string]struct{}{
		".jpg": struct{}{},
		".png": struct{}{},
		".gif": struct{}{},
		".svg": struct{}{},
		".pcx": struct{}{},
	}
	fmt.Println(Map_img)

	fmt.Println("==========================")
	str_time := "202012"
	timer, _ := time.Parse("200601", str_time)
	fmt.Println(timer)
	start_time := timer.Format("2006-01-02 15:04:05")
	fmt.Println(start_time)
	end_time := timer.AddDate(0, 1, 0).Format("2006-01-02 15:04:05")
	fmt.Println(end_time)

	var xais []string
	for i := 0; i < 6; i++ {
		month := time.Now().AddDate(0, -i, 0).Format("200601")
		xais = append(xais, month)
	}
	fmt.Println(xais)
	map_test := make(map[int]int)
	map_test[1] = 1
	fmt.Println(map_test[1])
	fmt.Println(map_test[2])

	slice_test := []string{}
	fmt.Println(slice_test)
	for _, v := range slice_test {
		fmt.Println(v)
	}

	fmt.Println("+++++++++++++++++++++++++++++++++")
	type Peo struct {
		Name string
		Age  int
	}
	map_demo1 := make(map[int]*Peo)
	map_demo1[1] = &Peo{
		"Joel",
		12,
	}
	map_demo1[2] = &Peo{
		"Alice",
		16,
	}
	fmt.Println(map_demo1)
	fmt.Println(map_demo1[1])
	fmt.Println(map_demo1[1].Name)

	value, ok := map_demo1[1]
	if ok {
		value.Name = "Joel2.0"
		value.Age = 16
		// map_demo1[1] = value
	}
	fmt.Println(map_demo1[1])

}

func main2() {
	map_struct := make(map[int]*struct {
		Num  int
		Name string
	})
	map_struct[1] = &struct {
		Num  int
		Name string
	}{
		100, "Jole",
	}
	fmt.Println(map_struct)
	fmt.Println(map_struct[1])

	value, ok := map_struct[2]
	if !ok {
		map_struct[2] = &struct {
			Num  int
			Name string
		}{}
	}
	fmt.Println(map_struct[2].Num)
	fmt.Println(value)
	// fmt.Println(value.Num)

	// map_struct[2] = &struct {
	// 	Num  int
	// 	Name string
	// }{
	// 	Num:  value1.Num + 1,
	// 	Name: "Alice",
	// }

	fmt.Println(map_struct)
	fmt.Println(map_struct[1])
	fmt.Println(map_struct[2])

	sec_struct := make(map[int]*SecLimit)
	sec_struct[1] = &SecLimit{
		100, 88888888,
	}
	value_2, ok := sec_struct[2]
	if !ok {
		value_2 = &SecLimit{}
		sec_struct[2] = value_2
	}

	// fmt.Println(sec_struct[2])
	// value_2.Count(6666666)
	// fmt.Println(sec_struct)
	// fmt.Println(sec_struct[2])

	// value_2.count = 200
	// value_2.curTime = 666666
	sec_struct[2].count = 200
	sec_struct[2].curTime = 666666666

	fmt.Println(sec_struct)
	fmt.Println(sec_struct[1])
	fmt.Println(sec_struct[2])
	for v, vv := range sec_struct {
		fmt.Println(v)
		fmt.Println(vv)
	}

	fmt.Println(sec_struct)
	for v := range sec_struct {
		fmt.Println(v)
	}
}

type SecLimit struct {
	count   int   //每秒访问数量
	curTime int64 //访问的时间(精确到秒)
}

//Count 更新每秒访问数量
func (s *SecLimit) Count(nowTime int64) (newCount int) {
	if s.curTime == nowTime {
		return s.count + 1
	}
	s.count = 1
	s.curTime = nowTime
	return 1
}

func main888() {
	all_price := fmt.Sprintf("%.2f", 1.234)
	fmt.Println(all_price)

	timer, _ := time.Parse("20060102", "20201231")
	y, m, _ := timer.Date()
	timer_fix := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	fmt.Println(timer_fix)

	var xais []string
	for i := 5; i >= 0; i-- {
		y, m, _ := time.Now().Date()
		timer := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
		month := timer.AddDate(0, -i, 0).Format("200601")
		xais = append(xais, month)
	}
	fmt.Println(xais)

	imgPath := "C:\\Users\\Asche\\go\\src\\GoSpiderTest\\"
	imgUrl := "http://hbimg.b0.upaiyun.com/32f065b3afb3fb36b75a5cbc90051b1050e1e6b6e199-Ml6q9F_fw320"

	fileName := path.Base(imgUrl)

	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(imgPath + fileName)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)
}

/*-------------------------------分割线----------------------------*/
//KMP字符串匹配
func SindexKMP(S, T string) int {
	//next := get_next(T)
	next := NextArray(T)
	i := 0
	j := 0
	//同时满足才可以  找除字符串出现的第一个位置
	for i <= len(S)-1 && j <= len(T)-1 {

		if j == -1 || S[i] == T[j] {
			//当字符匹配时 i j 都加1
			i++
			j++
		} else {
			//子串的 偏移量 从next数组中取  i 不变
			j = next[j]
		}
	}
	//如果 j 大于 或者 等于 T串的长度 说明匹配成功
	if j >= len(T)-1 {
		return i - len(T) + 1
	}

	return 0
}
func NextArray(needle string) []int {
	l := len(needle)
	next := make([]int, l)
	next[0] = -1
	k := -1
	i := 0
	for i < l-1 {
		if k == -1 || needle[k] == needle[i] {
			i++
			k++
			next[i] = k
		} else {
			k = next[k]
		}
	}
	return next
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", *v)
	}
}

func main222() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("++++")
			f := err.(func() string)
			fmt.Println(err, f(), reflect.TypeOf(err).Kind().String())
		} else {
			fmt.Println("fatal")
		}
	}()
	defer func() {
		panic(func() string {
			return "defer panic"
		})
	}()
	panic("panic")

}

func testF(s []*int) {
	tmp := []int{1, 2, 3}
	t := make([]int, len(tmp), len(tmp))
	copy(t, tmp)
	a := 123
	s = append(s, &a)
}

//struct 转 map
func StructToMap(source interface{}) map[string]interface{} {

	valueof := reflect.ValueOf(source)
	if valueof.Kind() == reflect.Ptr {
		valueof = valueof.Elem()
	}
	if valueof.Kind() != reflect.Struct {
		panic("not struct")
	}

	typeof := reflect.TypeOf(source)
	if typeof.Kind() == reflect.Ptr {
		typeof = typeof.Elem()
	}

	res := make(map[string]interface{}, valueof.NumField())

	for i := 0; i < valueof.NumField(); i++ {
		filedvalue := valueof.Field(i)
		filedtype := typeof.Field(i)

		if filedvalue.Kind() == reflect.Ptr {
			filedvalue = filedvalue.Elem()
		}

		if filedtype.Tag.Get("rft") == "omit" {
			continue
		}
		switch filedvalue.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
			res[filedtype.Name] = filedvalue.Int()
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
			res[filedtype.Name] = filedvalue.Uint()
		case reflect.Float32, reflect.Float64:
			res[filedtype.Name] = filedvalue.Float()
		case reflect.String:
			res[filedtype.Name] = filedvalue.String()
		case reflect.Bool:
			res[filedtype.Name] = filedvalue.Bool()
		case reflect.Struct:
			if filedvalue.Type().Name() == "Time" {
				res[filedtype.Name] = filedvalue.Interface().(time.Time).Format(time.RFC3339)
			}
		}
	}
	return res
}

type MyErr struct {
	Code int
	Msg  string
}

func (this *MyErr) Error() string {
	return this.Msg
}

func NewMyErr(code int, msg string) error {
	return &MyErr{Code: code, Msg: msg}
}
func (this *MyErr) Unwrap(err error) error {
	return &MyErr{Code: 2, Msg: "bottom err"}
}

type User struct {
	Name string
	Age  int
}

func (this *User) Test() {
	this.Age++
}

type item struct {
	Name string
}

func main21() {
	item := &item{}
	// defer func() {
	// 	fmt.Println(item.Name)
	// }()
	defer func() {
		log.Infof("Item:%s", item.Name)
	}()
	item.Name = "joel"
	fmt.Println("over")
	return
	t := User{Age: 2}
	t.Test()
	fmt.Println(t)
	return
	list := make([]int, 10)
	alist := list //引用类型
	for i := 0; i < 10; i++ {
		alist[i] = i
	}
	fmt.Println(list)
	fmt.Println(alist)
	return
	err := NewMyErr(1, "failed")
	fmt.Println(errors.Unwrap(err))
	fmt.Println(err)

	return
	type idType int
	type Hoby struct {
		Name string
	}
	type User struct {
		Hoby
		Id        idType
		Name      *string
		Man       bool `rft:"omit"`
		Childuser *User
		Current   *time.Time
	}
	name := "joel"
	id := 1
	now := time.Now()
	childuser := User{Current: &now}
	u := User{Name: &name, Id: idType(id), Man: true, Childuser: &childuser, Current: &now}

	m := StructToMap(&u)
	for k, v := range m {
		fmt.Printf("%s-%v\n", k, v)
	}

	// mm := StructToMap(user{})
	// fmt.Println(mm)

	return
	// s1 := []*int{}
	s1 := make([]*int, 20)
	testF(s1)
	fmt.Println(s1)

	maptest := make(map[int]int, 2)
	fmt.Println(maptest)
	fmt.Println(len(maptest))
	maptest[1] = 1
	fmt.Println(maptest)
	fmt.Println(maptest[2])
	return
	atest := []int{1, 2, 3}
	for k, v := range atest {
		println(k, v)
	}
	return
	var ss []*int
	ss = append(ss, nil)
	fmt.Println(ss)
	fmt.Println(len(ss))
	return
	jointime, err := time.Parse(time.RFC3339, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jointime)
	return
	mapT := make(map[int]string)
	mapT[1] = "A"
	fmt.Println(mapT[1])
	fmt.Println(mapT[2])

	return
	aa := "ajdsfajf"
	saa := &aa
	fmt.Printf("%v", saa)
	return
	ss1 := "asfkashdfk"
	for _, v := range ss1 {
		fmt.Println(v)
	}
	return
	s := []int64{1, 2}
	func(req []int64) {
		// req = append(req, 3) //发生了扩容，重新分配内存
		fmt.Printf("%#v-%p\n", req, req)
		req[1] += 9
	}(s)
	fmt.Printf("%#v-%p\n", s, s)

	return
	var tt []int64
	fmt.Printf("%#v-%p", tt, tt)
	for _, v := range tt {
		fmt.Println("hello")
		fmt.Println(v)
	}

	return

	address := "联排21-120"
	valid := regexp.MustCompile("[0-9]")
	sli := valid.FindAllStringSubmatch(address, -1)
	sli2 := valid.FindAllString(address, 1)
	sli3 := valid.FindString(address)
	fmt.Printf("%#v\n", sli)
	fmt.Printf("%#v\n", sli2)
	fmt.Printf("%#v\n", sli3)

	return
	HouseID := fmt.Sprintf("%s%s%07d", "3302990200", "121", 1)
	num := HouseID[14:]
	new_num, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return
	}
	houseid := HouseID[:14] + fmt.Sprintf("%07d", new_num+1)
	fmt.Println(houseid)

	return

	var test_s []int
	fmt.Printf("%#v", test_s)
	if test_s == nil {
		fmt.Println("is null")
	}
	fmt.Println(len(test_s))
	a := make([]int, 0)
	fmt.Printf("%#v", a)
	fmt.Println(len(a))
	if a == nil {
		fmt.Println("is null")
	}
	return

	averagePrice, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", -1980/float64(-2)), 64)
	fmt.Println(averagePrice)
	return

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

	action_card := []byte{}
	str_action_card := string(action_card)
	fmt.Printf("%v", str_action_card)
	fmt.Printf("%v", action_card)
	// id := ""
	// fmt.Println(&id)

	// json_str := `{"小区id": null, "物业名称": "物业名称", "竣工时间": "竣工时间"}`
	type residential struct {
		Name    *int64  `json:"小区id"`
		PMCName *string `json:"物业名称"`
		EndTime *string `json:"竣工时间"`
	}
	out := new(residential)
	// json.Unmarshal([]byte(json_str), out)
	// fmt.Println(out)
	json_str2 := "[]"
	json.Unmarshal([]byte(json_str2), out)
	fmt.Printf("%#v\n", out)
	fmt.Printf("%v", out)
	marsh_str := residential{
		Name:    nil,
		PMCName: nil,
		EndTime: nil,
	}
	byte_marsh_str, _ := json.Marshal(&marsh_str)
	fmt.Println(string(byte_marsh_str))

	i, err := strconv.ParseInt("1611313585853", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i/1000, 0).Format("2006-01-02 15:04:05")
	fmt.Println(tm)

	fmt.Println("====================================================")
	path := "https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3887379252,3315859978&fm=26&gp=0.jpg"
	img_res, err := http.Get(path)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer img_res.Body.Close()

	// byte_img, err := ioutil.ReadAll(img_res.Body)
	// s := base64.StdEncoding.EncodeToString(byte_img)
	// s = "data:image/jpeg;base64," + s

	// img, _, err := imageorient.Decode(img_res.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	fmt.Printf("%#v\n", s)
	// fmt.Printf("%#v\n", img)
	var a2 []*string
	if a2 == nil {
		fmt.Println(123)
	}
	fmt.Printf("%#v", a2)
	fmt.Printf("%p\n", a2)
	str_accessory, _ := json.Marshal(a2)
	fmt.Println(len(str_accessory))
	fmt.Println(string(str_accessory))

	var a3 []string
	if a3 == nil {
		fmt.Println(456)
	}
	a1 := []*string{}
	if a1 == nil {
		fmt.Println(234)
	}
	fmt.Printf("%#v", a1)
	fmt.Printf("%p", a1)
	str_accessory, _ = json.Marshal(a1)
	fmt.Println(string(str_accessory))

	str_accessory, _ = json.Marshal(nil)
	str_res := string(str_accessory)
	fmt.Println(&str_res)
	if &str_res == nil {
		fmt.Println("is nil")
	}
	fmt.Printf("%#v\n", string(str_accessory))
	fmt.Println(len(str_accessory))

	var test1 *string
	fmt.Println(&test1)
	if &test1 == nil {
		fmt.Println("is nill")
	}

	//
	slice_test := [][]float32{
		{123, 234.9},
		{123, 234.9},
		{123, 234.9},
	}
	slice_b, _ := json.Marshal(slice_test)
	fmt.Println(string(slice_b))
	var Unmarshal_test [][]float32
	json.Unmarshal(slice_b, &Unmarshal_test)
	fmt.Printf("%#v\n", Unmarshal_test)

	// t := time.Now()
	// fmt.Println(int(t.Weekday()))

	s_test := "wzbwzt"
	fmt.Println(s_test)
	fmt.Println(string(s_test[1]))
	for _, v := range s_test {
		fmt.Println(string(v))
	}
}

type User1 struct {
	Name string
}

//go:generate go version
func main33() {
	a := make([]int, 10)
	fmt.Println(a)
	a = append(a, 1)
	fmt.Println(a)
}

func main() {
	data := `[
        {
            "date": "2022-2-23",
            "count": 2
        },
        {
            "date": "2022-2-22",
            "count": 1
        },
        {
            "date": "2022-2-28",
            "count": 10
        },
        {
            "date": "2022-2-27",
            "count": 0
        },
        {
            "date": "2022-2-26",
            "count": 12
        },
        {
            "date": "2022-2-25",
            "count": 0
        },
        {
            "date": "2022-2-24",
            "count": 0
        }
    ]`

	var res []*staticGotoCheckRspModel
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		log.Error(err)
	}
	for _, v := range res {
		fmt.Printf("%#v\n", v)
	}
	fmt.Println("sort res")
	sort.Sort(staticGotoCheckRspListModel(res))
	for _, v := range res {
		fmt.Printf("%#v\n", v)
	}

}

type staticGotoCheckRspModel struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
type staticGotoCheckRspListModel []*staticGotoCheckRspModel

func (s staticGotoCheckRspListModel) Len() int {
	return len(s)
}
func (s staticGotoCheckRspListModel) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s staticGotoCheckRspListModel) Less(i, j int) bool {
	return s[i].Date < s[j].Date
}
