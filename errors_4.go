package main

import (
	"errors"
	"fmt"
)

func main24() {
	res, _ := Factorial(3)
	fmt.Println(res)
}

func Factorial(n int) (int, error) {
	factorial := 1
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial, nil
}
