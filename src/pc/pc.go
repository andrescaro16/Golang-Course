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

func (myPC Pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.Ram, myPC.Disk, myPC.Brand)
}