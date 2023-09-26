package main

import "fmt"

func main17() {
	res := MoveZeroes([]int{0, 1, 0, 3, 12}) // 2 len of the little lists
	fmt.Println(res)
}

func MoveZeroes(nums []int) []int {
	zeroCount := 0
	newList := []int{}
	for _, num := range nums {
		if num == 0 {
			zeroCount += 1
		} else {
			newList = append(newList, num)
		}
	}

	for i := 0; i < zeroCount; i++ {
		newList = append(newList, 0)
	}
	return newList
}