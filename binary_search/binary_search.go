package main

import (
	"fmt"
)

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15}
	fmt.Println(binarySearch(input, 10))
	fmt.Println(binarySearch(input, 18))
}

func binarySearch(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	for max >= min {

		centerIndex := min + (max-min)/2

		centerValue := nums[centerIndex]

		if centerValue == target {
			return centerIndex
		}

		if centerValue > target {
			max = centerIndex - 1
		} else {
			min = centerIndex + 1
		}
	}

	return -1
}
