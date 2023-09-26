package main

import "fmt"

func main9() {
	res := SortSlice([]int {1,0,3,2})
	fmt.Println(res)
}

func SortSlice(slice []int) []int {
	length := len(slice)
	for i := 0; i < length - 1; i++ {
		for j := 0; j < length - i - 1; j++ {
			if slice[j] > slice[j + 1] {
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}
	}
	return slice
}