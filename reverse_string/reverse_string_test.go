package main

import "testing"

func Test_ReverseString(t *testing.T){
	cases := []struct{
		name string
		input string
		expected string
	}{
		{name: "empty", input: "", expected: ""},
		{name: "single", input: "a", expected: "a"},
		{name: "multiple", input: "abc", expected: "cba"},
		{name: "multiple", input: "abcd", expected: "dcba"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := reverseString(tc.input)
			if result != tc.expected {
				t.Errorf("expected %s but got %s", tc.expected, result)
			}
		})
	}
}