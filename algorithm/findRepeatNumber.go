//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{2, 3, 1, 0, 2, 5, 3}
	res := findRepeatNumber(test)
	fmt.Println(res)
}

func findRepeatNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		left := 1
		right := n - 1
		nums[0], nums[i] = nums[i], nums[0]
		for left <= right {
			m := (left + right) / 2
			if nums[m] == nums[0] {
				return nums[0]
			}
			if nums[m] > nums[0] {
				right = m - 1
			} else {
				left = m + 1
			}
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return 0

}
