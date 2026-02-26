package main

import (
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"a", true},
		{"ab", false},
		{"aba", true},
		{"abba", true},
		{"abcba", true},
		{"abcdba", false},
	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %t, but got %t", test.input, test.expected, result)
		}
	}
}

var tests []struct {
	input    string
	expected string
} = []struct {
	input    string
	expected string
}{
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"a", "a"},
	{"ac", "a"},
	{"", ""},
	{"abacdfgdcaba", "aba"},
	{"abacdfgdcabba", "abba"},
}

func Test_LongestPalindromeSubstring(t *testing.T) {
	for _, test := range tests {
		result := longestPalindromeSubstring(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}

func Test_LongestPalindromeSubstrinV2(t *testing.T) {
	for _, test := range tests {
		result := longestPalindromeSubstringv2(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}
