package main

import (
	"fmt"
	"strconv"
)

const (
	eat   int = 4 //二进制是100
	sleep int = 2 //二进制是010
	da    int = 1 //二进制是001
)

//111 二级制
//最左边1 标识eat
//中间的1 标识sleep
//最右边1 标识da

func f(arg int) {
	fmt.Printf("%b", arg)

}

func main1() {
	f(eat | sleep | da) //衍生到file_write文件中的 os.OpenFile()函数中间参数flag的 |传递；决定文件是否追加+创建等多重模式的操作
}

func main() {
	fmt.Println(strconv.FormatFloat(float64(2.3), 'b', -1, 32))
}
