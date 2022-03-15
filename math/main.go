package main

import (
	"fmt"
	"math"
)

//数学运算:

func main() {
	//保留2位小数
	//方法1：四舍五入，不会出现精确度问题;但是无法固定小数位
	n10 := math.Pow10(2)
	fmt.Println(n10)
	fmt.Println(math.Trunc((9.815+0.5/n10)*n10) / n10)              //9.82
	fmt.Println(math.Trunc((9.825+0.5/n10)*n10) / n10)              //9.83
	fmt.Println(math.Trunc((9.835+0.5/n10)*n10) / n10)              //9.84
	fmt.Println(math.Trunc((9.845+0.5/n10)*n10) / n10)              //9.85
	fmt.Println(math.Trunc((9.8003+0.5/n10)*n10) / n10)             //9.8
	fmt.Println(math.Trunc((3.3+0.5/n10)*n10) / n10)                //3.3
	fmt.Println(math.Trunc((3.3000000000000003+0.5/n10)*n10) / n10) //3.3

	fmt.Println(float64(13) / float64(11))
	fmt.Println(11 / 13)
}
