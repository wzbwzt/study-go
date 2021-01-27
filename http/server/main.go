package main

import (
	"fmt"
	"io/ioutil"
	"net"
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

//启动http服务的方式1
func main1() {
	http.HandleFunc("/home", f1)
	http.HandleFunc("/simulation/", f2)
	// http.ListenAndServe("127.0.0.1:9090", nil) //监听并服务于指定的IP和端口；
	http.ListenAndServe("0.0.0.0:9090", nil) //0.0.0.0 表示监听任何网段的信息
}

//启动http服务的方式2
type myHandleType struct {
	addr string
}

func new() *myHandleType {
	return &myHandleType{
		addr: "0.0.0.0:9090",
	}
}

//实现这个ServerHTTP这个方法从而实现http.Handler这个接口类
func (m *myHandleType) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/v1" {
		httpCode := http.StatusOK
		text := http.StatusText(httpCode)
		fmt.Println("200", text)
		rsp.Write([]byte("hello!"))
		return
	}
}

func main() {
	w := new()
	//http是基于tcp/Ip协议的，所以首先开启监听
	listener, err := net.Listen("tcp", w.addr)
	if err != nil {
		fmt.Println("net listen failed:", err)
		return
	}
	//启动http服务
	err = http.Serve(listener, w)
	if err != nil {
		fmt.Println("http serve failed:", err)
		return
	}
}
