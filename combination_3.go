package main

import "fmt"

func main19() {
	res := Convert2D([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 2) // 2 len of the little lists
	fmt.Println(res)
}

func Convert2D(nums []int, m, n int) [][]int {
	mainList := [][]int {}
	smallList := []int{}
	for _, num := range nums {
		smallList = append(smallList, num)
		if len(smallList) == n {
			mainList = append(mainList, smallList)
			smallList = []int{}
		}
	}
	return mainList
}