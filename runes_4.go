package main

import (
	"fmt"
)

func main13() {
	m := map[int]int{
		-7: 1,
		-1: 39,
	}
	fmt.Println(FindMaxKey(m))
}

func SumOfValuesInMap(m map[int]int) int {
	var sum int
	for _, val := range m {
		sum += val
	}
	return sum
}