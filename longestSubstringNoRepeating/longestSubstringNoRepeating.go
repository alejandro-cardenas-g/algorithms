package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstringv3("caaabdcefgda"))
}

func lengthOfLongestSubstringv1(s string) int {
	if len(s) == 0 {
		return 0
	}
	left := 0
	right := 1

	max := 1
	previous := map[byte]struct{}{}
	previous[s[left]] = struct{}{}
	current := 1
	for left < right && right <= len(s)-1 {
		if _, ok := previous[s[right]]; !ok {
			previous[s[right]] = struct{}{}
			right++
			current++

			if current > max {
				max = current
			}
		} else {
			left += 1
			clear(previous)
			previous[s[left]] = struct{}{}
			right = left + 1
			current = 1
		}
	}

	return max
}

func lengthOfLongestSubstringv2(s string) int {
	left := 0
	maxStr := 0

	lastCharSeenAt := map[byte]int{}

	for right := 0; right < len(s); right++ {

		currentChar := s[right]

		if idxAt, ok := lastCharSeenAt[currentChar]; ok && idxAt >= left /*This is important as we can set left to any value lower than the left itself*/ {
			left = idxAt + 1
		}

		lastCharSeenAt[currentChar] = right

		accLength := right - left + 1

		maxStr = max(maxStr, accLength)
	}

	return maxStr
}

func lengthOfLongestSubstringv3(s string) int {
	left := 0
	maxStr := 0

	lastCharSeenAt := [128]int{}
	for i := range lastCharSeenAt {
		lastCharSeenAt[i] = -1
	}

	for right := 0; right < len(s); right++ {

		currentChar := s[right]

		if val := lastCharSeenAt[int(currentChar)]; val != -1 && val >= left /*This is important as we can set left to any value lower than the left itself*/ {
			left = val + 1
		}

		lastCharSeenAt[currentChar] = right

		accLength := right - left + 1

		maxStr = max(maxStr, accLength)
	}

	return maxStr
}
