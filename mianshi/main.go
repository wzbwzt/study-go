package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(fmt.Sprintf("%.2f", float64(float64(20)/float64(50)*100)))

	return
	nums := []int{123, 2, 28, 18}
	num := sort.IntSlice(nums)
	num.Sort()
	fmt.Println(nums)
	return
	var stock float64 = -140
	res := strconv.FormatFloat(stock, 'e', 2, 64)
	fmt.Println(res)

}
