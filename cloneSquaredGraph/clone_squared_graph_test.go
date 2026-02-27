package main

import "testing"

func Test_CloneSquaredGraph(t *testing.T) {
	node := &Node{Val: 1, Neighbors: []*Node{}}
	node2 := &Node{Val: 2, Neighbors: []*Node{}}
	node3 := &Node{Val: 3, Neighbors: []*Node{}}
	node4 := &Node{Val: 4, Neighbors: []*Node{}}
	node.Neighbors = append(node.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node3, node)
	node3.Neighbors = append(node3.Neighbors, node4, node2)
	node4.Neighbors = append(node4.Neighbors, node, node3)

	result := cloneGraph(node)

	compareNodes(node, result, make(map[int]bool), t)
}

func Test_CloneSquaredGraph_EmptyGraph(t *testing.T) {
	node := &Node{Val: 1, Neighbors: []*Node{}}

	result := cloneGraph(node)

	compareNodes(node, result, make(map[int]bool), t)
}

func compareNodes(node1 *Node, node2 *Node, visited map[int]bool, t *testing.T) bool {
	if node1.Val != node2.Val {
		t.Errorf("Expected %d, got %d", node1.Val, node2.Val)
		return false
	}

	if _, ok := visited[node1.Val]; ok {
		return true
	}

	visited[node1.Val] = true

	if len(node1.Neighbors) != len(node2.Neighbors) {
		t.Errorf("Expected %d neighbors, got %d", len(node1.Neighbors), len(node2.Neighbors))
		return false
	}

	for i, neighbor := range node1.Neighbors {
		if !compareNodes(neighbor, node2.Neighbors[i], visited, t) {
			t.Errorf("Expected %d, got %d", neighbor.Val, node2.Neighbors[i].Val)
			return false
		}
	}
	return true
}
