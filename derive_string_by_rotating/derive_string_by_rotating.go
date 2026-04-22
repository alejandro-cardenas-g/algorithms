package main

import "fmt"

func main(){
	fmt.Println(deriveStringByRotating("abcd", "bcca"))
}

func deriveStringByRotating(s string, k string) bool{
	if len(s) != len(k) {
		return false
	}

	currentLeft := k
	currentRight := k

	if currentLeft == s || currentRight == s {
		return true
	}
	
	for i := 0; i < len(k); i++ {
		first := string(currentLeft[0])

		currentLeft = fmt.Sprint(currentLeft[1:len(k)], first)

		last := string(currentRight[len(k) - 1])
		
		currentRight = fmt.Sprint(last, currentRight[0: len(k) - 1])

		if currentLeft == s || currentRight == s {
			return true
		}
	}

	return false
}

