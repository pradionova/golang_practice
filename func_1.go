package main

import (
	"fmt"
	"math"
)

func main1() {
  SqRoots()
}

func SqRoots() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	var discriminant float64 = b*b - 4.0*a*c
	if discriminant < 0 {
		fmt.Println(0, 0)
	}
	if discriminant == 0 {
		fmt.Println(-b/(2.0*a)) //-b - math.Sqrt(discriminant) / 2*a
		}
	if discriminant > 0 {
		var dd float64 = math.Sqrt(discriminant)
		var x1 float64 = (-b - dd) / (2.0*a)
		var x2 float64 = (-b + dd) / (2.0*a)
		fmt.Println(math.Min(x1, x2), math.Max(x1, x2))
		
	}
}