package main

import (
	"fmt"
	"math"
)

type Shape interface {
  Area() float64
}

type Cicle struct {
  radius float64
}

func (c Cicle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main16() {
	c := Cicle{ radius: 5 }
	Culc(c)
  r := Rectangle{ width: 4, height: 8 }
	Culc(r)
}

func Culc(sh Shape) {
	fmt.Println(sh.Area())
}

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}