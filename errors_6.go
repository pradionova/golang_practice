package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	res, _ := SumTwoIntegers("-50", "2")
	fmt.Println(res)
}

func SumTwoIntegers(a, b string) (int, error) {
	if !IsInt(a) || !
	IsInt(b) {
		return 0, errors.New("invalid input, please provide two integers")
	}
	numbera, _ := strconv.ParseInt(a, 10, 64)
	numberb, _ := strconv.ParseInt(b, 10, 64)
	//if numbera < 0 || numberb < 0 {
	//	return 0, errors.New("invalid input, please provide two integers")
	//}

	return int(numbera)+ int(numberb), nil
}

func IsInt(num string) bool {
	_, err := strconv.ParseInt(num, 10, 64)
	return err == nil
}