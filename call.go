package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) call(count int) int {

	fmt.Print(cursor.Hide())

	var reg1Mask uint16 = 0x0700

	reg1 := core.RegValues[IR] & reg1Mask
	reg1 = reg1 >> 8

	switch count {
	case 7:

		core.callPushAll()

		core.cntlRegHandler(int(IA), "closein")
		core.cntlRegHandler(int(IA), "indirection")
		core.RegValues[DABA] = core.RegValues[IA]
		core.RegValues[DABB] = core.RegValues[IA]
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		core.cntlRegHandler(int(DR), "closein")
		core.cntlRegHandler(int(DR), "outdirection")
		core.RegValues[DR] = core.RegValues[IA]
		core.cntlRegHandler(int(DR), "show")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 8:
		core.cntlRegHandler(int(IA), "openin")
		core.cntlRegHandler(int(DR), "openin")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 9:
		core.genRegHandler(int(SR), "closebot")
		core.RegValues[B2] = core.RegValues[SR]
		core.busHandler(int(B2))

		core.busGateHandler(int(B2DAB), "close")
		core.RegValues[DABA] = core.RegValues[B2]
		core.RegValues[DABB] = core.RegValues[B2]
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))

		core.cntlRegHandler(int(DA), "closein")
		core.RegValues[DA] = core.RegValues[B2]
		core.cntlRegHandler(int(DA), "show")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 10:
		core.busGateHandler(int(B2DAB), "open")
		core.cntlRegHandler(int(DA), "openin")
		core.genRegHandler(int(SR), "openbot")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 11:
		core.cntlRegHandler(int(DA), "closeout")
		core.cntlRegHandler(int(DR), "closeout")
		core.cntlRegHandler(int(DR), "outdirection")

		core.memHandler(int(DA))
		core.memHandler(int(DR))
		core.CoreMemory[core.RegValues[DA]] = core.RegValues[DR]
		core.RegValues[SR] = core.RegValues[SR] - 1
		core.genRegHandler(int(SR), "show")
		time.Sleep(core.SleepTime * time.Millisecond)
	case 12:
		core.cntlRegHandler(int(DA), "openout")
		core.cntlRegHandler(int(DR), "openout")

		core.genRegHandler(int(reg1), "closetop")
		core.RegValues[B1] = core.RegValues[reg1]
		core.busHandler(int(B1))

		core.busGateHandler(int(B1DAB), "close")
		core.RegValues[DABA] = core.RegValues[B2]
		core.RegValues[DABB] = core.RegValues[B2]
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))

		core.cntlRegHandler(int(IA), "closein")
		core.cntlRegHandler(int(IA), "outdirection")
		core.RegValues[IA] = core.RegValues[B1]
		core.CoreMemPoint = int16(core.RegValues[IA])
		core.cntlRegHandler(int(IA), "show")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 13:
		core.busGateHandler(int(B1DAB), "open")
		core.cntlRegHandler(int(IA), "openin")
		core.genRegHandler(int(reg1), "opentop")
		time.Sleep(core.SleepTime * time.Millisecond)

		if core.DebugMode {
			core.printStack()
		}

		core.OperationClass = "fetch"
		core.clockTick(count)
		fmt.Print(cursor.Hide())
		return 1
	}

	core.clockTick(count)
	fmt.Print(cursor.Hide())
	return count + 1

}

func (core *Config) callPushAll() {

	var reg1 int

	for reg1 = 7; reg1 >= 0; reg1-- {

		core.genRegHandler(int(reg1), "closetop")
		core.RegValues[B1] = core.RegValues[reg1]
		core.busHandler(int(B1))

		core.genRegHandler(int(SR), "closebot")
		core.RegValues[B2] = core.RegValues[SR]
		core.busHandler(int(B2))
		time.Sleep(core.SleepTime * time.Millisecond)

		core.busGateHandler(int(B2DAB), "close")
		core.RegValues[DABA] = core.RegValues[B2]
		core.RegValues[DABB] = core.RegValues[B2]
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))

		core.cntlRegHandler(int(DA), "closein")
		core.RegValues[DA] = core.RegValues[B2]
		core.cntlRegHandler(int(DA), "show")
		time.Sleep(core.SleepTime * time.Millisecond)

		core.busGateHandler(int(B2DAB), "open")
		core.cntlRegHandler(int(DA), "openin")
		core.genRegHandler(int(SR), "openbot")
		time.Sleep(core.SleepTime * time.Millisecond)

		core.busGateHandler(int(B1DAB), "close")
		core.RegValues[DABA] = core.RegValues[B2]
		core.RegValues[DABB] = core.RegValues[B2]
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))

		core.cntlRegHandler(int(DR), "closein")
		core.cntlRegHandler(int(DR), "outdirection")
		core.RegValues[DR] = core.RegValues[B1]
		core.cntlRegHandler(int(DR), "show")
		time.Sleep(core.SleepTime * time.Millisecond)

		core.busGateHandler(int(B1DAB), "open")
		core.cntlRegHandler(int(DR), "openin")
		core.genRegHandler(int(reg1), "opentop")
		time.Sleep(core.SleepTime * time.Millisecond)

		core.cntlRegHandler(int(DA), "closeout")
		core.cntlRegHandler(int(DR), "closeout")
		core.cntlRegHandler(int(DR), "outdirection")

		core.memHandler(int(DA))
		core.memHandler(int(DR))
		core.CoreMemory[core.RegValues[DA]] = core.RegValues[DR]

		core.RegValues[SR] = core.RegValues[SR] - 1
		core.genRegHandler(int(SR), "show")
		time.Sleep(core.SleepTime * time.Millisecond)

		core.cntlRegHandler(int(DA), "openout")
		core.cntlRegHandler(int(DR), "openout")
		time.Sleep(core.SleepTime * time.Millisecond)

		if core.DebugMode {
			core.printStack()
		}

	}
}
