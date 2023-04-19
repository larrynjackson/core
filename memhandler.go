package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) memHandler(reg int) {

	switch reg {
	case int(DA):
		fmt.Print(cursor.MoveTo(core.MEM.daHexRow, core.MEM.memoryCol))
		fmt.Printf("0x%04X", core.RegValues[DA])

		fmt.Print(cursor.MoveTo(core.MEM.daDecRow, core.MEM.memoryCol))
		fmt.Printf("%05d", core.RegValues[DA])
	case int(DR):
		fmt.Print(cursor.MoveTo(core.MEM.drHexRow, core.MEM.memoryCol))
		fmt.Printf("0x%04X", core.RegValues[DR])

		fmt.Print(cursor.MoveTo(core.MEM.drDecRow, core.MEM.memoryCol))
		fmt.Printf("%05d", core.RegValues[DR])
	case int(IA):
		fmt.Print(cursor.MoveTo(core.MEM.iaHexRow, core.MEM.memoryCol))
		fmt.Printf("0x%04X", core.RegValues[IR])

	case int(IR):
		fmt.Print(cursor.MoveTo(core.MEM.irHexRow, core.MEM.memoryCol))
		fmt.Printf("0x%04X", core.RegValues[IR])

	}

}
