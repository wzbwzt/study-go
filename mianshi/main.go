package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stock float64 = -140
	res := strconv.FormatFloat(stock, 'e', 2, 64)
	fmt.Println(res)
}
