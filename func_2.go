package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mainn() {
	fmt.Println(Multiply(10, 2))
	
}

func Add(a, b float64) float64 {
	var c float64 = a + b
	 return c
}

func Multiply(a, b float64) float64 {
	var c float64 = a * b
	return c
}

func PrintNumbersAscending(n int) {
	var res string
	for i := 1; i <= n; i++ {
		res += strconv.Itoa(i)
		res += " "
	}
	res = strings.TrimSpace(res)
	fmt.Println(res)
}