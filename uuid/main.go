package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	d := uuid.NewV1().String()
	fmt.Println(d)
}
