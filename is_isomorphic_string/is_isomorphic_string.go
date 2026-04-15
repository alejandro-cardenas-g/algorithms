package main

import "fmt"

func main() {

	s := "abc"
	t := "bad"

	// true
	fmt.Println(isIsomorphic(s, t))

	s = "aa12"
	t = "bb56"

	// true
	fmt.Println(isIsomorphic(s, t))

	s = "aa12"
	t = "bb55"

	// false
	fmt.Println(isIsomorphic(s, t))

}

func isIsomorphic(s, t string) bool {

	return isLeftValid(s, t) && isLeftValid(t, s)
}

func isLeftValid(s, t string) bool {
	mapper := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		currentS := s[i]
		currentT := t[i]

		if v, ok := mapper[currentS]; ok {

			if v != currentT {
				return false
			}
		} else {
			mapper[currentS] = currentT
		}
	}

	return true
}
