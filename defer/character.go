//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
)

func foo() {
	defer func() {
		fmt.Println("foo defer func invoked")
	}()
	fmt.Println("foo invoked")

	bar()
	fmt.Println("do something after bar in foo")
}

func bar() {
	defer func() {
		fmt.Println("bar defer func invoked")
	}()
	fmt.Println("bar invoked")

	zoo()
	fmt.Println("do something after zoo in bar")
}

func zoo() {
	defer func() {
		fmt.Println("zoo defer func invoked")
	}()

	fmt.Println("zoo invoked")
	panic("runtime exception")
}

func main() {
	foo()
}

/*
foo invoked
bar invoked
zoo invoked
zoo defer func invoked
bar defer func invoked
foo defer func invoked
panic: runtime exception

*/

func zoo2() {
	defer func() {
		fmt.Println("zoo defer func1 invoked")
	}()

	defer func() {
		if x := recover(); x != nil {
			log.Printf("recover panic: %v in zoo recover defer func", x)
		}
	}()

	defer func() {
		fmt.Println("zoo defer func2 invoked")
	}()

	fmt.Println("zoo invoked")
	panic("zoo runtime exception")
}
