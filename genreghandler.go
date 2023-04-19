package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) genRegHandler(reg int, cmd string) {

	switch cmd {
	case "show":
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].hexRow, core.GenRegs[reg].valueCol))
		fmt.Printf("0x%04X", core.RegValues[reg])

		fmt.Print(cursor.MoveTo(core.GenRegs[reg].decRow, core.GenRegs[reg].valueCol))
		fmt.Printf("%05d", core.RegValues[reg])
	case "closein":
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].ingateRow, core.GenRegs[reg].ingateCol))
		fmt.Printf("%s", closeGate)
	case "closetop":
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateTopRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", closeGate)
	case "closebot":
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateBotRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", closeGate)
	case "openin":
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].ingateRow, core.GenRegs[reg].ingateCol))
		fmt.Printf("%s", openGate)
	case "opentop":
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateTopRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", openGate)
	case "openbot":
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateBotRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", openGate)
	}

}
