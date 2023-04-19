package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) busGateHandler(gate int, cmd string) {

	switch cmd {
	case "close":
		switch gate {
		case int(B1DAB):
			fmt.Print(cursor.MoveTo(core.B1DABGate.row, core.B1DABGate.col))
			fmt.Printf("%s", closeGate)
		case int(B2DAB):
			fmt.Print(cursor.MoveTo(core.B2DABGate.row, core.B2DABGate.col))
			fmt.Printf("%s", closeGate)
		}
	case "open":
		switch gate {
		case int(B1DAB):
			fmt.Print(cursor.MoveTo(core.B1DABGate.row, core.B1DABGate.col))
			fmt.Printf("%s", openGate)
		case int(B2DAB):
			fmt.Print(cursor.MoveTo(core.B2DABGate.row, core.B2DABGate.col))
			fmt.Printf("%s", openGate)
		}
	}
}
