package main

import (
	"strings"
	"testing"
)

type Case struct {
	name  string
	list1 []int
	list2 []int
	expected []int
}

var cases []Case = []Case {
	{
		name: "Same length",
		list1: []int {1,2,4},
		list2: []int {2,5,6},
		expected: []int {1,2,2,4,5,6},
	},
	{
		name: "Different lenght. List 1 larger",
		list1: []int {1,2,3,4},
		list2: []int {2,5,6},
		expected: []int {1,2,2,3,4,5,6},
	},
	{
		name: "Different lenght. List 2 larger",
		list1: []int {1,2,4},
		list2: []int {2,5,6,8},
		expected: []int {1,2,2,4,5,6,8,9},
	},
} 

func Test_MergeTwoListSameLenght(t *testing.T){

	for _, x := range(cases) {
		t.Logf("%v CASE: %v %v", strings.Repeat("=", 20), x.name, strings.Repeat("=", 20))

		result := mergeTwoLists(x.list1, x.list2)

		if len(x.expected) != len(result){
			t.Errorf("result lenght is %d, but expeted %d", len(result), len(x.expected))
			t.Log(strings.Repeat("=", 100))
			continue
		}
	
		validateArrays(x.expected, result, t)
		t.Log(strings.Repeat("=", 100))
	}
}

func validateArrays(expected, result []int, t *testing.T){
	for i := 0; i < len(expected); i++ {

		expectedValue := expected[i]
		resultValue := result[i]

		if expectedValue != resultValue {
			t.Errorf("expected value %d but %d received at index %d", expectedValue, resultValue, i)
			return
		}
	}
}