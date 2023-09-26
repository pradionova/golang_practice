package main

import "fmt"

func main3() {
	res := IsPowerOfTwoRecursive(64)
	fmt.Println(res)
}

func IsPowerOfTwoRecursive(n int) int {
	if n%2 != 0 {
		if n != 1 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
		return n
	}
	return IsPowerOfTwoRecursive(n / 2)
}