package main

import (
	"fmt"
)

func Scanln(pp *string) {
	fmt.Println(pp)
	*pp = "rty"
}

func mainnnn() {
	var str string = "qwe"
	Scanln(&str)
	fmt.Println(str)
}