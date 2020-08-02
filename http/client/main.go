package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
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
