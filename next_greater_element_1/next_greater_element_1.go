package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	stack := []int{}
	mapper := make(map[int]int)

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			mapper[nums2[i]] = -1
		} else {
			mapper[nums2[i]] = stack[len(stack)-1]
		}

		stack = append(stack, nums2[i])
	}

	for i := 0; i < len(nums1); i++ {
		if v, ok := mapper[nums1[i]]; ok {
			nums1[i] = v
		}
	}

	return nums1
}
