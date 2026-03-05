package main

import "fmt"

func main(){
	list1 := []int{2,3,2}
	list2 := []int{1,2}

	fmt.Println(findIntersectionValues(list1, list2))
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
    result := []int{}

	set := make(map[int] struct{})

	for _, v := range(nums2) {
		set[v] = struct{}{}
	}

	count := 0

	for i := 0; i < len(nums1); i++ {
		if _,ok := set[nums1[i]]; ok {
			count++
		}
	}

	result = append(result, count)

	clear(set)
	count = 0
	for _, v := range(nums1) {
		set[v] = struct{}{}
	}

	for i := 0; i < len(nums2); i++ {
		if _,ok := set[nums2[i]]; ok {
			count++
		}
	}

	result = append(result, count)

	return result
}