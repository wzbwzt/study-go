package main

import (
	"strings"
	"unicode"
)

//问题描述
//请编写⼀个⽅法，将字符串中的空格全部替换为“%20”。 假定该字符串有⾜够的空间存
//放新增的字符，并且知道字符串的真实⻓度(⼩于等于1000)，同时保证字符串由【⼤⼩
//写的英⽂字⺟组成】。 给定⼀个string为原始的串，返回替换后的string。

//解题思路
//两个问题，第⼀个是只能是英⽂字⺟，第⼆个是替换空格。

func replaceStr(str string) (string, bool) {
	if len([]rune(str)) > 1000 {
		return str, false
	}
	for _, r := range []rune(str) {
		if string(r) != " " && !unicode.IsLetter(r) {
			return str, false
		}
	}
	return strings.Replace(str, " ", "%20", -1), true
}

// func main() {
// 	fmt.Println(replaceStr("aldfjlsdkafj ajdsfsdf"))
// }
