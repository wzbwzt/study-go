//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	res := targetIndices([]int{1, 2, 5, 2, 3}, 2)
	fmt.Println("res:", res)

}

//获取指定下标值
func targetIndices(nums []int, target int) []int {
	n := len(nums) - 1
	mergeSort(nums, 0, n)

	l, r, mid := 0, n, 0
	for l <= r {
		mid = (r-l)>>1 + l
		fmt.Println(mid)
		if nums[mid] < target {
			l = mid + 1
		} else {
			if nums[mid] == target && (mid-1 < 0 || nums[mid-1] < target) {
				break
			}
			r = mid - 1
		}
	}

	res := make([]int, 0)
	for i := mid; i < len(nums); i++ {
		if nums[i] == target {
			res = append(res, i)
		} else {
			break
		}
	}
	return res
}

func mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r-l)>>1 + l
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)

	merge(nums, l, mid, r)
}

func merge(nums []int, l, mid, r int) {
	p1, p2 := l, mid+1
	res := make([]int, 0)
	for p1 <= mid && p2 <= r {
		if nums[p1] < nums[p2] {
			res = append(res, nums[p1])
			p1++
		} else {
			res = append(res, nums[p2])
			p2++
		}
	}
	for p1 <= mid {
		res = append(res, nums[p1])
		p1++
	}
	for p2 <= r {
		res = append(res, nums[p2])
		p2++
	}
	for i := range res {
		nums[i+l] = res[i]
	}
}
