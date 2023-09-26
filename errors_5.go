package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main25() {
	res, _ := IntToBinary(20)
	fmt.Println(res)
}

func IntToBinary(num int) (string, error) {
	i := 0
	arr := make([]int, 32) 
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}

	for num > 0 {
		arr[i] = num % 2
		num /= 2
		i += 1
	}
	str := ""
	//arr = ReverseSlice(arr)
	numStarted := false
	for _, num := range ReverseSlice(arr) {
		if !numStarted && num == 0 {
      continue
		}
		numStarted = true
		str += strconv.Itoa(num)
	}

	return str, nil
}