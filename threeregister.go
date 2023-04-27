package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) threeRegister(count int) int {

	fmt.Print(cursor.Hide())

	var reg1Mask uint16 = 0x0700
	var reg2Mask uint16 = 0x00E0
	var reg3Mask uint16 = 0x001C

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8
	reg2 := core.RegValues[IR] & reg2Mask
	reg2 = reg2 >> 5
	reg3 := core.RegValues[IR] & reg3Mask
	reg3 = reg3 >> 2

	switch count {
	case 7:
		core.genRegHandler(int(reg2), "closetop")
		core.RegValues[B1] = core.RegValues[reg2]

		core.genRegHandler(int(reg3), "closebot")
		core.RegValues[B2] = core.RegValues[reg3]
		core.busHandler(int(B1))
		core.busHandler(int(B2))
		time.Sleep(core.SleepTime * time.Millisecond)
	case 8:

		core.RegValues[ALUL] = core.RegValues[B1]
		core.RegValues[ALUR] = core.RegValues[B2]

		core.aluHandler("left")
		core.aluHandler("right")
		core.aluHandler("code")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 9:
		core.genRegHandler(int(reg2), "opentop")
		core.genRegHandler(int(reg3), "openbot")

		if core.Opcode == "ADD" {
			core.RegValues[ACC] = core.RegValues[ALUL] + core.RegValues[ALUR]
		} else if core.Opcode == "SUB" {
			core.RegValues[ACC] = core.RegValues[ALUL] - core.RegValues[ALUR]
		} else if core.Opcode == "AND" {
			core.RegValues[ACC] = core.RegValues[ALUL] & core.RegValues[ALUR]
		} else if core.Opcode == "XOR" {
			core.RegValues[ACC] = core.RegValues[ALUL] ^ core.RegValues[ALUR]
		} else if core.Opcode == "OR" {
			core.RegValues[ACC] = core.RegValues[ALUL] | core.RegValues[ALUR]
		} else if core.Opcode == "LDW" || core.Opcode == "STW" {
			core.RegValues[ACC] = core.RegValues[ALUL] + core.RegValues[ALUR]
		}

		core.accHandler("show")
		core.accHandler("flag")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 10:

		core.accHandler("close")
		core.RegValues[DABA] = core.RegValues[ACC]
		core.RegValues[DABB] = core.RegValues[ACC]

		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		if core.Opcode == "LDW" || core.Opcode == "STW" {

			core.cntlRegHandler(int(DA), "closein")
			core.RegValues[DA] = core.RegValues[DABB]

			core.cntlRegHandler(int(DA), "show")
			time.Sleep(core.SleepTime * time.Millisecond)

		} else {
			core.genRegHandler(int(reg1), "closein")
			core.RegValues[reg1] = core.RegValues[DABA]
			core.genRegHandler(int(reg1), "show")
			time.Sleep(core.SleepTime * time.Millisecond)
		}

	case 11:

		core.accHandler("open")

		if core.Opcode == "LDW" || core.Opcode == "STW" {
			core.cntlRegHandler(int(DA), "openin")
			time.Sleep(core.SleepTime * time.Millisecond)
		} else {
			core.genRegHandler(int(reg1), "openin")
			time.Sleep(core.SleepTime * time.Millisecond)

			core.OperationClass = "fetch"
			core.clockTick(count)
			fmt.Print(cursor.Hide())
			return 1
		}
	case 12:
		if core.Opcode == "LDW" {

			core.cntlRegHandler(int(DA), "closeout")

			core.memHandler(int(DA))

			core.RegValues[DR] = core.CoreMemory[core.RegValues[DA]]
			core.memHandler(int(DR))
			time.Sleep(core.SleepTime * time.Millisecond)
		} else if core.Opcode == "STW" {
			core.genRegHandler(int(reg1), "closetop")
			core.RegValues[B1] = core.RegValues[reg1]
			core.RegValues[DABA] = core.RegValues[reg1]
			core.RegValues[DABB] = core.RegValues[reg1]
			core.busGateHandler(int(B1DAB), "close")
			core.cntlRegHandler(int(DR), "closein")
			core.cntlRegHandler(int(DR), "outdirection")
			core.busHandler(int(B1))
			core.busHandler(int(DABA))
			core.busHandler(int(DABB))
			core.RegValues[DR] = core.RegValues[reg1]
			core.cntlRegHandler(int(DR), "show")
			time.Sleep(core.SleepTime * time.Millisecond)
		}
	case 13:
		if core.Opcode == "LDW" {

			core.cntlRegHandler(int(DA), "openout")

			core.cntlRegHandler(int(DR), "closeout")
			core.cntlRegHandler(int(DR), "indirection")
			core.cntlRegHandler(int(DR), "show")

			time.Sleep(core.SleepTime * time.Millisecond)
		} else if core.Opcode == "STW" {
			core.genRegHandler(int(reg1), "opentop")
			core.busGateHandler(int(B1DAB), "open")
			core.cntlRegHandler(int(DR), "openin")
			time.Sleep(core.SleepTime * time.Millisecond)
		}
	case 14:
		if core.Opcode == "LDW" {
			core.cntlRegHandler(int(DR), "closeout")
			core.cntlRegHandler(int(DR), "indirection")
			core.cntlRegHandler(int(DR), "show")

			core.cntlRegHandler(int(DR), "openout")

			core.cntlRegHandler(int(DR), "closein")
			core.cntlRegHandler(int(DR), "indirection")

			core.RegValues[DABA] = core.RegValues[DR]
			core.busHandler(int(DABA))

			core.RegValues[DABB] = core.RegValues[DR]
			core.busHandler(int(DABB))
			time.Sleep(core.SleepTime * time.Millisecond)
		} else if core.Opcode == "STW" {

			core.cntlRegHandler(int(DA), "closeout")
			core.cntlRegHandler(int(DR), "closeout")
			core.cntlRegHandler(int(DR), "outdirection")
			core.CoreMemory[core.RegValues[DA]] = core.RegValues[DR]

			core.memHandler(int(DA))
			core.memHandler(int(DR))
			time.Sleep(core.SleepTime * time.Millisecond)
		}
	case 15:
		if core.Opcode == "LDW" {

			core.cntlRegHandler(int(DR), "openin")

			core.RegValues[reg1] = core.RegValues[DABA]

			core.genRegHandler(int(reg1), "closein")

			core.genRegHandler(int(reg1), "show")
			time.Sleep(core.SleepTime * time.Millisecond)
		} else if core.Opcode == "STW" {
			core.cntlRegHandler(int(DA), "openout")
			core.cntlRegHandler(int(DR), "openout")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.OperationClass = "fetch"
			core.clockTick(count)
			fmt.Print(cursor.Hide())
			return 1
		}
	case 16:
		if core.Opcode == "LDW" {

			core.genRegHandler(int(reg1), "openin")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.OperationClass = "fetch"
			core.clockTick(count)
			fmt.Print(cursor.Hide())
			return 1
		}
	}

	core.clockTick(count)
	fmt.Print(cursor.Hide())
	return count + 1
}
