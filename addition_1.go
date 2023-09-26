package main

import (
	"fmt"
	"strings"
)

func main() {
	res := MoveZeroes([]string{"hi"}, {"ih"}) // 2 len of the little lists
	fmt.Println(res)
}

func IsPalindrome(input string) bool {
	var reversed = input
	input = input.Tolower()
	input = input.reverse()
	if reversed == input {
		return true
	}
	return false
}

func reverse(s string) string {
    rns := []rune(s) // convert to rune
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
  
        // swap the letters of the string,
        // like first with last and so on.
        rns[i], rns[j] = rns[j], rns[i]
    }
  
    // return the reversed string.
    return string(rns)
}