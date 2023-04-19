package main

import (
	"time"

	"lnj.com/core/assembler"
)

type Config struct {
	ASM               assembler.Assembler
	DABA              Bus
	DABB              Bus
	B1                Bus
	B2                Bus
	MEM               Memory
	ACC               ACCReg
	ALU               ALUReg
	Instruction       Instruction
	B1DABGate         Gate
	B2DABGate         Gate
	OUTPUTR           IOReg
	INPUTR            IOReg
	IA                ControlReg
	DA                ControlReg
	DR                ControlReg
	IR                ControlReg
	CoreMemory        [65536]uint16
	CoreMemPoint      int16
	OperationClass    string
	Opcode            string
	InstructionString string
	RegValues         [24]uint16
	GenRegs           [9]Register
	SleepTime         time.Duration
	OutRow            int
	OutCol            int
}

type Regval int

const (
	R0 Regval = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	SR
	DA
	DR
	IA
	IR
	IN
	OUT
	B1
	B2
	DABA
	DABB
	ALUL
	ALUR
	ACC
	INST
	FLAG
	B1DAB // no storage, just used as a key
	B2DAB // no storage, just used as a key
)

func main() {

	var core Config

	core.OperationClass = "fetch"
	core.CoreMemPoint = 0

	core.CreateCoreComponents()

	core.drawScreen()

	core.resetCore()

	core.shell()
}
