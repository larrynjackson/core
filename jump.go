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
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[B1] = core.RegValues[reg1]
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busHandler(int(B1))
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busGateHandler(int(B1DAB), "close")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[DABA] = core.RegValues[B1]
		core.RegValues[DABB] = core.RegValues[B1]
		core.busHandler(int(DABA))
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busHandler(int(DABB))
		time.Sleep(core.SleepTime * time.Millisecond)
		core.cntlRegHandler(int(IA), "closein")
		core.cntlRegHandler(int(IA), "outdirection")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[IA] = core.RegValues[DABB]
		time.Sleep(core.SleepTime * time.Millisecond)
		core.cntlRegHandler(int(IA), "show")
	case 8:
		core.genRegHandler(int(reg1), "opentop")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busGateHandler(int(B1DAB), "open")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.cntlRegHandler(int(IA), "openin")
		core.CoreMemPoint = int16(core.RegValues[IA])

		core.OperationClass = "fetch"
		core.clockTick(count)
		fmt.Print(cursor.Hide())
		return 1
	}

	core.clockTick(count)
	fmt.Print(cursor.Hide())
	return count + 1

}
