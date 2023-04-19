package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) instructionHandler(cmd string) {

	switch cmd {
	case "show":
		fmt.Print(cursor.MoveTo(core.Instruction.row, core.Instruction.col))
		fmt.Printf("%14s", SP)
		fmt.Print(cursor.MoveTo(core.Instruction.row, core.Instruction.col))
		fmt.Printf("%s", core.InstructionString)
	}
}
