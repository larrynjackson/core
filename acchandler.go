package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) accHandler(cmd string) {

	switch cmd {
	case "show":
		fmt.Print(cursor.MoveTo(core.ACC.decRow, core.ACC.decCol))
		fmt.Printf("%05d", core.RegValues[ACC])
	case "flag":
		fmt.Print(cursor.MoveTo(core.ACC.flagsRow, core.ACC.flagsCol))
		fmt.Printf("%11s", SP)
		var eq, ne, lt, gt uint16
		if core.RegValues[FLAG] == 1 {
			eq = 1
		} else if core.RegValues[FLAG] == 6 {
			ne = 1
			lt = 1
		} else if core.RegValues[FLAG] == 10 {
			ne = 1
			gt = 1
		}
		fmt.Print(cursor.MoveTo(core.ACC.flagsRow, core.ACC.flagsCol))
		fmt.Printf("%d%2s%d%2s%d%2s%d", gt, SP, lt, SP, ne, SP, eq)
	case "close":
		fmt.Print(cursor.MoveTo(core.ACC.busRow, core.ACC.busCol))
		fmt.Printf("%s", BCU)

		fmt.Print(cursor.MoveTo(core.ACC.boxBusRow, core.ACC.busCol))
		fmt.Printf("%s", GBP)
	case "open":
		fmt.Print(cursor.MoveTo(core.ACC.busRow, core.ACC.busCol))
		fmt.Printf("%s", HP)

		fmt.Print(cursor.MoveTo(core.ACC.boxBusRow, core.ACC.busCol))
		fmt.Printf("%s", HB)
	}
}
