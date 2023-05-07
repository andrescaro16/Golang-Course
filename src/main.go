// 1. "go mod init <module-name>"
// 2. Create the folder with the .go file
// 3. "go mod tidy" to update the go.mod

package main

import (
	"fmt"

	"github.com/andrescaro16/Golang-Course/transport"
)

func main() {

	var myCar transport.Car
	myCar.Brand = "Tesla"
	fmt.Println(myCar)

	transport.PrintMessage("Hey")

}
