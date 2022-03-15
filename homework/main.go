package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	//判断一个字符串中的汉字
	s1 := "hello你好啊안녕하세요"
	var count int
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	//判断“how do you do”中单词中出现的次数
	s2 := "how do you do"
	s3 := strings.Split(s2, " ") //分割后成为切片
	m1 := make(map[string]int, 10)
	for _, v := range s3 {
		if _, ok := m1[v]; !ok { //map中的key不存在时 默认value为0
			m1[v] = 1
		} else {
			m1[v]++
		}
	}
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//回文判断：“上海自来水来自海上”
	s4 := "a上海自来水来自海上a"

	r := make([]rune, 0, len(s4))
	for _, c := range s4 {
		r = append(r, c)
	}
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

	//分金币
	left := dispatchCoin()
	fmt.Println("剩下：", left)

	//上台阶
	fmt.Println(digui(3))

}
func dispatchCoin() (left int) {
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return
}

//递归面试问题上台阶：n个台阶，一次可以走1步，一次也可以走2步，有多少走法
func digui(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return digui(n-1) + digui(n-2)
}
