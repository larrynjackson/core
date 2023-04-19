package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) busHandler(bus int) {

	switch bus {
	case int(B1):
		fmt.Print(cursor.MoveTo(core.B1.hexRow, core.B1.hexCol))
		fmt.Printf("0x%04X", core.RegValues[B1])

		fmt.Print(cursor.MoveTo(core.B1.decRow, core.B1.decCol))
		fmt.Printf("%05d", core.RegValues[B1])
	case int(B2):
		fmt.Print(cursor.MoveTo(core.B2.hexRow, core.B2.hexCol))
		fmt.Printf("0x%04X", core.RegValues[B2])

		fmt.Print(cursor.MoveTo(core.B2.decRow, core.B2.decCol))
		fmt.Printf("%05d", core.RegValues[B2])
	case int(DABA):
		fmt.Print(cursor.MoveTo(core.DABA.hexRow, core.DABA.hexCol))
		fmt.Printf("0x%04X", core.RegValues[DABA])

		fmt.Print(cursor.MoveTo(core.DABA.decRow, core.DABA.decCol))
		fmt.Printf("%05d", core.RegValues[DABA])
	case int(DABB):
		fmt.Print(cursor.MoveTo(core.DABB.hexRow, core.DABB.hexCol))
		fmt.Printf("0x%04X", core.RegValues[DABB])

		fmt.Print(cursor.MoveTo(core.DABB.decRow, core.DABB.decCol))
		fmt.Printf("%05d", core.RegValues[DABB])
	}

}
