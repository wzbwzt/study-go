package main

import (
	"fmt"
)

//排序方法总结
var a = []int{12, 6, 15, 22, 100, 88, 66, 3, 234, 23, 54, 14, 3, 134, 676, 9, 134, 854}

//冒泡排序
/*
比较相邻的元素。如果第一个比第二个大，就交换它们两个；
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
针对所有的元素重复以上的步骤，除了最后一个；
重复步骤1~3，直到排序完成。
*/
func BubbleSort(x []int) []int {
	for i := 0; i < len(x)-1; i++ {
		for j := 0; j < len(x)-1-i; j++ {
			if x[j] > x[j+1] {
				var tmp int
				tmp = x[j+1]
				x[j+1] = x[j]
				x[j] = tmp
			}
		}
	}
	return x
}

//选择排序（Selection Sort）
/*
工作原理：首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
以此类推，直到所有元素均排序完毕。
*/
func SelectSort(x []int) []int {
	var minIndex int
	var tmp int
	for i := 0; i < len(x)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(x); j++ {
			if x[minIndex] > x[j] {
				minIndex = j
			}
		}
		tmp = x[i]
		x[i] = x[minIndex]
		x[minIndex] = tmp
	}
	return x
}

//插入排序
func InsertSort(x []int) []int {
	for j := 1; j < len(x)-1; j++ {
		for i := j - 1; i >= 0; i-- {
			if x[j] < x[i] {
				x[i] = x[j]
			}
		}
	}
	return x
}

func main() {
	fmt.Println(BubbleSort(a))
	fmt.Println(SelectSort(a))
	fmt.Println(InsertSort(a))
}
