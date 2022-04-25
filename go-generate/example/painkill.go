package main

import (
	"fmt"

	painkiller "github.com/wzbwzt/studyGo/go-generate/painkill"
)

func main() {
	fmt.Println(painkiller.Pill(7).String())
	fmt.Println(painkiller.Ibuprofen)
}
