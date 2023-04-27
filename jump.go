package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) jump(count int) int {

	fmt.Print(cursor.Hide())

	var reg1Mask uint16 = 0x0700

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8

	switch count {
	case 7:
		core.genRegHandler(int(reg1), "closetop")
		core.RegValues[B1] = core.RegValues[reg1]
		core.busHandler(int(B1))
		core.busGateHandler(int(B1DAB), "close")
		core.RegValues[DABA] = core.RegValues[B1]
		core.RegValues[DABB] = core.RegValues[B1]
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		core.cntlRegHandler(int(IA), "closein")
		core.cntlRegHandler(int(IA), "outdirection")
		core.RegValues[IA] = core.RegValues[DABB]
		core.cntlRegHandler(int(IA), "show")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 8:
		core.genRegHandler(int(reg1), "opentop")
		core.busGateHandler(int(B1DAB), "open")
		core.cntlRegHandler(int(IA), "openin")
		core.CoreMemPoint = int16(core.RegValues[IA])
		time.Sleep(core.SleepTime * time.Millisecond)

		core.OperationClass = "fetch"
		core.clockTick(count)
		fmt.Print(cursor.Hide())
		return 1
	}

	core.clockTick(count)
	fmt.Print(cursor.Hide())
	return count + 1

}
