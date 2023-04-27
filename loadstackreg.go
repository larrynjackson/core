package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) loadSR(count int) int {

	fmt.Print(cursor.Hide())

	var reg1Mask uint16 = 0x0700

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8

	switch count {
	case 7:

		core.genRegHandler(int(reg1), "closetop")
		core.RegValues[B1] = core.RegValues[reg1]

		core.busHandler(int(B1))
		time.Sleep(core.SleepTime * time.Millisecond)

	case 8:

		core.RegValues[ALUL] = core.RegValues[B1]
		core.RegValues[ALUR] = 0

		core.aluHandler("left")
		core.aluHandler("right")
		core.aluHandler("code")
		core.genRegHandler(int(reg1), "opentop")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 9:

		core.RegValues[ACC] = core.RegValues[ALUL]

		core.accHandler("show")
		core.accHandler("flag")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 10:
		core.accHandler("close")
		core.RegValues[DABA] = core.RegValues[ACC]
		core.RegValues[DABB] = core.RegValues[ACC]

		core.busHandler(int(DABA))
		core.busHandler(int(DABB))

		core.genRegHandler(int(SR), "closein")
		core.RegValues[SR] = core.RegValues[DABA]
		core.genRegHandler(int(SR), "show")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 11:
		core.accHandler("open")
		core.genRegHandler(int(SR), "openin")
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
