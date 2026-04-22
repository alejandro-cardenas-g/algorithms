package main

import "fmt"

func main(){
	fmt.Println(firstUniqChar("leetcode"))
}

func firstUniqChar(s string) int {
    chars := [26]int{}

    ref := byte('a');

    for i := 0; i < len(s); i++ {
        current := s[i];
        pos := current - ref 
        chars[pos]++;
    }

    for i := 0; i < len(s); i++ {

        current := s[i];
        pos := current - ref 

        if(chars[pos] == 1){
            return i;
        }
    }

    return -1
}