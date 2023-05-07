package pc

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func (myPC Pc) Ping() {
	fmt.Println(myPC.Brand, "Pong")
}

func (myPC *Pc) DuplicateRAM() {
	myPC.Ram = myPC.Ram * 2

}