package main

import (
	"fmt"
)

//defer  执行顺序按照先后后出的原则，即先入栈的最后执行
func main1() {
	i := 10
	defer fmt.Printf("defer i=%d", i) //defer 虽然是最后执行 但是会先把值压入到栈，执行是最后执行
	i = 100
	fmt.Println(i)
}
/*
output:
100
defer i=10
*/


func calc(x,y int )int{
	res:=x+y
	fmt.Println(x,y,res)
	return res
}

func main2(){
	a:=10
	b:=20
	defer calc(a,calc(a,b))
	a=100
	defer calc(a,calc(a,b))
	a=1000

}
//output:
//10 20 30
//100 20 120
//100 120 220
//10 30 40

func main4(){
	a:=10
	b:=20
	defer func(i int) {
		calc(a,calc(a,b))
		fmt.Println(i)
	}(1000)
	a=100
}

//闭包的原理：a会从函数外面找最新的那个值，这个区别于main2中的压栈执行（先将a传进去，等待执行）；
//output:
//100 20 120
//100 120 220
//1000


func main3(){
	a:=make([]int ,6)
	a=append(a,1,2,3)
	fmt.Println(a)
}




