package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	address string
}

func (r Person) print() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", r.name, r.age, r.address)
}

func main15() {
	p := Person{}

	p.name = "qwe"
	p.age = 1
	p.address = "wwwe"
	p.print()

}