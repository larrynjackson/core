package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) ioRegHandler(reg int, cmd string) {

	switch cmd {
	case "show":
		switch reg {
		case int(IN):
			fmt.Print(cursor.MoveTo(core.INPUTR.decRow, core.INPUTR.valueCol))
			fmt.Printf("%05d", core.RegValues[IN])
		case int(OUT):
			fmt.Print(cursor.MoveTo(core.OUTPUTR.decRow, core.OUTPUTR.valueCol))
			fmt.Printf("%05d", core.RegValues[OUT])
		}
	case "close":
		switch reg {
		case int(IN):
			fmt.Print(cursor.MoveTo(core.INPUTR.gateRow, core.INPUTR.gateCol))
			fmt.Printf("%s", closeGate)
		case int(OUT):
			fmt.Print(cursor.MoveTo(core.OUTPUTR.gateRow, core.OUTPUTR.gateCol))
			fmt.Printf("%s", closeGate)
		}
	case "open":
		switch reg {
		case int(IN):
			fmt.Print(cursor.MoveTo(core.INPUTR.gateRow, core.INPUTR.gateCol))
			fmt.Printf("%s", openGate)
		case int(OUT):
			fmt.Print(cursor.MoveTo(core.OUTPUTR.gateRow, core.OUTPUTR.gateCol))
			fmt.Printf("%s", openGate)
		}
	}
}
