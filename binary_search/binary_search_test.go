package main

import (
	"strings"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		target   int
		expected int
	}{
		{"found by binary", []int{1, 2, 3, 4, 5, 6}, 4, 3},
		{"not found by left edge", []int{1, 2, 3, 4, 5, 6}, -7, -1},
		{"not found by right edge", []int{1, 2, 3, 4, 5, 6}, 7, -1},
		{"not found by empty", []int{}, -7, -1},
		{"not found by binary", []int{1, 3, 5, 7, 9}, 4, -1},
	}

	for index, caseI := range cases {
		t.Run(caseI.name, func(t *testing.T) {
			t.Logf("%v case: #%d %v", strings.Repeat("=", 20), index, strings.Repeat("=", 20))
			t.Log(caseI.name)
			result := binarySearch(caseI.input, caseI.target)
			if result != caseI.expected {
				t.Errorf("expected %d but received %d", caseI.expected, result)
			}
		})
	}

}
