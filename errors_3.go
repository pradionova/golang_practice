package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main23() {
	res, _ := GetCharacterAtPosition("Зенит чемпион!", 20)
	fmt.Println(res)
}

func GetCharacterAtPosition(str string, position int) (rune, error) {
	length := utf8.RuneCountInString(str)
	if position > length || position < 0 {
		return 0, errors.New("0")
	}
	return []rune(str)[position - 1], nil
}