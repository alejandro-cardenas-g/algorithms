package main

import "fmt"

func main(){
	list1 := []int{1,2,4,6,7,8}
	list2 := []int{1,3,4,5}

	fmt.Println(mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 []int, list2 []int) []int {

	if len(list1) == 0 {
		return list2
	}
	if len(list2) == 0 {
		return list1
	}

	result := make([]int, len(list1) + len(list2))

	resultIndex := 0

	p1 := 0
	p2 := 0

	for p1 < len(list1) && p2 < len(list2){
		if list1[p1] > list2[p2] {
			result[resultIndex] = list2[p2]
			p2++
		}else {
			result[resultIndex] = list1[p1]
			p1++
		}

		resultIndex++
	}

	if p1 == len(list1) {
		for i := p2; i < len(list2); i++ {
			result[resultIndex] = list2[i]
			resultIndex++
		}
	}

	if p2 == len(list2) {
		for i := p1; i < len(list1); i++ {
			result[resultIndex] = list1[i]
			resultIndex++
		}
	}

	return result
}