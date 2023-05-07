package main

import (
	"fmt"
	"github.com/andrescaro16/Golang-Course/pc"
)


func main() {

	a := 50
	b := &a				// & to access the memory address

	fmt.Println(b)
	fmt.Println(*b)		//* to access the value in that address

	*b = 100
	fmt.Println(a)

	myPC := pc.Pc{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPC)

	myPC.Ping()

	fmt.Println(myPC)
	myPC.DuplicateRAM()

	fmt.Println(myPC)
	myPC.DuplicateRAM()

	fmt.Println(myPC)

}