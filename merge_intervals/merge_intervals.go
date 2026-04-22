package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{15, 18},
		[]int{8, 10},
	}
	fmt.Println(merge(arr))
}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	merge := [][]int{}

	for i, interval := range intervals {

		if i == 0 {
			merge = append(merge, interval)
			continue
		}

		lastIndex := len(merge) - 1

		if merge[lastIndex][1] >= interval[0] {
			merge[lastIndex][1] = max(interval[1], merge[lastIndex][1])
		} else {
			merge = append(merge, interval)
		}
	}

	return merge
}
