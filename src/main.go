package main

import (
	"fmt"
	"github.com/andrescaro16/Golang-Course/figures"
)


func main() {

	square := figures.NewSquare(4)
	fmt.Println(square)

	rectangle := figures.NewRectangle(8, 4)
	fmt.Println(rectangle)

	// Before the interfaces:
	// fmt.Println(square.Area())
	// fmt.Println(rectangle.Area())

	// Now with interfaces:
	figures.Calculate(square)
	figures.Calculate(rectangle)
}