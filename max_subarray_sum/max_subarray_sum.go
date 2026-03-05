package main

import "fmt"

func main(){
	list := []int { -1, -2, -3, -4, -5}
	fmt.Println(maxSubArraySum(list))

	list = []int { 1 }
	fmt.Println(maxSubArraySum(list))
}

func maxSubArraySum(list []int) int {

	if len(list) < 1 {
		return 0
	}

	acc := list[0]
	max := list[0]

	for i := 1; i < len(list); i++ {
		if  list[i] > acc + list[i] {
			acc = list[i]			
		} else {
			acc = acc + list[i]
		}

		if acc > max {
			max = acc
		}
	}

	return max
}