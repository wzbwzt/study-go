//go:build ignore
// +build ignore

package main

import "fmt"

//链表相关

type Node struct {
	value int
	next  *Node
}

//链表反转
func reverse(head *Node) *Node {

	var pre *Node = nil
	for head != nil {
		temp := head.next
		head.next = pre
		pre = head
		head = temp
	}
	return pre
}

func printNode(head *Node) {
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

//单链表是否存在环
func hasCycle(head *Node) bool {
	fast := head
	slow := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return true
		}
	}
	return false
}
