package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	root := &TreeNode{
		Val: 1,
	}

	node2 := &TreeNode{
		Val: 2,
	}

	node3 := &TreeNode{
		Val: 3,
	}

	node4 := &TreeNode{
		Val: 4,
	}

	node5 := &TreeNode{
		Val: 5,
	}

	node6 := &TreeNode{
		Val: 6,
	}

	node3.Left = node2
	node3.Right = node4
	root.Left = node3
	root.Right = node5
	node5.Right = node6

	reverseBinaryTree(root)
	fmt.Println(root)
}

func reverseBinaryTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	aux := node.Left;
	node.Left = reverseBinaryTree(node.Right)
	node.Right = reverseBinaryTree(aux)
	return node
}

