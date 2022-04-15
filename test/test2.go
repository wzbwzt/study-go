//go:build ignore
// +build ignore

package main

import (
	"github.com/gookit/goutil/dump"
)

/**
 * Definition for a Node.

 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	return deepCopy(head)
}

//递归
//
var nodeMap = make(map[*Node]*Node)

func deepCopy(head *Node) *Node {
	if head == nil {
		return nil
	}
	newnode, ok := nodeMap[head]
	if ok {
		return newnode
	}
	newnode = &Node{Val: head.Val}
	nodeMap[head] = newnode
	newnode.Next = deepCopy(head.Next)
	newnode.Random = deepCopy(head.Random)
	return newnode
}
func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}

	node1.Next = node2
	node1.Random = node3
	node2.Next = node3
	node2.Random = node1
	node3.Next = nil
	node3.Random = nil

	dump.P(node1)
	res := copyRandomList(node1)
	dump.P(res)

}
