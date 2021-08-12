package main

//问题描述
//请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字
//符串(可以使⽤单个过程变量)。
//给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于
//5000。

//问题描述
//请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字
//符串(可以使⽤单个过程变量)。
//给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于
//5000。

func reverString(str string) (string, bool) {
	if len(str) > 5000 {
		return str, false
	}
	strRune := []rune(str)

	for i := 0; i < len(strRune)/2; i++ {
		strRune[i], strRune[len(strRune)-1-i] = strRune[len(strRune)-1-i], strRune[i]
	}
	return string(strRune), true
}
