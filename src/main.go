package main

import "fmt"

func main() {

	// Println
	var num uint8 = 200
	fmt.Println(num)

	// Printf
	var num2 int8 = 100
	var store string = "MarketPlace"
	fmt.Printf("There are %d in the %s\n", num2, store)
	var data string = "To know the data type..."
	fmt.Printf("Data type: %T\n", data)

	// Sprintf
	message := fmt.Sprintf("Guarda el valor, no lo imprime... %d", num)
	fmt.Println(message)

}
