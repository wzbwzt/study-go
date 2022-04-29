//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"time"
)

func takeARecvChannel() chan int {
	fmt.Println("invoke takeARecvChannel")
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		c <- 1
	}()

	return c
}

func getAStorageArr() *[5]int {
	fmt.Println("invoke getAStorageArr")
	var a [5]int
	return &a
}

func takeASendChannel() chan int {
	fmt.Println("invoke takeASendChannel")
	return make(chan int)
}

func takeBSendChannel() chan int {
	fmt.Println("invoke takeBSendChannel")
	return make(chan int)
}

func getANumToChannel() int {
	fmt.Println("invoke getANumToChannel")
	return 2
}

func getBNumToChannel() int {
	fmt.Println("invoke getBNumToChannel")
	return 2
}

func main() {
	select {
	// //send channels
	// case takeBSendChannel() <- getBNumToChannel():
	// 	fmt.Println("send somethingB to a send channel")
	//recv channels
	case (getAStorageArr())[0] = <-takeARecvChannel():
		fmt.Println("recv something from a recv channel")

	//send channels
	case takeASendChannel() <- getANumToChannel():
		fmt.Println("send somethingA to a send channel")

	}
}

/*
通过例子我们可以看出：
1) select执行开始时，首先所有case expression的表达式都会被求值一遍，按语法先后次序。

invoke takeARecvChannel
invoke takeASendChannel
invoke getANumToChannel

例外的是recv channel的位于赋值等号左边的表达式（这里是：(getAStorageArr())[0]）不会被求值。

2) 如果选择要执行的case是一个recv channel，那么它的赋值等号左边的表达式会被求值：
如例子中当goroutine 3s后向recvchan写入一个int值后，select选择了recv channel执行，
此时对=左侧的表达式 (getAStorageArr())[0] 开始求值，输出“invoke getAStorageArr”。
*/

/*
*Go 语言规范中原句：

For all the cases in the statement, the channel operands of receive operations
and the channel and right-hand-side expressions of send statements are
evaluated exactly once, in source order, upon entering the “select” statement.
The result is a set of channels to receive from or send to,
and the corresponding values to send. Any side effects in that evaluation will
occur irrespective of which (if any) communication operation is
selected to proceed. Expressions on the left-hand side of a RecvStmt with
a short variable declaration or assignment are not yet evaluated.

理解为：case语句首先判断是从chan中recv还是send to chan,
如果是recv会执行右边表达式
如果是send会先执行左边表达式再执行右边表达式
这个按照case顺序从上往下执行

最终如果走的是recv的case,才会执行左边的expression
*/
