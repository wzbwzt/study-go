package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)
/*
github.com/gorilla/mux:
golang自带的http.SeverMux路由实现简单,本质是一个map[string]Handler,是请求路径与该路
径对应的处理函数的映射关系。实现简单功能也比较单一：
	1.不支持正则路由， 这个是比较致命的
	2.只支持路径匹配，不支持按照Method，header，host等信息匹配，所以也就没法实现RESTful架构
而gorilla/mux是一个强大的路由，小巧但是稳定高效，不仅可以支持正则路由还可以按照Method，
header，host等信息匹配，可以从我们设定的路由表达式中提取出参数方便上层应用，而且完全兼容http.ServerMux
*/

var (
	hostname string
	port int
)

func init(){
	flag.StringVar(&hostname,"addr","0.0.0.0","service listen IP or hostname")
	flag.IntVar(&port,"port",8000,"service listen port")
}
//get 方式请求handle  //客户端请求命令：curl -X GET http://localhost:8000/api/service/get?name=wzb
func GetResquest(w http.ResponseWriter, r *http.Request){
	//解析路由中的参数
	querys := r.URL.Query()
	strings := querys["name"]
	resq:=map[string]interface{}{"Status":"OK","name":strings[0]}
	marshal, _ := json.Marshal(resq)
	w.Header().Set("Content-Type","application/json")
	w.Write(marshal)
}
//post 请求handle
func PostResquest(w http.ResponseWriter, r *http.Request) {
	//post 提交中根据route获取参数
	//curl -X POST -d  "{\"teacher\":\"wzb\"}"  http://localhost:8000/api/service/Alex/post
	//返回 map[authName:Alex]
	vars := mux.Vars(r)
	auth:=vars["authName"]
	//获取post提交体
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
	}
	s := req["teacher"].(string)
	//返回数据
	rep:= map[string]interface{}{
		"status":"OK",
		"auth":auth,
		"teacher":s,
	}
	w.Header().Set("content-Type","application/json")
	//marshal, err := json.Marshal(rep)
	//格式话输出 json
	indent, err := json.MarshalIndent(rep, "", "\t")
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found data"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(indent)
}
func main(){
	flag.Parse()
	address:=fmt.Sprintf("%s:%v",hostname,port)
	log.Println("REST service on",address)
	//register route
	route:=mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/api/service/get",GetResquest).Methods("GET")
	route.HandleFunc("/api/service/{authName}/post",PostResquest).Methods("POST")
	//开启监听
	err:=http.ListenAndServe(address,route)
	if err !=nil{
		log.Println("service listen err:",err)
	}
}




