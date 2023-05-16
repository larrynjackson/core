package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) printStack() {

	fmt.Print(cursor.MoveTo(ioStackMemTopRow-2, ioMemCol))
	fmt.Printf("%6s", "stack")
	core.clearPrintStack()

	var stackSize int = 65535 - int(core.RegValues[SR])
	var j uint16 = core.RegValues[SR] + 1

	for i := 0; i < stackSize; i++ {

		if j <= 65535 {
			fmt.Print(cursor.MoveTo(ioStackMemTopRow+i, ioMemCol))
			fmt.Printf("%5d", core.CoreMemory[j])
			j++
		}

	}

}

func (core *Config) clearPrintStack() {
	for i := 0; i < 20; i++ {
		fmt.Print(cursor.MoveTo(ioStackMemTopRow+i, ioMemCol))
		fmt.Printf("%6s", SP)
	}
}
