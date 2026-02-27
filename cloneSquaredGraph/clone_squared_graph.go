package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 * Val int
 * Neighbors []*Node
 * }
 */

func main() {
	fmt.Println("clone squared graph")
	node := &Node{Val: 1, Neighbors: []*Node{}}
	node2 := &Node{Val: 2, Neighbors: []*Node{}}
	node3 := &Node{Val: 3, Neighbors: []*Node{}}
	node4 := &Node{Val: 4, Neighbors: []*Node{}}
	node.Neighbors = append(node.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node3, node)
	node3.Neighbors = append(node3.Neighbors, node4, node2)
	node4.Neighbors = append(node4.Neighbors, node, node3)

	fmt.Println(cloneGraph(node))
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var root *Node = &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	if len(node.Neighbors) < 1 {
		return root
	}

	visited := make(map[int]*Node)

	addNeighbors(node, root, root, visited)

	return root
}

func addNeighbors(node *Node, newNode *Node, root *Node, visited map[int]*Node) {
	if node.Val == root.Val && len(root.Neighbors) == len(node.Neighbors) {
		return
	}

	if _, isProcessing := visited[node.Val]; isProcessing {
		return
	}

	if node.Val == root.Val {
		visited[node.Val] = root
	} else {
		visited[node.Val] = newNode
	}

	for _, neigbhor := range node.Neighbors {

		var nextNode *Node
		if val, ok := visited[neigbhor.Val]; ok {
			nextNode = val
		} else {
			nextNode = &Node{
				Val:       neigbhor.Val,
				Neighbors: []*Node{},
			}
		}

		newNode.Neighbors = append(newNode.Neighbors, nextNode)
		addNeighbors(neigbhor, nextNode, root, visited)
	}
}
