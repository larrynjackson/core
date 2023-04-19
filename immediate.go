package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) immediate(count int) int {

	fmt.Print(cursor.Hide())

	var reg1Mask uint16 = 0x0700

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8
	value := core.RegValues[IR] & 0x00FF

	switch count {
	case 7:

		if core.Opcode != "LDI" {
			core.genRegHandler(int(reg1), "closetop")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.RegValues[B1] = core.RegValues[reg1]

			core.busHandler(int(B1))
			time.Sleep(core.SleepTime * time.Millisecond)
		}

		core.cntlRegHandler(int(IR), "closeout")
		core.RegValues[B2] = value
		core.busHandler(int(B2))
		time.Sleep(core.SleepTime * time.Millisecond)

	case 8:
		if core.Opcode == "LDI" {
			core.RegValues[ALUL] = 0
			core.RegValues[ALUR] = core.RegValues[B2]
			time.Sleep(core.SleepTime * time.Millisecond)
			core.cntlRegHandler(int(IR), "openout")

		} else {
			core.RegValues[ALUL] = core.RegValues[B1]
			core.RegValues[ALUR] = core.RegValues[B2]
			time.Sleep(core.SleepTime * time.Millisecond)

			core.genRegHandler(int(reg1), "opentop")
			core.cntlRegHandler(int(IR), "openout")
		}

		core.aluHandler("left")
		core.aluHandler("right")
		core.aluHandler("code")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 9:

		if core.Opcode == "ADDI" {
			core.RegValues[ACC] = core.RegValues[ALUL] + core.RegValues[ALUR]
		} else if core.Opcode == "SUBI" {
			core.RegValues[ACC] = core.RegValues[ALUL] - core.RegValues[ALUR]
		} else if core.Opcode == "ANDI" {
			core.RegValues[ACC] = core.RegValues[ALUL] & core.RegValues[ALUR]
		} else if core.Opcode == "XORI" {
			core.RegValues[ACC] = core.RegValues[ALUL] ^ core.RegValues[ALUR]
		} else if core.Opcode == "ORI" {
			core.RegValues[ACC] = core.RegValues[ALUL] | core.RegValues[ALUR]
		} else if core.Opcode == "LDI" {
			core.RegValues[ACC] = core.RegValues[ALUR]
		} else if core.Opcode == "CMPI" {
			var flags uint16 = 0

			if core.RegValues[ALUL] > core.RegValues[ALUR] {
				flags = 10
			} else if core.RegValues[ALUL] < core.RegValues[ALUR] {
				flags = 6
			} else if core.RegValues[ALUL] == core.RegValues[ALUR] {
				flags = 1
			}
			core.RegValues[FLAG] = flags

			core.accHandler("show")
			core.accHandler("flag")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.OperationClass = "fetch"
			core.clockTick(count)
			fmt.Print(cursor.Hide())
			return 1
		}

		core.accHandler("show")
		core.accHandler("flag")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 10:
		core.accHandler("close")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[DABA] = core.RegValues[ACC]
		core.RegValues[DABB] = core.RegValues[ACC]

		core.busHandler(int(DABA))
		core.busHandler(int(DABB))

		core.genRegHandler(int(reg1), "closein")
		core.RegValues[reg1] = core.RegValues[DABA]
		core.genRegHandler(int(reg1), "show")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 11:
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
