package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width int
	Height int
}

func main(){
	//连接远程rpc服务
	reps, err := rpc.DialHTTP("tcp", "127.0.0.1:8002")
	if err != nil {
		log.Fatal(err)
	}
	reply:=0
	err = reps.Call("Rect.Area",Params{10,30}, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

