package main

import "fmt"

//defer  执行顺序按照先后后出的原则，即先入栈的最后执行
func main() {
	i := 10
	defer fmt.Printf(" defer i=%d", i) //defer 虽然是最后执行 但是会先把值压入到栈，执行是最后执行
	i = 100
	fmt.Println(i)
}
