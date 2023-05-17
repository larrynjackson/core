package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) printStack() {

	fmt.Print(cursor.MoveTo(ioStackMemTopRow-2, ioMemCol))
	fmt.Printf("%6s", "stack")
	core.clearPrintStack()

	var start int = int(core.RegValues[SR]) + 1
	var end int

	if 65535-int(core.RegValues[SR]) > 20 {
		end = int(core.RegValues[SR]) + 20
	} else {
		end = 65535
	}

	var j uint16 = uint16(start)
	var k int = 0

	for i := start; i < end; i++ {
		fmt.Print(cursor.MoveTo(ioStackMemTopRow+k, ioMemCol))
		fmt.Printf("%5d", core.CoreMemory[j])
		j++
		k++
	}

}

func (core *Config) clearPrintStack() {
	for i := 0; i < 20; i++ {
		fmt.Print(cursor.MoveTo(ioStackMemTopRow+i, ioMemCol))
		fmt.Printf("%6s", SP)
	}
}
