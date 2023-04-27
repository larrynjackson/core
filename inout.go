package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) scrollInOutArea() {

	keyOrder := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for key := range keyOrder {
		core.IOmem[key] = ""
		core.IOmem[key] = core.IOmem[key+1]
	}

	for k, v := range core.IOmem {
		if k < 15 {
			fmt.Print(cursor.MoveTo(ioMemTopRow+k, ioMemCol))
			fmt.Printf("%85s", SP)
			fmt.Print(cursor.MoveTo(ioMemTopRow+k, ioMemCol))
			fmt.Print(v)
		}

	}

}

func (core *Config) inOut(count int) int {

	fmt.Print(cursor.Hide())

	var reg1Mask uint16 = 0x0700

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8

	switch count {
	case 7:

		if core.Opcode == "IN" {

			fmt.Print(cursor.Hide())
			fmt.Print(cursor.MoveTo(ioTypingLine, ioTypingCol))
			fmt.Print("$ ")

			fmt.Print(cursor.MoveTo(ioTypingLine, core.IOrwCol))
			fmt.Print(cursor.Show())
			var b []byte = make([]byte, 1)
			var character byte

			for {

				os.Stdin.Read(b)
				character = b[0]
				if character > 8 {
					break
				} else if character == 8 && core.IOrwCol > ioMemCol {
					break
				}
			}

			if character >= 'a' && character <= 'z' {
				character = character - 32 // space character
			}

			if character == 8 && core.IOrwCol > ioMemCol {
				core.IOrwCol--
				character = 32
				fmt.Print(cursor.MoveTo(ioTypingLine, core.IOrwCol))
				fmt.Printf("%c", character)

				tmpString := core.IOmem[15]
				var newString string = ""
				idx := len(tmpString) - 1
				newString = tmpString[0:idx]

				core.IOmem[15] = newString

				fmt.Print(cursor.MoveTo(ioTypingLine, core.IOrwCol))
				core.RegValues[reg1] = core.RegValues[reg1] - 1
				core.genRegHandler(int(reg1), "show")
				return 7

			} else if character == 13 {
				character = 0
			} else {

				core.IOmem[15] = core.IOmem[15] + string(character)

				fmt.Print(cursor.MoveTo(ioTypingLine, core.IOrwCol))
				fmt.Printf("%c", character)

				core.IOrwCol++

			}
			core.RegValues[IN] = uint16(character)
			core.ioRegHandler(int(IN), "show")
			time.Sleep(core.SleepTime * time.Millisecond)

		} else if core.Opcode == "OUT" {

			core.genRegHandler(int(reg1), "closetop")
			core.RegValues[B1] = core.RegValues[reg1]
			core.busHandler(int(B1))

			core.busGateHandler(int(B1DAB), "close")

			core.RegValues[DABA] = core.RegValues[B1]
			core.RegValues[DABB] = core.RegValues[B1]
			core.busHandler(int(DABA))
			core.busHandler(int(DABB))

			core.cntlRegHandler(int(DA), "closein")
			core.RegValues[DA] = core.RegValues[DABB]
			core.cntlRegHandler(int(DA), "show")
			time.Sleep(core.SleepTime * time.Millisecond)

		}

	case 8:
		if core.Opcode == "IN" {

			core.ioRegHandler(int(IN), "close")
			core.RegValues[DABA] = core.RegValues[IN]
			core.RegValues[DABB] = core.RegValues[IN]

			core.busHandler(int(DABA))
			core.busHandler(int(DABB))

			core.cntlRegHandler(int(DR), "closein")
			core.RegValues[DR] = core.RegValues[IN]
			core.cntlRegHandler(int(DR), "outdirection")
			core.cntlRegHandler(int(DR), "show")
			time.Sleep(core.SleepTime * time.Millisecond)

		} else if core.Opcode == "OUT" {

			core.cntlRegHandler(int(DA), "openin")
			core.genRegHandler(int(reg1), "opentop")
			core.RegValues[int(reg1)] = core.RegValues[int(reg1)] + 1
			core.genRegHandler(int(reg1), "show")
			time.Sleep(core.SleepTime * time.Millisecond)
		}
	case 9:
		if core.Opcode == "IN" {
			core.cntlRegHandler(int(DR), "openin")
			core.ioRegHandler(int(IN), "open")
			time.Sleep(core.SleepTime * time.Millisecond)

		} else if core.Opcode == "OUT" {

			core.cntlRegHandler(int(DA), "closeout")

			core.cntlRegHandler(int(DR), "closeout")
			core.RegValues[DR] = core.CoreMemory[core.RegValues[DA]]
			core.cntlRegHandler(int(DR), "indirection")
			core.cntlRegHandler(int(DR), "show")
			core.memHandler(int(DA))
			core.memHandler(int(DR))
			time.Sleep(core.SleepTime * time.Millisecond)

		}
	case 10:
		if core.Opcode == "IN" {
			core.genRegHandler(int(reg1), "closetop")
			core.RegValues[B1] = core.RegValues[reg1]
			core.RegValues[DABA] = core.RegValues[reg1]
			core.RegValues[DABB] = core.RegValues[reg1]
			core.busGateHandler(int(B1DAB), "close")

			core.busHandler(int(B1))
			core.busHandler(int(DABA))
			core.busHandler(int(DABB))
			core.RegValues[DA] = core.RegValues[reg1]
			core.cntlRegHandler(int(DA), "closein")
			core.cntlRegHandler(int(DA), "show")
			time.Sleep(core.SleepTime * time.Millisecond)

		} else if core.Opcode == "OUT" {
			core.cntlRegHandler(int(DA), "openout")
			core.cntlRegHandler(int(DR), "openout")
			time.Sleep(core.SleepTime * time.Millisecond)
		}
	case 11:
		if core.Opcode == "IN" {
			core.genRegHandler(int(reg1), "opentop")
			core.busGateHandler(int(B1DAB), "open")
			core.cntlRegHandler(int(DA), "openin")
			time.Sleep(core.SleepTime * time.Millisecond)
		} else if core.Opcode == "OUT" {
			core.cntlRegHandler(int(DR), "closein")
			core.cntlRegHandler(int(DR), "indirection")
			core.RegValues[DABA] = core.RegValues[DR]
			core.RegValues[DABB] = core.RegValues[DR]

			core.ioRegHandler(int(OUT), "close")
			core.RegValues[OUT] = core.RegValues[DR]
			core.ioRegHandler(int(OUT), "show")
			time.Sleep(core.SleepTime * time.Millisecond)
		}

	case 12:
		if core.Opcode == "IN" {
			core.cntlRegHandler(int(DA), "closeout")

			core.cntlRegHandler(int(DR), "closeout")
			core.cntlRegHandler(int(DR), "outdirection")
			core.memHandler(int(DA))
			core.memHandler(int(DR))
			time.Sleep(core.SleepTime * time.Millisecond)

		} else if core.Opcode == "OUT" {
			core.cntlRegHandler(int(DR), "openin")
			core.ioRegHandler(int(OUT), "open")

			if core.RegValues[OUT] == 0 {
				core.scrollInOutArea()
				core.IOmem[15] = ""

				fmt.Print(cursor.MoveTo(ioTypingLine, ioMemCol))
				fmt.Printf("%85s", SP)

				core.IOrwCol = ioMemCol
				core.OperationClass = "fetch"
				return 1
			}
		}
	case 13:
		if core.Opcode == "IN" {
			core.RegValues[reg1] = core.RegValues[reg1] + 1
			core.genRegHandler(int(reg1), "show")
			core.cntlRegHandler(int(DA), "openout")
			core.cntlRegHandler(int(DR), "openout")
			core.CoreMemory[core.RegValues[DA]] = core.RegValues[DR]
			core.clockTick(count)

			if core.RegValues[IN] == 0 {
				core.scrollInOutArea()
				core.IOmem[15] = ""

				fmt.Print(cursor.MoveTo(ioTypingLine, ioMemCol))
				fmt.Printf("%85s", SP)

				core.IOrwCol = ioMemCol
				core.OperationClass = "fetch"
				return 1
			}
			return 7
		} else if core.Opcode == "OUT" {

			fmt.Print(cursor.MoveTo(ioTypingLine, ioTypingCol))
			fmt.Printf("%85s", SP)
			fmt.Print(cursor.MoveTo(ioTypingLine, ioTypingCol))
			fmt.Print("$ ")
			core.IOmem[15] = core.IOmem[15] + string(byte(core.RegValues[OUT]))

			fmt.Print(cursor.MoveTo(ioTypingLine, core.IOrwCol))
			fmt.Printf("%c", core.RegValues[OUT])
			core.IOrwCol++
			return 7
		}
	}
	core.clockTick(count)
	fmt.Print(cursor.Hide())
	return count + 1

}
