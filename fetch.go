package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) fetchInstruction(count int) int {

	fmt.Print(cursor.Hide())
	switch count {
	case 1:

		core.cntlRegHandler(int(IA), "closeout")
		core.cntlRegHandler(int(IA), "outdirection")
		core.RegValues[IR] = core.CoreMemory[core.CoreMemPoint]
		core.memHandler(int(IA))
		time.Sleep(core.SleepTime * time.Millisecond)

	case 2:

		core.memHandler(int(IA))
		core.memHandler(int(IR))
		time.Sleep(core.SleepTime * time.Millisecond)

	case 3:
		core.cntlRegHandler(int(IR), "closein")
		core.cntlRegHandler(int(IR), "show")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 4:

		core.cntlRegHandler(int(IA), "openout")
		core.cntlRegHandler(int(IR), "openin")
		time.Sleep(core.SleepTime * time.Millisecond)

	case 5:

		core.CoreMemPoint = core.CoreMemPoint + 1
		core.RegValues[IA] = uint16(core.CoreMemPoint)
		core.cntlRegHandler(int(IA), "show")

		time.Sleep(core.SleepTime * time.Millisecond)

		core.OperationClass = "translate"

	}

	core.clockTick(count)
	fmt.Print(cursor.Show())
	return count + 1
}
