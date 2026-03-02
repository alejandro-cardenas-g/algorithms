package main

import (
	"fmt"
)

func main(){
	input := "racecar"
	fmt.Println(isPalindrome(input))

	input = "hello"
	fmt.Println(isPalindrome(input))

	input = "madam"
	fmt.Println(isPalindrome(input))

	input = "raddar"
	fmt.Println(isPalindrome(input))

	input = "raddarr"
	fmt.Println(isPalindrome(input))

	input = "level"
	fmt.Println(isPalindrome(input))
}

func isPalindrome(input string) bool {
	if len(input) < 2 {
		return true
	}

	isEven := len(input) % 2

	
	left := (len(input) - 1 )/2
	if(isEven == 0){
		return IsPalindromeFromCenter(input, left, left + 1)
	}
	return IsPalindromeFromCenter(input, left, left)
}

func IsPalindromeFromCenter(input string, l, r int) bool {
	for (l <= r && r < len(input)){
		if(input[l] != input[r]){
			return false
		}
		l--
		r++
	}
	return true
}
