//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	var nodes = &ListNode{Val: 1}
	nodes.Next = &ListNode{Val: 2}
	nodes.Next.Next = &ListNode{Val: 2}
	nodes.Next.Next.Next = &ListNode{Val: 1}
	fmt.Println(isPalindrome(nodes))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}
