package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) translateInstruction(count int) int {

	fmt.Print(cursor.Hide())
	OpCodeArray := []string{"HALT", "ADD", "SUB", "AND", "XOR", "OR", "LDW", "STW", "SHL", "SHR", "NOT", "CMP", "BFLAG", "ADDI", "SUBI", "ANDI", "XORI", "ORI", "LDI", "CMPI", "JUMP", "PUSH", "POP", "OUT", "IN", "LDSR", "NOOP", "CALL", "RTRN", "MVSR"}
	regArray := []string{"R0", "R1", "R2", "R3", "R4", "R5", "R6", "R7"}

	var opCodeMask uint16 = 0xF800
	var reg1Mask uint16 = 0x0700
	var reg2Mask uint16 = 0x00E0
	var reg3Mask uint16 = 0x001C

	opCode := opCodeMask & core.RegValues[IR]
	opCode = opCode >> 11

	core.Opcode = OpCodeArray[opCode]

	var instString string

	if opCode == 0 {
		core.OperationClass = "Halt"
	} else if opCode < 8 {
		core.OperationClass = "3Register"
		reg1 := core.RegValues[IR] & reg1Mask
		reg1 = reg1 >> 8
		reg2 := core.RegValues[IR] & reg2Mask
		reg2 = reg2 >> 5
		reg3 := core.RegValues[IR] & reg3Mask
		reg3 = reg3 >> 2
		instString += "  " + regArray[reg1] + "," + regArray[reg2] + "," + regArray[reg3]
	} else if opCode < 10 {
		reg1 := core.RegValues[IR] & reg1Mask
		reg1 = reg1 >> 8
		reg2 := core.RegValues[IR] & reg2Mask
		reg2 = reg2 >> 5
		value := core.RegValues[IR] & 0x001E
		value = value >> 1
		core.OperationClass = "shift"
		instString += "  " + regArray[reg1] + "," + regArray[reg2] + "," + fmt.Sprint(value)
	} else if opCode < 12 {
		reg1 := core.RegValues[IR] & reg1Mask
		reg1 = reg1 >> 8
		reg2 := core.RegValues[IR] & reg2Mask
		reg2 = reg2 >> 5
		core.OperationClass = "2Register"
		instString += "  " + regArray[reg1] + "," + regArray[reg2]
	} else if opCode < 13 {
		reg1 := core.RegValues[IR] & reg1Mask
		reg1 = reg1 >> 8
		flag := core.RegValues[IR] & 0x00F0
		flag = flag >> 4
		core.OperationClass = "branch"
		instString += "  " + regArray[reg1] + ", "
		if flag == 1 {
			instString += "EQ"
		} else if flag == 2 {
			instString += "NE"
		} else if flag == 4 {
			instString += "LT"
		} else if flag == 8 {
			instString += "GT"
		}
	} else if opCode < 20 {
		reg1 := core.RegValues[IR] & reg1Mask
		reg1 = reg1 >> 8
		value := core.RegValues[IR] & 0x00FF
		core.OperationClass = "immediate"
		instString += "  " + regArray[reg1] + "," + fmt.Sprint(value)
	} else if opCode < 26 {
		reg1 := core.RegValues[IR] & reg1Mask
		reg1 = reg1 >> 8
		instString += "  " + regArray[reg1]
		switch core.Opcode {
		case "JUMP":
			core.OperationClass = "jump"
		case "PUSH":
			core.OperationClass = "push"
		case "POP":
			core.OperationClass = "pop"
		case "OUT":
			core.OperationClass = "inout"
		case "IN":
			core.OperationClass = "inout"
		case "LDSR":
			core.OperationClass = "ldsr"
		}
	} else if opCode < 27 {
		core.OperationClass = "NOOP"
	} else if opCode < 28 {
		core.OperationClass = "call"
		reg1 := core.RegValues[IR] & reg1Mask
		reg1 = reg1 >> 8
		instString += "  " + regArray[reg1]
	} else if opCode == 28 {
		core.OperationClass = "rtrn"
	} else if opCode == 29 {
		reg1 := core.RegValues[IR] & reg1Mask
		reg1 = reg1 >> 8
		instString += "  " + regArray[reg1]
		core.OperationClass = "mvsr"
	}

	core.InstructionString = core.Opcode + instString
	core.RegValues[INST] = core.RegValues[IR] // not used, but saved anyway

	core.instructionHandler("show")
	core.clockTick(count)
	fmt.Print(cursor.Show())
	return count + 1
}
