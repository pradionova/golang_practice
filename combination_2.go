package main

import "fmt"

func main18() {
	res := FindValue([]int{2, 2, 1, 1, 1, 3})
	fmt.Println(res)
}

func FindValue(nums []int) int {
	single := make(map[int]int)
	multiple := make(map[int]int)
	for _, num := range nums {
		if _, ok := multiple[num]; ok {
			continue
		} else if _, ok := single[num]; ok {
			multiple[num] = 100
			delete(single, num)
		} else {
			single[num] = 1
			//
		}
	}

	for key, _ := range single {
		return key
	}

	return 0
}