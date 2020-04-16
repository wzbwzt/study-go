package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机数
func main() {
	rand.Seed(time.Now().UnixNano()) //随机数种子(必需的) 要不然下面执行的每次随机数都是一样的
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    //int64的随机数
		r2 := rand.Intn(10) //0<=r2<10
		fmt.Println(r1, r2)
	}

}
