package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productOfArrayExceptSelf(nums))
}

/* Must be done in O(n) time and O(1) space without using division */

func productOfArrayExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	// the acc is the product of the numbers before the current index
	acc := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			acc = 1
		} else {
			acc = acc * nums[i-1]
		}
		result[i] = acc
	}

	// the acc is the product of the numbers after the current index
	acc = 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			acc = 1
		} else {
			acc = acc * nums[i+1]
		}
		result[i] = result[i] * acc
	}
	return result
}
