package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) shift(count int) int {

	var reg1Mask uint16 = 0x0700
	var reg2Mask uint16 = 0x00E0

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8
	reg2 := core.RegValues[IR] & reg2Mask
	reg2 = reg2 >> 5
	value := core.RegValues[IR] & 0x001E
	value = value >> 1

	switch count {
	case 7:
		core.genRegHandler(int(reg2), "closetop")
		core.RegValues[B1] = core.RegValues[reg2]
		core.busHandler(int(B1))

		core.cntlRegHandler(int(IR), "closeout")
		core.RegValues[B2] = value
		core.busHandler(int(B2))
		time.Sleep(core.SleepTime * time.Millisecond)
	case 8:
		core.genRegHandler(int(reg2), "opentop")
		core.cntlRegHandler(int(IR), "openout")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 9:
		core.RegValues[ALUL] = core.RegValues[B1]
		core.RegValues[ALUR] = core.RegValues[B2]

		core.aluHandler("left")
		core.aluHandler("right")
		core.aluHandler("code")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 10:
		if core.Opcode == "SHL" {
			core.RegValues[ACC] = core.RegValues[ALUL] << core.RegValues[ALUR]
		} else if core.Opcode == "SHR" {
			core.RegValues[ACC] = core.RegValues[ALUL] >> core.RegValues[ALUR]
		}
		core.accHandler("show")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 11:
		core.accHandler("close")
		core.genRegHandler(int(reg1), "closein")
		core.RegValues[DABA] = core.RegValues[ACC]
		core.RegValues[DABB] = core.RegValues[ACC]
		core.RegValues[reg1] = core.RegValues[ACC]
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		core.genRegHandler(int(reg1), "show")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 12:
		core.accHandler("open")
		core.genRegHandler(int(reg1), "openin")

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
