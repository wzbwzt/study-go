package main

import "fmt"

//select 多路复用
//同一时刻对多个通道实现操作的场景；
/*
可处理一个或多个channel的发送/接收操作。
如果多个case同时满足，select会随机选择一个。
对于没有case的select{}会一直等待，可用于阻塞main函数。
*/

func main() {
	var ch1 = make(chan int, 1) //0,2,4,6,8;因为只有一个缓存区;如果缓存大于1的话，case随机执行；无法预测结果
	for i := 0; i < 10; i++ {
		select { //哪个case可以执行就执行哪个
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:

		}
	}

}
