package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
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
	urlData.Set("appKey", "123")
	queryStr := urlData.Encode() //编码后的url 地址
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr

	req, err := http.NewRequest("get", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req) //发送请求
	// req.Header.Add()
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

//客户端post请求实例
func postDemo() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=joel&age=18"
	// json
	contentType := "application/json"
	data := `{"name":"joel","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

func main1() {
	imgPath := "C:/Users/LM-LL/Desktop/"
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
