//go:build ignore
// +build ignore

package main

import "strings"

//大整数加法
func stringReverse(str string) string {
	reverse := []rune(str)
	strLen := len(str)
	for i, j := 0, strLen-1; i < j; i, j = i+1, j-1 {
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse)
}

func add(str1 string, str2 string) string {

	if len(str1) < len(str2) {
		str1 = strings.Repeat("0", len(str2)-len(str1)) + str1
	} else if len(str1) > len(str2) {
		str2 = strings.Repeat("0", len(str1)-len(str2)) + str2
	}
	str1 = stringReverse(str1)
	str2 = stringReverse(str2)

	count := len(str1)
	nums := make([]byte, count)
	carry := false

	for i := 0; i < count; i++ {
		sum := str1[i] - '0' + str2[i] - '0'
		if carry {
			sum++
		}
		if sum > 9 {
			sum = sum - 10
			carry = true
		} else {
			carry = false
		}
		nums[i] = sum + '0'
	}

	result := stringReverse(string(nums))
	if carry {
		result = "1" + result
	}
	return result

}

//大整数乘法
func LargeNumberMultiplication(a string, b string) (result string) {

	a = stringReverse(a)
	b = stringReverse(b)
	c := make([]byte, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c[i+j] += (a[i] - '0') * (b[j] - '0')
		}
	}

	var plus byte = 0

	for i := 0; i < len(c); i++ {
		if c[i] == 0 {
			break
		}
		temp := c[i] + plus
		plus = 0
		if temp > 9 {
			plus = temp / 10
			result += string(temp - plus*10 + '0')
		} else {
			result += string(temp + '0')
		}
	}
	return stringReverse(result)
}

func main() {

}
