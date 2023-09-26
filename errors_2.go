package main

import (
	"errors"
	"fmt"
)

func main22() {
	res, _ := DivideIntegers(3, 2)
	fmt.Println(res)
}

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return float64(0.0), errors.New("division by zero")
	}
	return float64(a) / float64(b), nil
}