package main

import "fmt"

func main() {
	nums := []int{-1, -3, -2, 6, -4, -2, 3, 2}
	fmt.Println(maxSubArray3(nums))
}

func maxSubArray(nums []int) int {

	right := 1
	var maxSum int = nums[0]
	var acc int = nums[0]

	for right <= len(nums)-1 {
		if nums[right] >= acc {
			acc = nums[right]
		} else {
			acc = acc + nums[right]
		}
		right++

		if acc > maxSum {
			maxSum = acc
		}
	}
	return maxSum
}

func maxSubArray2(nums []int) int {
	var maxSum int = nums[0]
	var acc int = 0

	for _, s := range nums {
		if acc < 0 {
			acc = 0
		}
		acc += s
		maxSum = max(maxSum, acc)
	}
	return maxSum
}

func maxSubArray3(nums []int) int {
	var maxSum int = nums[0]
	var acc int = 0

	for _, s := range nums {
		if s >= acc {
			acc = s
		} else {
			acc = acc + s
		}

		maxSum = max(maxSum, acc)

	}
	return maxSum
}
