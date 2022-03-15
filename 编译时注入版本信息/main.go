package main

import "os"

var version string = "1.0"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		println(version)
	}
}

//编译时注入版本信息
// go build  -ldflags "-X main.version=1.2" -o test.exe main.go
