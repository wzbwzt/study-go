//go:build ignore
// +build ignore

package main

//一个数组里只有一个数出现一次，其他元素都出现3次，在不允许使用辅助空间，时间复杂度位O(n) 的情况下找到它
func singleNumber(arr []int) int {
	count := len(arr)
	ones := 0
	twos := 0
	xthrees := 0
	for i := 0; i < count; i++ {
		twos |= ones & arr[i]
		ones ^= arr[i]
		xthrees = ^(ones & twos)
		ones &= xthrees
		twos &= xthrees
	}
	return ones
}

//一个数组超过一半以上都是同一个数，求这个数
func findMostNum(nums []int) int {

	candidate := 0
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else {
			if candidate == nums[i] {
				count++
			} else {
				count--
			}
		}
	}
	return candidate
}
