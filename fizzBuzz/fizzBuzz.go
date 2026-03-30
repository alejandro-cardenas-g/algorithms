package main

import "fmt"

func main(){
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {

    output := []string {}
    for i:=1; i <= n;i++ {

        divisibleBy3 := i%3 == 0
        divisibleBy5 := i%5 == 0

        newItem := fmt.Sprintf("%d", i);
        if divisibleBy3 && divisibleBy5 {
            newItem = "FizzBuzz"
        } else if divisibleBy3 {
            newItem = "Fizz"    
        } else if divisibleBy5 {
            newItem = "Buzz"
        }

        output = append(output, newItem)
    }

    return output
}