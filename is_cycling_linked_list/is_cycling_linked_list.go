package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	head := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node2

	fmt.Println(isClyclingLInkedList(head))
}

func isClyclingLInkedList(node *Node) bool {

	if node == nil {
		return false
	}

	slow := node
	fast := node.Next

	for fast != nil {

		if fast == slow {
			return true
		}

		slow = slow.Next

		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

	}

	return false
}
