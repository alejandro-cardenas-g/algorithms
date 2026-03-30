package main

import "fmt"

func main(){
	fmt.Println(reverseString("hello"))
}

func reverseString(input string) string{

		output := make([]byte, len(input))

		for i := 0; i < len(input); i++{

			if i > len(input) - 1 - i {
				break;
			}

			left := input[i]
			right := input[len(input) - 1 - i];

			output[i] = right;
			output[len(input) - 1 -i] = left
		}

		return string(output)
}