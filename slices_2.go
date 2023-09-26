package main

import (
	"fmt"
)

func main6() {
	var numbers = []int{}
	min, max := FindMinMaxInSlice(numbers)
	fmt.Println(min, max)
}

func FindMinMaxInSlice (slice []int) (int, int) {
	length := len(slice)
	if length == 0 {
		return 0,0
	}
	min := slice[0]
	max := slice[0]
	for i := 1; i < length; i++ {
		min = Min(min, slice[i])
		max = Max(max, slice[i])
		
	}
	return min, max
}

func Min(a, b int) int {
	if a < b {
			return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
			return a
	}
	return b
}