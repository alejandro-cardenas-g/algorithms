package main

import "fmt"

func main(){
	input := []int {
		1,0,3,4,0,5,
	} 

	output := moveZerosInPlaceV2(input)
	fmt.Println(output)

	input = []int {
		0,0,0,1,2,3,
	} 

	output = moveZerosInPlaceV2(input)
	fmt.Println(output)

	input = []int {
		0,0,0,2,3,
	} 

	output = moveZerosInPlaceV2(input)
	fmt.Println(output)

	input = []int {
		0,0,1,2,3,
	} 

	output = moveZerosInPlaceV2(input)
	fmt.Println(output)
}

func moveZeros(input []int ) []int {
	if len(input) < 2 {
		return input
	}

	result := make([]int, len(input))

	lastIndex := len(input) - 1
	addAtIndex := 0

	for i := 0; i < len(input); i++ {
		if addAtIndex > lastIndex {
			return result
		}

		current := input[i]

		if current == 0 {
			result[lastIndex] = current
			lastIndex--
		}else {
			result[addAtIndex] = current
			addAtIndex++
		}
	}

	return result
}

func moveZerosInPlace(input []int ) []int {
	if len(input) < 2 {
		return input
	}

	currentAt := 0
	lastZeroSeenAt := -1
	max := len(input) -1

	for currentAt < max{
		
		first := input[currentAt]
		second := input[currentAt + 1]

		if first == 0 {
			if(second == 0 && lastZeroSeenAt == -1) {
				lastZeroSeenAt = currentAt
			}

			input[currentAt + 1] = first
			input[currentAt] = second
			// last iteration per zero
			if currentAt + 1== max {
				max--	
			}
		}

		if currentAt == max && lastZeroSeenAt != -1 {
			currentAt = lastZeroSeenAt
			lastZeroSeenAt = -1
			continue
		}
		
		currentAt++
		
	}

	return input
}

func moveZerosInPlaceV2(input []int ) []int {
	if len(input) < 2 {
		return input
	}

	numberOfZeros := 0

	valueIndexAt := 0

	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			numberOfZeros++
		}else {
			input[valueIndexAt] = input[i]
			valueIndexAt++
		}
	}

	for i := len(input) - 1; i >= len(input) - numberOfZeros; i-- {
		input[i] = 0		
	}

	return input
}