package test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Gettest(){
	fmt.Println(os.Args)
	path, _ := exec.LookPath(os.Args[0])
	fmt.Println(path)
	abs, _ := filepath.Abs(path)
	fmt.Println(abs)
	index := strings.LastIndex(abs, string(os.PathSeparator))
	fmt.Println(index)
}