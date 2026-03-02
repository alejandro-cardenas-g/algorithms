package main

import (
	"fmt"
	"math"
)

func main(){
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(secondLargestNumber(input))

	input = []int{8, 2, 6, 0, 3}
	fmt.Println(secondLargestNumber(input))

	input = []int{1, 1, 1, 1, 1}
	fmt.Println(secondLargestNumber(input))

	input = []int{10, 4, 0, 8, 7, 9, 5}
	fmt.Println(secondLargestNumber(input))
}

func secondLargestNumber(input []int) int{
	if len(input) < 2 {
		return 0
	}

	largest := math.MinInt
	second := math.MinInt

	for i := 0; i < len(input); i++ {
		if input[i] > largest {
			second = largest
			largest = input[i]
			continue
		}
		second = max(second, input[i])
	}

	return second
}