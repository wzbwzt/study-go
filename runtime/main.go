package main

import (
	"fmt"
	"path"
	"runtime"
)

func main(){
	name, funcName, line := f2(0)
	fmt.Printf("file:%v;function:%v;line:%d",name,funcName,line)
}


func getLocation(skip int)(fileName ,funcName string ,line int){
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("get info failed")
		return
	}
	fmt.Println(pc,file)
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}




