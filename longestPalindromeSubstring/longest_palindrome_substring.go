package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubstringv2("abacda"))
}

// v1 brute force

func longestPalindromeSubstring(s string) string {

	if len(s) < 2 {
		return s
	}

	left := 0
	right := 1

	maxPalindrome := string(s[left])
	accStr := string(s[left])

	for left < right {
		accStr += string(s[right])
		if isPalindrome(accStr) {
			if len(accStr) > len(maxPalindrome) {
				maxPalindrome = accStr
			}
		}

		if right == len(s)-1 {
			left++
			right = min(left+1, len(s)-1)
			accStr = string(s[left])
			continue
		}

		right++
	}

	return maxPalindrome
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		right := len(s) - i - 1
		if i > right {
			break
		}
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// v2 expanding from the center

func longestPalindromeSubstringv2(s string) string {
	if len(s) < 2 {
		return s
	}

	maxPalindrome := ""
	for i := 0; i < len(s)-1; i++ {

		// 	left, right := i, i
		// 	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		// 		lengOfCurrentPalindrome := right - left + 1
		// 		if len(maxPalindrome) < lengOfCurrentPalindrome {
		// 			maxPalindrome = s[left : right+1]
		// 		}
		// 		left--
		// 		right++
		// 	}

		// 	left, right = i, i+1
		// 	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		// 		lengOfCurrentPalindrome := right - left + 1
		// 		if len(maxPalindrome) < lengOfCurrentPalindrome {
		// 			maxPalindrome = s[left : right+1]
		// 		}
		// 		left--
		// 		right++
		// 	}

		newMax := getMaxPalindromePositions(s, i, i)
		if newMax > len(maxPalindrome) {
			maxPalindrome = s[i-newMax/2 : i+newMax/2+1]
		}

		newMax = getMaxPalindromePositions(s, i, i+1)
		if newMax > len(maxPalindrome) {
			maxPalindrome = s[i-(newMax-1)/2 : (i+1)+(newMax-1)/2+1]
		}
	}
	return maxPalindrome
}

func getMaxPalindromePositions(s string, left, right int) int {
	lengthOfCurrentPalindrome := 0
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		lengthOfCurrentPalindrome = right - left + 1
		left--
		right++
	}
	return lengthOfCurrentPalindrome
}
