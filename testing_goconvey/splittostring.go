package splittostring

import "strings"

//Splittostring  测试testing 包
func Splittostring(s, sep string) (res []string) {
	res = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i >= 0 {
		res = append(res, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	res = append(res, s)
	return
}

//测试性能比较：比较同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
