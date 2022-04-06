package main

import (
	"fmt"
	"os"
)

var version string = "1.0"

func main() {
	args := os.Args
	if len(args) >= 2 && args[1] == "version" {
		v := utils.GetVersion()
		fmt.Printf("Version: %s\nGitBranch: %s\nCommitId: %s\nBuild Date: %s\nGo Version: %s\nOS/Arch: %s\n",
			v.Version, v.GitBranch, v.GitCommit, v.BuildDate, v.GoVersion, v.Platform)
	} else {
		fmt.Printf("Version(hard code): %s\n", "0.1")
	}
}

//编译时注入版本信息
//go build -ldflags "-X 'github.com/wzbwzt/studyGo/utils.version=0.1' -X 'github.com/wzbwzt/studyGo/utils.gitBranch=test' -X 'github.com/wzbwzt/studyGo/utils.gitTag=test' -X 'github.com/wzbwzt/studyGo/utils.gitCommit=test' -X 'github.com/wzbwzt/studyGo/utils.buildDate=2022-03-25' -X 'github.com/wzbwzt/studyGo/utils.osArch=darwin/amd64'"

/*
使用makeFile来指定注入的编译信息

*/
