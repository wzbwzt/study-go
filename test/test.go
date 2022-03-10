// +build ignore
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "提交工单TS20222"
	sp := strings.Split(s, "TS")
	fmt.Printf("%v", s)
	fmt.Printf("%v", sp)
}
