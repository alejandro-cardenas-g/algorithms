package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head

	current := dummy

	for current != nil {

		if current.Next != nil && current.Next.Val == val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}

	return head
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 6}
	node5 := &ListNode{Val: 4}
	node6 := &ListNode{Val: 5}
	node7 := &ListNode{Val: 6}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	n := removeElements(node1, 6)

	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}

	node11 := &ListNode{Val: 1}
	node22 := &ListNode{Val: 2}
	node33 := &ListNode{Val: 3}

	node11.Next = node22
	node22.Next = node33

	n = removeElements(node11, 6)

	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
