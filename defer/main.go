package main

import "fmt"

//defer  执行顺序按照先后后出的原则，即先入栈的最后执行
func main1() {
	i := 10
	defer fmt.Printf(" defer i=%d", i) //defer 虽然是最后执行 但是会先把值压入到栈，执行是最后执行
	i = 100
	fmt.Println(i)
}



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

func main(){
	a:=10
	b:=20
	defer func(i int) {
		calc(a,calc(a,b))
		fmt.Println(i)
	}(1000)
	a=100
}

//output:
//100 20 120
//100 120 220
//1000


func main3(){
	a:=make([]int ,6)
	a=append(a,1,2,3)
	fmt.Println(a)
}
