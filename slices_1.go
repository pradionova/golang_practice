package main

import "fmt"

func main5() {
	res := ReverseSlice([]int{9,8,7,6,5,4,3,2,1})
	fmt.Println(res)
}

func ReverseSlice(slice []int) []int {
	length := len(slice)
	reversed := make([] int, 0)
	i := 0
	for i = length - 1; i > -1; i-- {
		reversed = append(reversed, slice[i])
	}
	return reversed
}