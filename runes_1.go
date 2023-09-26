package main

import (
	"fmt"
)

func main10() {
	res := StringLength(string("1234"))
	fmt.Println(res)
}

func StringLength(input string) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += 1
	}
	return sum
}