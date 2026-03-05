package main

import (
	"testing"
)

func Test_ValuesIntersecting(t *testing.T){

	list1 := []int { 2, 3, 2 }
	list2 := []int { 2, 1 }

	expected := []int {2, 1}

	result := findIntersectionValues(list1, list2)

	if len(result) != len(expected) {
		t.Error("should have the same lenght")
	}

	for i := 0; i < len(result); i++ {
		ev := expected[i]
		rv := result[i]

		if ev != rv {
			t.Errorf("Values are not equal. expected %d but go %d", ev, rv)
		}
	}

}