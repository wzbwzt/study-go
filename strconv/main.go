package main

import (
	"fmt"
	"strconv"
)

func main1() {
	//从字符串转解析数字
	str := "123123"
	strToInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v,%T\n", strToInt, strToInt)
	//字符串转int
	atoi, _ := strconv.Atoi(str)
	fmt.Printf("%v,%T\n", atoi, atoi)

	//数字转字符串
	var i int32 = 86
	fmt.Println(string(i)) //v 是根据Asclic表来解析数字的
	str2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", str2)

	itoa := strconv.Itoa(int(i))
	fmt.Printf("%v,%T\n", itoa, itoa)
	//从字符串中解析bool
	var sb = "true"
	bs, _ := strconv.ParseBool(sb)
	fmt.Printf("%v,%T\n", bs, bs)
	//从字符串中解析浮点数
	var sf = "12.12"
	fs, _ := strconv.ParseFloat(sf, 64)
	fmt.Printf("%v,%T\n", fs, fs)


}
