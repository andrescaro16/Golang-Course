package main

import "fmt"

func main() {

	// Square area
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println("Suma:", result)

	// Subtract
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplication
	result = x * y
	fmt.Println("Multiplicación:", result)

	// Division
	result = y / x
	fmt.Println("División:", result)

	// Module
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)
	
}