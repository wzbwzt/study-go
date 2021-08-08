package main

import "fmt"

func main() {
	a := "a\r\nb"
	fmt.Println(a)
	fmt.Println(len(a))
	for _, v := range a {
		fmt.Println(string(v))
	}
}
