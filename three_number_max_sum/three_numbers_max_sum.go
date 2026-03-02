package main

import "fmt"

func main(){
	nums := []int{5,2,3,4,1,6,7,11,0,22,0,0,1,1,1,12,1}
	fmt.Println(threeNumbersMaxSum(nums))
}

func threeNumbersMaxSum(nums []int) [3]int{
	maxSum := 0;

	result := [3]int{0, 0, 0};

	min := 0;
	maxI := 2;

	for (min < maxI && maxI < len(nums)){

		first := nums[min];
		second := nums[min + 1]
		third := nums[maxI];


		fmt.Println(first, second, third);
		currentSum := first + second + third

		fmt.Println(currentSum);
		if(currentSum > maxSum){
			maxSum = currentSum;
			result[0] = first
			result[1] = second;
			result[2] = third;
		}
		
		min++;
		maxI++;
	}

	return result;
}