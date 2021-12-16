package main

import "fmt"

//节点信息
type singleNode struct {
	Value string
	Next  *singleNode
}

func (this *singleNode) String() string {
	return this.Value
}

//##############################################################################

//单向链表
type singleList struct {
	Head *singleNode //头节点：nil
}

func NewSingleList() *singleList {
	return &singleList{}
}

func (this *singleList) IsEmpty() bool {
	return this.Head == nil
}

func (this *singleList) LastNode() *singleNode {
	cur := this.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	return cur
}

func (this *singleList) Length() int {
	if this.IsEmpty() {
		return 0
	}
	count := 1
	cur := this.Head
	for cur.Next != nil {
		cur = cur.Next
		count++
	}
	return count
}

//头部插入
func (this *singleList) Unshift(value string) {
	if this == nil {
		return
	}
	this.Head = &singleNode{Value: value, Next: this.Head}
	return
}

//指定位置添加
func (this *singleList) Insert(index int, value string) {
	if this == nil {
		return
	}
	if index == 0 {
		this.Unshift(value)
		return
	}
	if index >= this.Length() {
		this.Append(value)
		return
	}

	cur := this.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	cur.Next = &singleNode{Value: value, Next: cur.Next}
	return
}

func (this *singleList) Append(v string) {
	if this == nil {
		panic("无效追加")
	}
	if this.IsEmpty() {
		this.Head = &singleNode{Value: v}
		return
	}

	this.LastNode().Next = &singleNode{Value: v}

	return
}

func main() {
	list := NewSingleList()
	list.Append("node1")
	list.Append("node2")
	list.Unshift("node3")
	list.Insert(2, "node6")

	println(list.Length())

	cur := list.Head
	for cur != nil {
		fmt.Println(cur)
		cur = cur.Next
	}

}
