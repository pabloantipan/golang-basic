package structs

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Brand string
	Year  int
}

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

func Structing() interface{} {
	myCar := Car{
		Brand: "Porsche",
		Year:  2022,
	}

	return myCar
}

func (pc Pc) PingPc() (string, string) {
	return pc.Brand, "Pong!"
}

func (pc *Pc) DuplicateRAM() {
	pc.Ram = pc.Ram * 2
}

func (pc Pc) Stringfy() string {
	out, err := json.Marshal(pc)
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (pc Pc) String() string {
	return fmt.Sprintf("I have %d GB RAM, %d GB Disk and I am %s", pc.Ram, pc.Disk, pc.Brand)
}
