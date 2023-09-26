package main

import "fmt"

func main4() {
	res := CalculateSeriesSum(3)
	fmt.Println(res)
}

func CalculateSeriesSum(n int) float64 {
	var sum float64
	i := 1
	for i = 1; i <= n; i++ {
		fraction := 1.0/float64(i)
		sum += fraction
	}
	return sum
}