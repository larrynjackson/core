package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) inOut(count int) int {

	fmt.Print(cursor.Hide())

	var reg1Mask uint16 = 0x0700

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8

	switch count {
	case 7:

		if core.Opcode == "IN" {

			fmt.Print(cursor.Hide())

			fmt.Print(cursor.MoveTo(41, 108))
			fmt.Print("$  ")
			fmt.Print(cursor.MoveTo(41, 110))
			fmt.Print(cursor.Show())
			var b []byte = make([]byte, 1)
			os.Stdin.Read(b)

			character := b[0]
			if character >= 'a' && character <= 'z' {
				character = character - 32 // space character
			}

			fmt.Print(cursor.MoveTo(41, 110))
			fmt.Printf("%c", character)

			core.RegValues[IN] = uint16(character)
			core.ioRegHandler(int(IN), "show")

		} else if core.Opcode == "OUT" {
			core.genRegHandler(int(reg1), "closetop")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.busHandler(int(B1))
			core.RegValues[B1] = core.RegValues[reg1]

			core.busGateHandler(int(B1DAB), "close")
			time.Sleep(core.SleepTime * time.Millisecond)

			core.RegValues[DABA] = core.RegValues[B1]
			core.RegValues[DABB] = core.RegValues[B1]
			core.busHandler(int(DABA))
			time.Sleep(core.SleepTime * time.Millisecond)
			core.busHandler(int(DABA))
			time.Sleep(core.SleepTime * time.Millisecond)
			core.ioRegHandler(int(OUT), "close")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.RegValues[OUT] = core.RegValues[DABB]
			core.ioRegHandler(int(OUT), "show")

		}

	case 8:
		if core.Opcode == "IN" {
			core.ioRegHandler(int(IN), "close")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.RegValues[DABA] = core.RegValues[IN]
			core.RegValues[DABB] = core.RegValues[IN]

			core.busHandler(int(DABA))
			time.Sleep(core.SleepTime * time.Millisecond)
			core.busHandler(int(DABB))
			time.Sleep(core.SleepTime * time.Millisecond)

			core.RegValues[reg1] = core.RegValues[DABA]

			core.genRegHandler(int(reg1), "closein")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.genRegHandler(int(reg1), "show")
			time.Sleep(core.SleepTime * time.Millisecond)
		} else if core.Opcode == "OUT" {

			core.ioRegHandler(int(OUT), "open")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.genRegHandler(int(reg1), "opentop")

			fmt.Print(cursor.Hide())

			fmt.Print(cursor.MoveTo(core.OutRow, core.OutCol))
			fmt.Printf("%c", core.RegValues[OUT])
			fmt.Print(cursor.Show())

			core.OperationClass = "fetch"
			core.clockTick(count)
			fmt.Print(cursor.Hide())
			return 1
		}

	case 9:
		if core.Opcode == "IN" {
			core.ioRegHandler(int(IN), "open")
			time.Sleep(core.SleepTime * time.Millisecond)
			core.genRegHandler(int(reg1), "openin")

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
