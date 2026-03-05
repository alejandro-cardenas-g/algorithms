package main

import "testing"

func Test_MaxSubarraySum(t *testing.T){
	cases := []struct{
		name string
		input []int
		expected int
	}{
		{name: "empty", input: []int{}, expected: 0},
		{name: "single", input: []int{1}, expected: 1},
		{name: "multiple", input: []int{1, 2, 3, 4, 5}, expected: 15},
		{name: "negative", input: []int{-1, -2, -3, -4, -5}, expected: -1},
		{name: "mixed", input: []int{1, -2, 3, -4, 5}, expected: 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := maxSubArraySum(tc.input)
			if result != tc.expected {
				t.Errorf("expected %d but got %d", tc.expected, result)
			}
		})
	}
}