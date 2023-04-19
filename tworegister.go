package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) twoRegister(count int) int {

	fmt.Print(cursor.Hide())

	var reg1Mask uint16 = 0x0700
	var reg2Mask uint16 = 0x00E0

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8
	reg2 := core.RegValues[IR] & reg2Mask
	reg2 = reg2 >> 5

	switch count {
	case 7:
		if core.Opcode == "CMP" {
			core.genRegHandler(int(reg1), "closetop")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.RegValues[B1] = core.RegValues[reg1]

			core.genRegHandler(int(reg2), "closebot")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.RegValues[B2] = core.RegValues[reg2]
			core.busHandler(int(B1))
			time.Sleep(core.SleepTime * time.Millisecond)
			core.busHandler(int(B2))
			time.Sleep(core.SleepTime * time.Millisecond)
		} else if core.Opcode == "NOT" {
			core.genRegHandler(int(reg2), "closetop")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.RegValues[B1] = core.RegValues[reg1]
			core.busHandler(int(B1))
		}
	case 8:
		if core.Opcode == "CMP" {
			core.RegValues[ALUL] = core.RegValues[B1]
			core.RegValues[ALUR] = core.RegValues[B2]

			core.aluHandler("left")
			core.aluHandler("right")
			core.aluHandler("code")
			time.Sleep(core.SleepTime * time.Millisecond)
		} else if core.Opcode == "NOT" {
			core.RegValues[ALUL] = core.RegValues[B1]
			core.RegValues[ALUR] = 0xFFFF
			core.aluHandler("left")
			core.aluHandler("right")
			core.aluHandler("code")
			time.Sleep(core.SleepTime * time.Millisecond)

		}

	case 9:
		if core.Opcode == "CMP" {

			core.genRegHandler(int(reg1), "opentop")
			core.genRegHandler(int(reg2), "openbot")
			time.Sleep(core.SleepTime * time.Millisecond)

			var flags uint16 = 0

			if core.Opcode == "CMP" {
				if core.RegValues[ALUL] > core.RegValues[ALUR] {
					flags = 10
				} else if core.RegValues[ALUL] < core.RegValues[ALUR] {
					flags = 6
				} else if core.RegValues[ALUL] == core.RegValues[ALUR] {
					flags = 1
				}
				core.RegValues[FLAG] = flags
			}

			core.RegValues[ACC] = flags

			core.accHandler("show")
			core.accHandler("flag")

			time.Sleep(core.SleepTime * time.Millisecond)
			core.OperationClass = "fetch"
			core.clockTick(count)
			fmt.Print(cursor.Hide())
			return 1
		} else if core.Opcode == "NOT" {
			core.genRegHandler(int(reg2), "opentop")
			core.RegValues[ACC] = core.RegValues[ALUL] ^ core.RegValues[ALUR]
			core.accHandler("show")
			core.accHandler("flag")
			time.Sleep(core.SleepTime * time.Millisecond)

		}
	case 10:
		core.accHandler("close")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[DABA] = core.RegValues[ACC]
		core.RegValues[DABB] = core.RegValues[ACC]

		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		time.Sleep(core.SleepTime * time.Millisecond)
		core.genRegHandler(int(reg1), "closein")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 11:
		core.accHandler("open")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.genRegHandler(int(reg1), "openin")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[reg1] = core.RegValues[ACC]
		core.genRegHandler(int(reg1), "show")

		core.OperationClass = "fetch"
		core.clockTick(count)
		fmt.Print(cursor.Hide())
		return 1
	}

	core.clockTick(count)
	fmt.Print(cursor.Hide())
	return count + 1
}
