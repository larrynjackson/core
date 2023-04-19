package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) rtrn(count int) int {

	fmt.Print(cursor.Hide())

	switch count {
	case 7:
		core.RegValues[SR] = core.RegValues[SR] + 1
		core.genRegHandler(int(SR), "show")

		core.genRegHandler(int(SR), "closebot")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.RegValues[B2] = core.RegValues[SR]
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busHandler(int(B2))
	case 8:

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
	case 9:
		core.busGateHandler(int(B2DAB), "open")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.cntlRegHandler(int(DA), "openin")
		core.genRegHandler(int(SR), "openbot")

	case 10:

		core.cntlRegHandler(int(DA), "closeout")
		core.cntlRegHandler(int(DR), "closeout")
		core.cntlRegHandler(int(DR), "indirection")
		core.RegValues[DR] = core.CoreMemory[core.RegValues[DA]]

		core.memHandler(int(DA))
		core.memHandler(int(DR))

	case 11:
		core.cntlRegHandler(int(DA), "openout")
		core.cntlRegHandler(int(DR), "openout")
		core.cntlRegHandler(int(DR), "show")

	case 12:

		core.cntlRegHandler(int(DR), "closein")
		core.cntlRegHandler(int(DR), "indirection")
		core.RegValues[DABA] = core.RegValues[DR]
		core.RegValues[DABB] = core.RegValues[DR]
		time.Sleep(core.SleepTime * time.Millisecond)
		core.busHandler(int(DABA))
		core.busHandler(int(DABB))
		time.Sleep(core.SleepTime * time.Millisecond)

		core.cntlRegHandler(int(IA), "closein")
		core.cntlRegHandler(int(IA), "outdirection")
		core.RegValues[IA] = core.RegValues[DABB]
		core.CoreMemPoint = int16(core.RegValues[IA])
		core.cntlRegHandler(int(IA), "show")

	case 13:

		core.cntlRegHandler(int(IA), "openin")
		time.Sleep(core.SleepTime * time.Millisecond)
		core.cntlRegHandler(int(DR), "openin")

		core.OperationClass = "fetch"
		core.clockTick(count)
		fmt.Print(cursor.Hide())
		return 1
	}
	core.clockTick(count)
	fmt.Print(cursor.Hide())
	return count + 1
}
