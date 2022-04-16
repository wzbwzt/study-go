package main

import (
	"container/heap"
	"fmt"

	"github.com/gookit/goutil/dump"
)

type IntHeap []int

func main() {
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0} // 创建slice
	heap.Init(h)                                // 将数组切片进行堆化
	dump.P(h)                                   // [0 1 3 6 2 5 7 9 8 4] 由Less方法可控制小顶堆
	dump.P(heap.Pop(h))                         // 调用pop 0 返回移除的顶部最小元素
	heap.Push(h, 6)                             // 调用push [1 2 3 6 4 5 7 9 8] 添加一个元素进入堆中进行堆化
	dump.P("new", h)                            // [1 2 3 6 4 5 7 9 8 6]
	for len(*h) > 0 {                           // 持续推出顶部最小元素
		fmt.Printf("%d \n ", heap.Pop(h))
	}
}

func (h IntHeap) Len() int { return len(h) }

// 这里决定 大小顶堆 现在是小顶堆
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	dump.P("old", old) // [1 2 3 6 4 5 7 9 8 0] 将顶小堆元素与最后一个元素交换位置，在进行堆排序的结果

	x := old[n-1]
	*h = old[0 : n-1]
	dump.P(h) // [1 2 3 6 4 5 7 9 8]
	return x
}

func (h *IntHeap) Push(x any) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}
