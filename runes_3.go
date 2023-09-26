package main

import (
	"fmt"
	"math"
)

func main12() {
	m := map[int]int{
		-7: 1,
		-1: 38,
	}
	fmt.Println(FindMaxKey(m))
}

func FindMaxKey(m map[int]int) int {
	biggestKey := math.MinInt
	for p, _ := range m {
		biggestKey = Max2(biggestKey, p)
	}
	return biggestKey

}

func Max2(num int, num2 int) int {
	if num > num2 {
		return num
	}
	return num2
}