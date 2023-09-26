package main

import (
	"fmt"
)

func main7() {
	res := SumOfSlice([]int {1,2,3,4})
	fmt.Println(res)
}

func SumOfSlice(slice []int) int {
	length := len(slice)
	sum := 0
	for i := 0; i < length; i++ {
		sum += slice[i]
	}
	return sum
}