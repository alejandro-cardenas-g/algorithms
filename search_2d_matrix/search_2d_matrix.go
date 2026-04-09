package main

import "fmt"

// O(log (m)) + O(log(n)) = O(log(m * n))

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10, 11},
		{12, 13, 14, 15, 16},
	}, 3))
}

func searchMatrix(matrix [][]int, target int) bool {

	left := 0
	right := len(matrix[0]) - 1

	for left <= right {

		m := (left + right) / 2

		middleSlice := matrix[m]

		if target >= middleSlice[0] && target <= middleSlice[len(middleSlice)-1] {
			return binarySearch(middleSlice, target)
		}

		if target < middleSlice[0] {
			right = m - 1
		} else {
			left = m + 1
		}
	}

	return false
}

func binarySearch(input []int, target int) bool {

	left := 0
	right := len(input) - 1

	for left <= right {

		m := (left + right) / 2

		mV := input[m]

		if mV == target {
			return true
		}

		if target > mV {
			left = m + 1
		} else {
			right = m - 1
		}
	}

	return false

}
