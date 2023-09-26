package main

import (
	"fmt"
)

func main11() {
	res := StringLength(string("df"))
	fmt.Println(res)
}

func ConcatenateStrings(str1, str2 string) string {
	full := str1 + " " + str2
	return full
}