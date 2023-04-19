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
		core.cntlRegHandler(int(IA), "closein")
		core.cntlRegHandler(int(IA), "indirection")
		core.RegValues[DABA] = core.RegValues[IA]
		core.RegValues[DABB] = core.RegValues[IA]
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		time.Sleep(core.SleepTime * time.Millisecond)
		core.cntlRegHandler(int(DR), "closein")
		core.cntlRegHandler(int(DR), "outdirection")
		core.RegValues[DR] = core.RegValues[IA]
		core.cntlRegHandler(int(DR), "show")
	case 8:
		core.cntlRegHandler(int(IA), "openin")
		core.cntlRegHandler(int(DR), "openin")
	case 9:
		core.genRegHandler(int(SR), "closebot")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[B2] = core.RegValues[SR]
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busHandler(int(B2))

		core.busGateHandler(int(B2DAB), "close")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[DABA] = core.RegValues[B2]
		core.RegValues[DABB] = core.RegValues[B2]
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		time.Sleep(core.SleepTime * time.Millisecond)

		core.cntlRegHandler(int(DA), "closein")
		core.RegValues[DA] = core.RegValues[B2]
		core.cntlRegHandler(int(DA), "show")

	case 10:
		core.busGateHandler(int(B2DAB), "open")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.cntlRegHandler(int(DA), "openin")
		core.genRegHandler(int(SR), "openbot")
	case 11:
		core.cntlRegHandler(int(DA), "closeout")
		core.cntlRegHandler(int(DR), "closeout")
		core.cntlRegHandler(int(DR), "outdirection")
		time.Sleep(core.SleepTime * time.Millisecond)

		core.memHandler(int(DA))
		core.memHandler(int(DR))
		core.CoreMemory[core.RegValues[DA]] = core.RegValues[DR]
		core.RegValues[SR] = core.RegValues[SR] - 1
		core.genRegHandler(int(SR), "show")
	case 12:
		core.cntlRegHandler(int(DA), "openout")
		core.cntlRegHandler(int(DR), "openout")

		core.genRegHandler(int(reg1), "closetop")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[B1] = core.RegValues[reg1]
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busHandler(int(B1))

		core.busGateHandler(int(B1DAB), "close")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[DABA] = core.RegValues[B2]
		core.RegValues[DABB] = core.RegValues[B2]
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		time.Sleep(core.SleepTime * time.Millisecond)

		core.cntlRegHandler(int(IA), "closein")
		core.cntlRegHandler(int(IA), "outdirection")
		core.RegValues[IA] = core.RegValues[B1]
		core.CoreMemPoint = int16(core.RegValues[IA])
		core.cntlRegHandler(int(IA), "show")

	case 13:
		core.busGateHandler(int(B1DAB), "open")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.cntlRegHandler(int(IA), "openin")
		core.genRegHandler(int(reg1), "opentop")

		core.OperationClass = "fetch"
		core.clockTick(count)
		fmt.Print(cursor.Hide())
		return 1
	}

	core.clockTick(count)
	fmt.Print(cursor.Hide())
	return count + 1

}
