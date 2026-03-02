package main

import "fmt"

func main(){
	nums := []int{7,1,5,3,6,4,0,10, 2}
	fmt.Println(bestProfitStock(nums))
}

func bestProfitStock(nums []int) int{
	left := 0
	right := len(nums) - 1

	maxProfit := 0;

	for (left < right){

		profit := nums[right] - nums[left];

		maxProfit = max(maxProfit, profit)

		if(right == len(nums) - 1){
			left++;
			right = min(left + 1, len(nums) - 1);
		} else {
			right++;
		}
	}

	return maxProfit
}