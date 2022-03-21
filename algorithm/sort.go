// +build ignore

package main

import "fmt"

//排序算法集合

//冒泡算法
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

//选择排序
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
	s := []int{2, 3, 6, 1, 98, 55, 77}
	fmt.Println(SelectSort(s))
	fmt.Println(InsertSort(s))
}
