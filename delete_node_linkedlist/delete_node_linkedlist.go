package main

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func main(){
	
	head := &Node{Val:  1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	newNode := deleteNodeLinkedList(head, 2)

	curr := newNode
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}

func deleteNodeLinkedList(head *Node, deleteTerm int) *Node{

	if head == nil {
		return nil
	}

	dummy := &Node{Val: -1000}

	dummy.Next = head

	current := dummy

	for current != nil {
		
		if current.Next == nil {
			break;
		}

		if (current.Next.Val == deleteTerm){
			current.Next = current.Next.Next
			return dummy.Next
		}

		current = current.Next
	}

	return nil
}
