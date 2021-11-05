package main

import (
	"strings"
)

//问题描述
//请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允
//许使⽤额外的存储结构】。 给定⼀个string，请返回⼀个bool值,true代表所有字符全都
//不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的⻓
//度⼩于等于【3000】。

//解题思路
//这⾥有⼏个重点，第⼀个是 ASCII字符 ， ASCII字符 字符⼀共有256个，其中128个是常
//⽤字符，可以在键盘上输⼊。128之后的是键盘上⽆法找到的。
//然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使⽤额外的储存结
//构，且字符串⼩于等于3000。
//如果允许其他额外储存结构，这个题⽬很好做。如果不允许的话，可以使⽤golang内置
//的⽅式实现。

// func main() {
// 	fmt.Println(string(127))
// 	fmt.Println(string(128))
// }

//第⼀个⽅法使⽤的是golang内置⽅法 strings.Count ,可以⽤来判断在⼀个字符串中包含
//的另外⼀个字符串的数量。
func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

//第⼆个⽅法使⽤的是golang内置⽅法 strings.Index 和 strings.LastIndex ，⽤来判断指
//定字符串在另外⼀个字符串的索引未知，分别是第⼀次发现位置和最后发现位置。
func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
