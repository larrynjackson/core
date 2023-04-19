package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) aluHandler(cmd string) {

	switch cmd {
	case "left":
		fmt.Print(cursor.MoveTo(core.ALU.decRow, core.ALU.leftDecCol))
		fmt.Printf("%05d", core.RegValues[ALUL])
	case "right":
		fmt.Print(cursor.MoveTo(core.ALU.decRow, core.ALU.rightDecCol))
		fmt.Printf("%05d", core.RegValues[ALUR])
	case "code":
		fmt.Print(cursor.MoveTo(core.ALU.opCodeRow, core.ALU.opCodeCol))
		fmt.Printf("%6s", SP)
		fmt.Print(cursor.MoveTo(core.ALU.opCodeRow, core.ALU.opCodeCol))
		fmt.Printf("%s", core.Opcode)
	}
}
