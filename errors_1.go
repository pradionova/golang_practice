package main

import (
	"fmt"
	"strconv"
	//"error"
)

func main21() {
	res := ConcatStringsAndInt("hi", "hi", 1)
	fmt.Println(res)
}

func ConcatStringsAndInt(str1, str2 string, num int) string {
	return str1 + str2 + strconv.Itoa(num)
	/*result, err := str1 + str2 + num
	if err != nil {
		string(num)
			return result
	}
	fmt.Println("Результат:", result)*/
}