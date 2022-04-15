//go:build ignore
// +build ignore

package main

import "fmt"

//排序算法集合

//冒泡算法
//O(n^2)
func BubbleSort(x []int) []int {
	for i := 0; i < len(x)-1; i++ {
		for j := 0; j < len(x)-1-i; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	return x
}

//选择排序
//O(n^2)
func SelectSort(x []int) []int {
	var minIndex int
	for i := 0; i < len(x)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(x); j++ {
			if x[minIndex] > x[j] {
				minIndex = j
			}
		}
		x[i], x[minIndex] = x[minIndex], x[i]
	}
	return x
}

//插入排序
//O(n^2)
func InsertSort(x []int) []int {
	for j := 1; j < len(x)-1; j++ {
		for i := j - 1; i >= 0; i-- {
			if x[i+1] < x[i] {
				x[i], x[i+1] = x[i+1], x[i]
			}
		}
	}
	return x
}

func main() {
	s := []int{2, 3, 6, 1, 98, 55, 77}
	fmt.Println(BubbleSort(s))
	// fmt.Println(SelectSort(s))
	// fmt.Println(InsertSort(s))
}
