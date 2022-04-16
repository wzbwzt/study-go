//go:build ignore
// +build ignore

package main

import "sort"

//go container/heap  源码解析

//heap的接口类：使用时候需要实现它的方法
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

//shifDown()
//堆化(heapify):如果一个节点比它的子节点小（最大堆）或者大（最小堆），那么需要将它向下移动
func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1          //获取i这个下标所在节点的left child
		if j1 >= n || j1 < 0 { //验证是否超过堆数组范围
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		//父节点与最小（大）子节点比较，如果子节点比夫节点小(大)则替换
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

//shifUp()
//如果一个节点比它的父节点大（最大堆）或者小（最小堆），
//那么需要将它同父节点交换位置。这样是这个节点在数组的位置上升。
func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // floor((i - 1)/2):获取j的父节点
		//最小堆时，如果夫节点比子节点大，交换；
		//最大堆时，如果夫节点比子节点小,则交换
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

//初始化时，从上往下堆化处理
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

//向尾部插入节点
func Push(h Interface, x any) {
	h.Push(x)        //先执行定义的方法
	up(h, h.Len()-1) //从下而上的排序堆
}

//返回并删除根节点数据
func Pop(h Interface) any {
	n := h.Len() - 1
	h.Swap(0, n)  //将根节点与最后一个节点交换
	down(h, 0, n) //堆化处理
	return h.Pop()
}

//删除指定节点
func Remove(h Interface, i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)        //将删除的元素和最后一个元素交换
		if !down(h, i, n) { //堆化来修复，但是并不是唯一的情况，可能还需要up来向上修复
			up(h, i)
		}
	}
	return h.Pop()
}
