package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	n, err := ioutil.ReadFile("./home.html")
	if err != nil {
		w.Write([]byte("页面丢失了！"))
	}
	w.Write(n)
}

//模拟客户端请求；同浏览器
func f2(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()
	fmt.Println(queryMap) //自动识别请求URl中的参数 并以map形式展示
	fmt.Println(queryMap["name"])
	fmt.Println(queryMap["age"])

	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body)) //在服务端打印客户端发送的请求的body
	w.Write([]byte("hello"))
}
func main() {
	http.HandleFunc("/home", f1)
	http.HandleFunc("/simulation/", f2)
	// http.ListenAndServe("127.0.0.1:9090", nil) //监听并服务于指定的IP和端口；
	http.ListenAndServe("0.0.0.0:9090", nil) //0.0.0.0 表示监听任何网段的信息

}