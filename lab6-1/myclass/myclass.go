package myclass

import (
	"fmt"
)

var X int = 20
var Y int
var Z int = 25

func ModFields() {
	X = 6
	Y = 7
	Z = 8
}

type myclass struct {
	abc int
}

func New() *myclass {
	return &myclass{1000}
}

func (mc *myclass) GetFields() {
	fmt.Printf("static: %v, %v, %v\ninstance: %v \n", X, Y, Z, mc.abc)
}

func (mc *myclass) SetFields() {
	X = 3
	Y = 4
	Z = 5
	mc.abc += 1000
}
