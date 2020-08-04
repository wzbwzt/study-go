package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main1() {
	//get请求示例1；对于get的提交的query值无法编译
	// resp, err := http.Get("http://127.0.0.1:9090/simulation?name=铁柱&id=123")
	// if err != nil {
	// 	fmt.Println("get request failed err:", err)
	// 	return
	// }

	//get请求示例2；对get的提交的query值编译  和自定义
	var urlObj, _ = url.Parse("http://127.0.0.1:9090/simulation/")
	var urlData = url.Values{} //url  map  可以对其编码
	urlData.Set("name", "铁柱")
	urlData.Set("age", "123")
	queryStr := urlData.Encode() //编码后的url 地址
	 fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("get", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req) //发送请求
	if err != nil {
		fmt.Println("send request failed;err:", err)

	}
	defer resp.Body.Close() //resp.Body 一定要关闭
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response failed err:", err)
		return
	}
	fmt.Println(string(b))
}

//KMP字符串匹配
func SindexKMP(S,T string) int {
	//next := get_next(T)
	next:=NextArray(T)
	i := 0
	j := 0
	//同时满足才可以  找除字符串出现的第一个位置
	for ;i <= len(S) -1  && j <= len(T) -1;{

		if j == -1|| S[i] == T[j]{
			//当字符匹配时 i j 都加1
			i++
			j++
		}else{
			//子串的 偏移量 从next数组中取  i 不变
			j = next[j]
		}
	}
	//如果 j 大于 或者 等于 T串的长度 说明匹配成功
	if j >= len(T) -1 {
		return i - len(T) +1
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
		{Name: "zhou",Age: 24},
		{Name: "li",Age: 23},
		{Name: "wang",Age: 22},
	}
	for i:=0;i<len(stus);i++{
		m[stus[i].Name]=&stus[i]
	}
	for _,stu:=range stus{
		m[stu.Name]=&stu
	}


	for k,v:=range m{
		fmt.Println(k,"=>",*v)
	}
}





func calc(index string, a, b int) int {
	ret := a+ b
	fmt.Println(index,a, b, ret)
	return ret
}
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}



