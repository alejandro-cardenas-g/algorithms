package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	head := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node4 := &Node{Val: 4}

	head.Next = node2
	node2.Next = node4

	head2 := &Node{Val: 1}
	node22 := &Node{Val: 3}
	node23 := &Node{Val: 4}
	node25 := &Node{Val: 5}

	head2.Next = node22
	node22.Next = node23
	node23.Next = node25

	newNode := mergeTwoLists(head, head2)

	curr := newNode
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func mergeTwoLists(list1 *Node, list2 *Node) *Node {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	dummy := &Node{}

	ref1 := list1
	ref2 := list2

	current := dummy

	for ref1 != nil && ref2 != nil {
		if ref1.Val <= ref2.Val {
			current.Next = ref1
			ref1 = ref1.Next
		} else {
			current.Next = ref2
			ref2 = ref2.Next
		}
		current = current.Next
	}

	if ref1 == nil {
		current.Next = ref2
	}

	if ref2 == nil {
		current.Next = ref1
	}

	return dummy.Next
}
