//go:build ingore
// +build ingore

package main

import (
	"fmt"
)

func main() {
	// done, _ := bbb()
	// done()
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}

//##############################################################################
func aaa() (done func(), err error) {
	return func() {
		print("aaa: done")
	}, nil
}

func bbb() (done func(), _ error) {
	done, err := aaa()
	return func() {
		print("bbb: surprise!")
		done() //导致递归执行，变成了一个递归函数
	}, err
}

//##############################################################################
func numSubarrayProductLessThanK(nums []int, k int) int {
	//思路1: 滑动窗口
	//参数处理
	if len(nums) == 0 || k <= 0 {
		return 0
	}

	//滑动窗口
	res := 0
	left, right, product := 0, 0, 1
	for ; right < len(nums); right++ {
		product *= nums[right]
		for left <= right && product >= k {
			product /= nums[left]
			left++
		}
		if left <= right {
			res += right - left + 1
		}
	}
	return res
}
