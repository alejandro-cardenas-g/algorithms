package main

import "fmt"

func main() {
	input := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	fmt.Println(floodFill(input, 1, 1, 2))
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	visited := make(map[string]*struct{})

	initialVal := image[sr][sc]

	dfs(image, sr, sc, color, initialVal, visited)

	return image
}

func dfs(image [][]int, sr int, sc int, color int, initVal int, visited map[string]*struct{}) {

	if sr < 0 || sr > len(image)-1 || sc < 0 || sc > len(image[0])-1 {
		return
	}

	key := fmt.Sprintf("%d:%d", sr, sc)

	if _, ok := visited[key]; ok {
		return
	}

	visited[key] = nil

	if image[sr][sc] == initVal {
		image[sr][sc] = color
		dfs(image, sr, sc-1, color, initVal, visited)
		dfs(image, sr+1, sc, color, initVal, visited)
		dfs(image, sr, sc+1, color, initVal, visited)
		dfs(image, sr-1, sc, color, initVal, visited)
	}
}
