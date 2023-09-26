package main

import (
	"fmt"
)

type Animal interface {
  MakeSound() string
}

type Dog struct {

}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")
}

type Cat struct {

}

func (c Cat) MakeSound() {
	fmt.Println("Мяу!")
}