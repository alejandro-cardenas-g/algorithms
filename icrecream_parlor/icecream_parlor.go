package main

import "fmt"

func main() {
	input := []int{1, 3, 5, 6}
	m := 7

	fmt.Println(icecreamParlor(m, input))
	fmt.Println(icecreamParlorV2(m, input))

	input = []int{1, 2, 4, 5, 6}
	m = 9

	fmt.Println(icecreamParlor(m, input))
	fmt.Println(icecreamParlorV2(m, input))

}

func icecreamParlor(m int, arr []int) []int {

	left := 0
	right := len(arr) - 1

	for left <= right {

		leftV := arr[left]
		righV := arr[right]

		total := leftV + righV

		if total == m {
			return []int{left, right}
		} else if total > m {
			right--
		} else {
			left++
		}
	}

	return []int{-1, -1}
}

func icecreamParlorV2(m int, arr []int) []int {

	myMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {

		if _, ok := myMap[arr[i]]; ok {
			return []int{myMap[arr[i]], i}
		} else {
			key := m - arr[i]
			myMap[key] = i
		}
	}

	return []int{-1, -1}
}
