package main

import (
	"fmt"
	"time"

	"github.com/ahmetalpbalkan/go-cursor"
)

type RegVal int

const (
	SP                   = " "
	TLB                  = "\u250C"
	TRB                  = "\u2510"
	VRB                  = "\u251C"
	VLB                  = "\u2524"
	VB                   = "\u2502"
	HB                   = "\u2500"
	BLB                  = "\u2514"
	BRB                  = "\u2518"
	TLP                  = "\u2554"
	TRP                  = "\u2557"
	VRP                  = "\u2560"
	VLP                  = "\u2563"
	VP                   = "\u2551"
	HP                   = "\u2550"
	BLP                  = "\u255A"
	BRP                  = "\u255D"
	GRP                  = "\u255E"
	GLP                  = "\u2561"
	GTP                  = "\u2568"
	GBP                  = "\u2565"
	BCU                  = "\u2569"
	RARROW               = "\u25B7"
	LARROW               = "\u25C1"
	openGate             = HP + GLP + GRP
	closeGate            = HP + HP + HP
	regInGateCol         = 20
	regOutGateCol        = 39
	regValueCol          = 29
	controlRegValueCol   = 93
	controlRegInGateCol  = 84
	controlRegOutGateCol = 103
	irInGateCol          = 103
	irOutGateCol         = 62
	memoryCol            = 112
	INDIRECTION          = LARROW + HB + HB
	OUTDIRECTION         = HB + HB + RARROW
)

type Register struct {
	hexRow        int
	decRow        int
	valueCol      int // R0-R7, SR
	ingateRow     int
	ingateCol     int
	outgateTopRow int
	outgateBotRow int
	outgateCol    int
}

type ControlReg struct {
	hexRow     int
	decRow     int
	valueCol   int // IA, DA, DR, IR
	directRow  int
	directCol  int
	gateRow    int
	ingateCol  int
	outgateCol int
}

type IOReg struct {
	decRow   int
	valueCol int
	gateRow  int
	gateCol  int
}

type ALUReg struct {
	decRow      int
	leftDecCol  int
	rightDecCol int
	opCodeRow   int
	opCodeCol   int
}

type ACCReg struct {
	decRow    int
	decCol    int
	flagsRow  int
	flagsCol  int
	boxBusRow int
	busRow    int
	busCol    int
}

type Memory struct {
	iaHexRow  int
	daHexRow  int
	daDecRow  int
	drHexRow  int
	drDecRow  int
	irHexRow  int
	memoryCol int
}

type Instruction struct {
	row int
	col int
}

type Bus struct {
	hexRow int
	hexCol int
	decRow int
	decCol int
}

type Gate struct {
	row int
	col int
}

func (core *Config) CreateCoreComponents() {

	core.DABA = Bus{
		hexRow: 1,
		hexCol: 18,
		decRow: 2,
		decCol: 18,
	}

	core.DABB = Bus{
		hexRow: 1,
		hexCol: 82,
		decRow: 2,
		decCol: 82,
	}

	core.B1 = Bus{
		hexRow: 1,
		hexCol: 43,
		decRow: 2,
		decCol: 43,
	}

	core.B2 = Bus{
		hexRow: 1,
		hexCol: 60,
		decRow: 2,
		decCol: 60,
	}

	core.MEM = Memory{
		iaHexRow:  10,
		daHexRow:  14,
		daDecRow:  15,
		drHexRow:  18,
		drDecRow:  19,
		irHexRow:  26,
		memoryCol: memoryCol,
	}

	core.ACC = ACCReg{
		decRow:    46,
		decCol:    50,
		flagsRow:  46,
		flagsCol:  66,
		boxBusRow: 47,
		busRow:    48,
		busCol:    52,
	}

	core.ALU = ALUReg{
		decRow:      43,
		leftDecCol:  44,
		rightDecCol: 57,
		opCodeRow:   42,
		opCodeCol:   66,
	}

	core.Instruction = Instruction{
		row: 31,
		col: 89,
	}

	core.B2DABGate = Gate{
		row: 6,
		col: 62,
	}

	core.B1DABGate = Gate{
		row: 5,
		col: 45,
	}

	core.OUTPUTR = IOReg{
		decRow:   37,
		valueCol: controlRegValueCol,
		gateRow:  37,
		gateCol:  controlRegInGateCol,
	}

	core.INPUTR = IOReg{
		decRow:   41,
		valueCol: controlRegValueCol,
		gateRow:  41,
		gateCol:  controlRegInGateCol,
	}

	core.IA = ControlReg{
		hexRow:     10,
		decRow:     11,
		valueCol:   controlRegValueCol,
		directRow:  12,
		directCol:  controlRegValueCol + 1,
		gateRow:    11,
		ingateCol:  controlRegInGateCol,
		outgateCol: controlRegOutGateCol,
	}

	core.DA = ControlReg{
		hexRow:     14,
		decRow:     15,
		valueCol:   controlRegValueCol,
		directRow:  16,
		directCol:  controlRegValueCol + 1,
		gateRow:    15,
		ingateCol:  controlRegInGateCol,
		outgateCol: controlRegOutGateCol,
	}

	core.DR = ControlReg{
		hexRow:     18,
		decRow:     19,
		valueCol:   controlRegValueCol,
		directRow:  20,
		directCol:  controlRegValueCol + 1,
		gateRow:    19,
		ingateCol:  controlRegInGateCol,
		outgateCol: controlRegOutGateCol,
	}

	core.IR = ControlReg{
		hexRow:     27,
		valueCol:   controlRegValueCol,
		directRow:  28,
		directCol:  controlRegValueCol + 1,
		gateRow:    27,
		ingateCol:  irInGateCol,
		outgateCol: irOutGateCol,
	}

	core.GenRegs[R0] = Register{
		hexRow:        6,
		decRow:        7,
		valueCol:      regValueCol,
		ingateRow:     7,
		ingateCol:     regInGateCol,
		outgateTopRow: 6,
		outgateBotRow: 8,
		outgateCol:    regOutGateCol,
	}

	core.GenRegs[R1] = Register{
		hexRow:        10,
		decRow:        11,
		valueCol:      regValueCol,
		ingateRow:     11,
		ingateCol:     regInGateCol,
		outgateTopRow: 10,
		outgateBotRow: 12,
		outgateCol:    regOutGateCol,
	}

	core.GenRegs[R2] = Register{
		hexRow:        14,
		decRow:        15,
		valueCol:      regValueCol,
		ingateRow:     15,
		ingateCol:     regInGateCol,
		outgateTopRow: 14,
		outgateBotRow: 16,
		outgateCol:    regOutGateCol,
	}

	core.GenRegs[R3] = Register{
		hexRow:        18,
		decRow:        19,
		valueCol:      regValueCol,
		ingateRow:     19,
		ingateCol:     regInGateCol,
		outgateTopRow: 18,
		outgateBotRow: 20,
		outgateCol:    regOutGateCol,
	}

	core.GenRegs[R4] = Register{
		hexRow:        22,
		decRow:        23,
		valueCol:      regValueCol,
		ingateRow:     23,
		ingateCol:     regInGateCol,
		outgateTopRow: 22,
		outgateBotRow: 24,
		outgateCol:    regOutGateCol,
	}

	core.GenRegs[R5] = Register{
		hexRow:        26,
		decRow:        27,
		valueCol:      regValueCol,
		ingateRow:     27,
		ingateCol:     regInGateCol,
		outgateTopRow: 26,
		outgateBotRow: 28,
		outgateCol:    regOutGateCol,
	}

	core.GenRegs[R6] = Register{
		hexRow:        30,
		decRow:        31,
		valueCol:      regValueCol,
		ingateRow:     31,
		ingateCol:     regInGateCol,
		outgateTopRow: 30,
		outgateBotRow: 32,
		outgateCol:    regOutGateCol,
	}

	core.GenRegs[R7] = Register{
		hexRow:        34,
		decRow:        35,
		valueCol:      regValueCol,
		ingateRow:     35,
		ingateCol:     regInGateCol,
		outgateTopRow: 34,
		outgateBotRow: 36,
		outgateCol:    regOutGateCol,
	}

	core.GenRegs[SR] = Register{
		hexRow:        38,
		decRow:        39,
		valueCol:      regValueCol,
		ingateRow:     39,
		ingateCol:     regInGateCol,
		outgateTopRow: 38,
		outgateBotRow: 40,
		outgateCol:    regOutGateCol,
	}

}

func (core *Config) clockTick(tickCount int) {

	fmt.Print(cursor.Hide())

	fmt.Print(cursor.MoveTo(5, 135))
	fmt.Printf("%10s", SP)
	fmt.Print(cursor.MoveTo(4, 135))
	fmt.Printf("%10s", SP)
	fmt.Print(cursor.MoveTo(3, 135))
	fmt.Printf("%10s", SP)

	fmt.Print(cursor.MoveTo(5, 135))
	fmt.Printf("%s", HB)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(5, 136))
	fmt.Printf("%s", BRB)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(4, 136))
	fmt.Printf("%s", VB)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(3, 136))
	fmt.Printf("%s", TLB)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(3, 137))
	fmt.Printf("%s", HB)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(3, 138))
	fmt.Printf("%s", TRB)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(4, 138))
	fmt.Printf("%s", VB)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(5, 138))
	fmt.Printf("%s", VB)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(2, 136))
	fmt.Printf("%2d", tickCount)

	fmt.Print(cursor.Show())
	core.cursorHome()
	//core.setOutputHome()
}

func (core *Config) drawScreen() {

	hex := "0x19AF"
	dec := "00157"

	FLAGS := "0  0  0  0"
	bline := HB + HB + HB + HB + HB + HB + HB + HB + HB + HB
	regTop := TLB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + TRB
	regBot := BLB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + BRB
	regRgate := GRP + HP + GLP + GRP + HP + HP
	regLgate := VRP + HP + GLP + GRP + HP + GLP
	regRpipeSet := HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP
	regLine := VRB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + VLB
	reg2reggate := GRP + HP + GLP + GRP + HP + GLP

	regVpipe := HP + HP + HP + HP + GLP
	regVgateSet := HP + GLP + GRP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP

	regIRtop := TLB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + TRB
	regIRbot := BLB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + BRB

	regALUtop := TLB + HB + HB + GTP + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + GTP + HB + HB + TRB
	regALUbot := BLB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + GBP + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + BRB

	regACCtop := TLB + HB + HB + HB + HB + HB + HB + HB + GTP + HB + HB + HB + HB + HB + HB + HB + HB + HB + TRB
	regACCbot := BLB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + HB + BRB

	ABtoB1LpipeSet := HP + GLP + GRP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP
	ABtoB1RpipeSet := HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP
	ABtoB2pipeSet := HP + GLP + GRP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP

	ABB0 := BLP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP
	ABB0 += HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP
	ABB0 += HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + HP + BRP

	OPCODE := " ADD  "
	INSTRUCTION := "  ADD R0,R1,R2   "

	fmt.Print(cursor.ClearEntireScreen())
	fmt.Print(cursor.MoveTo(0, 0))

	fmt.Printf("   %10s%s%1s%s%15s%s%2s%s%7s%s%2s%s%11s%s%2s%s", SP, "DAB", SP, hex, SP, "B1", SP, hex, SP, "B2", SP, hex, SP, "DAB", SP, hex)
	fmt.Println()
	fmt.Printf("   %14s%s%20s%s%12s%s%17s%s", SP, dec, SP, dec, SP, dec, SP, dec)
	fmt.Println()
	fmt.Printf("   %10s%s%15s%s%7s%s%12s%s", SP, bline, SP, bline, SP, bline, SP, bline)
	fmt.Println()
	fmt.Printf("   %15s%s%24s%s%16s%s%21s%s%27s%s", SP, VP, SP, VP, SP, VP, SP, VP, SP, "Memory")
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%s%s%s%s%23s%s", SP, VP, SP, regTop, SP, VRP, ABtoB1LpipeSet, VP, ABtoB1RpipeSet, VLP, SP, regTop) //one
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%s%s%23s%s%13s%s",
		SP, VP, SP, VB, SP, "R0", SP, hex, SP, regRgate, VLP, SP, VRP, ABtoB2pipeSet, VLP, SP, VB, SP, VB) //two
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%21s%s%23s%s%13s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VP, SP, VP, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%23s%s%13s%s", SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%13s%s", SP, VP, SP, regLine, SP, VP, SP, VP, SP, VP, SP, regTop, SP, VB, SP, VB)
	fmt.Println()

	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%21s%s%4s%s%s%s%s%s%3s%s%4s%s%4s%s%3s%s",
		SP, VP, SP, VB, SP, "R1", SP, hex, SP, regRgate, VLP, SP, VP, SP, VP, SP, VB, SP, "IA", SP, hex, SP, VB, SP, VB, SP, hex, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%4s%5s%4s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VP, SP, regLgate, SP, dec, SP, reg2reggate, SP, SP, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%4s%s%5s%3s%5s%s%4s%s%13s%s",
		SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, VB, SP, SP, SP, VB, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%13s%s", SP, VP, SP, regLine, SP, VP, SP, VP, SP, VP, SP, regLine, SP, VB, SP, VB)
	fmt.Println()

	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%21s%s%4s%s%s%s%s%s%3s%s%4s%s%4s%s%3s%s",
		SP, VP, SP, VB, SP, "R2", SP, hex, SP, regRgate, VLP, SP, VP, SP, VP, SP, VB, SP, "DA", SP, hex, SP, VB, SP, VB, SP, hex, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%4s%s%4s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VP, SP, regLgate, SP, dec, SP, reg2reggate, SP, dec, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%4s%s%5s%3s%5s%s%4s%s%13s%s",
		SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, VB, SP, SP, SP, VB, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%13s%s", SP, VP, SP, regLine, SP, VP, SP, VP, SP, VP, SP, regLine, SP, VB, SP, VB)
	fmt.Println()

	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%21s%s%4s%s%s%s%s%s%3s%s%4s%s%4s%s%3s%s",
		SP, VP, SP, VB, SP, "R3", SP, hex, SP, regRgate, VLP, SP, VP, SP, VP, SP, VB, SP, "DR", SP, hex, SP, VB, SP, VB, SP, hex, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%4s%s%4s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VP, SP, regLgate, SP, dec, SP, reg2reggate, SP, dec, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%4s%s%5s%3s%5s%s%4s%s%13s%s",
		SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, VB, SP, SP, SP, VB, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%13s%s", SP, VP, SP, regLine, SP, VP, SP, VP, SP, VP, SP, regBot, SP, VB, SP, VB)
	fmt.Println()

	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%21s%s%23s%s%13s%s",
		SP, VP, SP, VB, SP, "R4", SP, hex, SP, regRgate, VLP, SP, VP, SP, VP, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%21s%s%23s%s%13s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VP, SP, VP, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%23s%s%13s%s",
		SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%13s%s", SP, VP, SP, regLine, SP, VP, SP, VP, SP, VP, SP, regTop, SP, VB, SP, VB)
	fmt.Println()

	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%21s%s%4s%s%s%s%s%6s%3s%s%4s%s%4s%s%3s%s",
		SP, VP, SP, VB, SP, "R5", SP, hex, SP, regRgate, VLP, SP, VP, SP, VP, SP, VB, SP, "IR", SP, SP, SP, VB, SP, VB, SP, hex, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%s%s%s%4s%s%4s%s%4s%5s%4s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VRP, regVgateSet, VP, regVpipe, SP, dec, SP, reg2reggate, SP, SP, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%4s%s%5s%3s%5s%s%4s%s%13s%s",
		SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, VB, SP, SP, SP, VB, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%4s%s%4s%s%13s%s", SP, VP, SP, regLine, SP, VP, SP, VP, SP, VP, SP, regBot, SP, VB, SP, VB)
	fmt.Println()

	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%21s%s%2s%s%2s%s%13s%s",
		SP, VP, SP, VB, SP, "R6", SP, hex, SP, regRgate, VLP, SP, VP, SP, VP, SP, regIRtop, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%21s%s%2s%s%s%s%2s%s%13s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VP, SP, VP, SP, VB, INSTRUCTION, VB, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%2s%s%2s%s%13s%s",
		SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, regIRbot, SP, VB, SP, VB)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%9s%s%5s%s%3s%s%3s%s%13s%s", SP, VP, SP, regLine, SP, VP, SP, VP, SP, VP, SP, VB, SP, VB, SP, VB, SP, VB, SP, VB)
	fmt.Println()

	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%21s%s%8s%s%4s%s%1s%s",
		SP, VP, SP, VB, SP, "R7", SP, hex, SP, regRgate, VLP, SP, VP, SP, VP, SP, "OPC", SP, "DR", SP, "DW")
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%21s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VP, SP, VP)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%4s%s",
		SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, regTop)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%s%1s%s%4s%s%s%s%s",
		SP, VP, SP, regLine, SP, VP, SP, VP, SP, regLgate, "OUT", SP, dec, SP, VRB, HB, HB, RARROW) //
	fmt.Println()

	fmt.Printf("   %15s%s%4s%s%s%s%s%s%3s%s%s%16s%s%21s%s%4s%s",
		SP, VP, SP, VB, SP, "SR", SP, hex, SP, regRgate, VLP, SP, VP, SP, VP, SP, regBot)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%4s%s%5s%s%16s%s%21s%s",
		SP, regLgate, SP, dec, SP, VB, SP, VP, SP, VP, SP, VP)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%13s%s%s%s%s%21s%s%4s%s",
		SP, VP, SP, VB, SP, regRgate, VP, regRpipeSet, VLP, SP, VP, SP, regTop)
	fmt.Println()
	fmt.Printf("   %15s%s%4s%s%5s%s%16s%s%21s%s%s%1s%s%4s%s%s%s%s",
		SP, VP, SP, regBot, SP, VP, SP, VP, SP, regLgate, " IN", SP, dec, SP, VRB, HB, HB, LARROW) //
	fmt.Println()

	fmt.Printf("   %15s%s%21s%s%1s%12s%s%4s%s", SP, VP, SP, regALUtop, OPCODE, SP, VP, SP, regBot)
	fmt.Println()

	fmt.Printf("   %15s%s%17s%s%s%2s%s%8s%s%2s%s%s%s%s%s%s%s%s%8s%s", SP, VP, SP, "ALU ", VB, SP, dec, SP, dec, SP, VB, HB, HB, HB, HB, HB, HB, " OPC", SP, VP)
	fmt.Println()

	fmt.Printf("   %15s%s%21s%s%18s%s", SP, VP, SP, regALUbot, SP, VP)
	fmt.Println()

	fmt.Printf("   %15s%s%24s%s%s%6s%s", SP, VP, SP, regACCtop, "   GT LT NE EQ", SP, VP)
	fmt.Println()

	fmt.Printf("   %15s%s%20s%s%s%5s%s%7s%s%3s%s%7s%s", SP, VP, SP, "ACC ", VB, SP, dec, SP, VB, SP, FLAGS, SP, VP)
	fmt.Println()

	fmt.Printf("   %15s%s%24s%s%20s%s", SP, VP, SP, regACCbot, SP, VP)
	fmt.Println()

	fmt.Printf("   %15s%s", SP, ABB0)
	fmt.Println()

	//fmt.Println("012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789")

}

func (core *Config) resetCore() {

	core.RegValues[R0] = 0
	core.RegValues[R1] = 0
	core.RegValues[R2] = 0
	core.RegValues[R3] = 0
	core.RegValues[R4] = 0
	core.RegValues[R5] = 0
	core.RegValues[R6] = 0
	core.RegValues[R7] = 0
	core.RegValues[SR] = 0
	core.RegValues[IA] = 0
	core.RegValues[IR] = 0
	core.RegValues[DA] = 0
	core.RegValues[DR] = 0
	core.RegValues[B1] = 0
	core.RegValues[B2] = 0
	core.RegValues[DABA] = 0
	core.RegValues[DABB] = 0
	core.RegValues[IN] = 0
	core.RegValues[OUT] = 0
	core.RegValues[ALUL] = 0
	core.RegValues[ALUR] = 0
	core.RegValues[ACC] = 0
	core.RegValues[FLAG] = 0
	core.RegValues[INST] = 0

	core.CoreMemPoint = 0

	core.SleepTime = 0

	cntlArray := []ControlReg{core.IA, core.DA, core.DR}

	ioCntlArray := []IOReg{core.OUTPUTR, core.INPUTR}

	busGateArray := []Gate{core.B1DABGate, core.B2DABGate}

	fmt.Print(cursor.Hide())

	fmt.Print(cursor.MoveTo(core.DABA.hexRow, core.DABA.hexCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABA.hexRow, core.DABA.hexCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABA.hexRow, core.DABA.hexCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABA.decRow, core.DABA.decCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABA.decRow, core.DABA.decCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABA.decRow, core.DABA.decCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.B1.hexRow, core.B1.hexCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B1.hexRow, core.B1.hexCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B1.hexRow, core.B1.hexCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B1.decRow, core.B1.decCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B1.decRow, core.B1.decCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B1.decRow, core.B1.decCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.B2.hexRow, core.B2.hexCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B2.hexRow, core.B2.hexCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B2.hexRow, core.B2.hexCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B2.decRow, core.B2.decCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B2.decRow, core.B2.decCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.B2.decRow, core.B2.decCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.DABB.hexRow, core.DABB.hexCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABB.hexRow, core.DABB.hexCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABB.hexRow, core.DABB.hexCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABB.decRow, core.DABB.decCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABB.decRow, core.DABB.decCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.DABB.decRow, core.DABB.decCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	for reg := range core.GenRegs {
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].ingateRow, core.GenRegs[reg].ingateCol))
		fmt.Printf("%s", closeGate)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateTopRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", closeGate)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateBotRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", closeGate)
		time.Sleep(core.SleepTime * time.Millisecond)

		fmt.Print(cursor.MoveTo(core.GenRegs[reg].hexRow, core.GenRegs[reg].valueCol))
		fmt.Printf("%9s", SP)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].hexRow, core.GenRegs[reg].valueCol))
		fmt.Printf("0x%04X", 65535)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].hexRow, core.GenRegs[reg].valueCol))
		fmt.Printf("0x%04X", core.RegValues[reg])
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].decRow, core.GenRegs[reg].valueCol))
		fmt.Printf("%9s", SP)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].decRow, core.GenRegs[reg].valueCol))
		fmt.Printf("%05d", 65535)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].decRow, core.GenRegs[reg].valueCol))
		fmt.Printf("%05d", core.RegValues[reg])
		time.Sleep(core.SleepTime * time.Millisecond)
	}

	time.Sleep(core.SleepTime * time.Millisecond)

	for reg := range core.GenRegs {
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].ingateRow, core.GenRegs[reg].ingateCol))
		fmt.Printf("%s", openGate)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateTopRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", openGate)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateBotRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", openGate)
		time.Sleep(core.SleepTime * time.Millisecond)
	}

	for reg := range cntlArray {
		fmt.Print(cursor.MoveTo(cntlArray[reg].gateRow, cntlArray[reg].ingateCol))
		fmt.Printf("%s", closeGate)
		fmt.Print(cursor.MoveTo(cntlArray[reg].gateRow, cntlArray[reg].outgateCol))
		fmt.Printf("%s", closeGate)
		time.Sleep(core.SleepTime * time.Millisecond)

		fmt.Print(cursor.MoveTo(cntlArray[reg].hexRow, cntlArray[reg].valueCol))
		fmt.Printf("%9s", SP)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(cntlArray[reg].hexRow, cntlArray[reg].valueCol))
		fmt.Printf("0x%04X", 65535)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(cntlArray[reg].hexRow, cntlArray[reg].valueCol))
		fmt.Printf("0x%04X", 0)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(cntlArray[reg].decRow, cntlArray[reg].valueCol))
		fmt.Printf("%9s", SP)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(cntlArray[reg].decRow, cntlArray[reg].valueCol))
		fmt.Printf("%05d", 65535)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(cntlArray[reg].decRow, cntlArray[reg].valueCol))
		fmt.Printf("%05d", 0)
		time.Sleep(core.SleepTime * time.Millisecond)
	}

	fmt.Print(cursor.MoveTo(core.IR.hexRow, core.IR.valueCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.IR.hexRow, core.IR.valueCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.IR.hexRow, core.IR.valueCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	for reg := range ioCntlArray {
		fmt.Print(cursor.MoveTo(ioCntlArray[reg].gateRow, ioCntlArray[reg].gateCol))
		fmt.Printf("%s", closeGate)
		fmt.Print(cursor.MoveTo(ioCntlArray[reg].gateRow, ioCntlArray[reg].gateCol))
		fmt.Printf("%s", closeGate)

		fmt.Print(cursor.MoveTo(ioCntlArray[reg].decRow, ioCntlArray[reg].valueCol))
		fmt.Printf("%9s", SP)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(ioCntlArray[reg].decRow, ioCntlArray[reg].valueCol))
		fmt.Printf("%05d", 65535)
		time.Sleep(core.SleepTime * time.Millisecond)
		fmt.Print(cursor.MoveTo(ioCntlArray[reg].decRow, ioCntlArray[reg].valueCol))
		fmt.Printf("%05d", 0)
		time.Sleep(core.SleepTime * time.Millisecond)
	}

	for reg := range busGateArray {
		fmt.Print(cursor.MoveTo(busGateArray[reg].row, busGateArray[reg].col))
		fmt.Printf("%s", closeGate)
		time.Sleep(core.SleepTime * time.Millisecond)
	}

	for reg := range core.GenRegs {
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].ingateRow, core.GenRegs[reg].ingateCol))
		fmt.Printf("%s", openGate)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateTopRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", openGate)
		fmt.Print(cursor.MoveTo(core.GenRegs[reg].outgateBotRow, core.GenRegs[reg].outgateCol))
		fmt.Printf("%s", openGate)
		time.Sleep(core.SleepTime * time.Millisecond)
	}
	for reg := range cntlArray {
		fmt.Print(cursor.MoveTo(cntlArray[reg].gateRow, cntlArray[reg].ingateCol))
		fmt.Printf("%s", openGate)
		fmt.Print(cursor.MoveTo(cntlArray[reg].gateRow, cntlArray[reg].outgateCol))
		fmt.Printf("%s", openGate)
		time.Sleep(core.SleepTime * time.Millisecond)
	}

	for reg := range ioCntlArray {
		fmt.Print(cursor.MoveTo(ioCntlArray[reg].gateRow, ioCntlArray[reg].gateCol))
		fmt.Printf("%s", openGate)
		fmt.Print(cursor.MoveTo(ioCntlArray[reg].gateRow, ioCntlArray[reg].gateCol))
		fmt.Printf("%s", openGate)
		time.Sleep(core.SleepTime * time.Millisecond)
	}

	for reg := range busGateArray {
		fmt.Print(cursor.MoveTo(busGateArray[reg].row, busGateArray[reg].col))
		fmt.Printf("%s", openGate)
		time.Sleep(core.SleepTime * time.Millisecond)
	}

	fmt.Print(cursor.MoveTo(core.Instruction.row, core.Instruction.col))
	fmt.Printf("%14s", SP)
	fmt.Print(cursor.MoveTo(core.Instruction.row, core.Instruction.col))
	fmt.Printf("%s", "          ")
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.ALU.decRow, core.ALU.leftDecCol))
	fmt.Printf("%6s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ALU.decRow, core.ALU.leftDecCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ALU.decRow, core.ALU.leftDecCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.ALU.decRow, core.ALU.rightDecCol))
	fmt.Printf("%6s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ALU.decRow, core.ALU.rightDecCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ALU.decRow, core.ALU.rightDecCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.ALU.opCodeRow, core.ALU.opCodeCol))
	fmt.Printf("%6s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ALU.opCodeRow, core.ALU.opCodeCol))
	fmt.Printf("%s", "      ")
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.ACC.decRow, core.ACC.decCol))
	fmt.Printf("%6s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ACC.decRow, core.ACC.decCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ACC.decRow, core.ACC.decCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.ACC.flagsRow, core.ACC.flagsCol))
	fmt.Printf("%11s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ACC.flagsRow, core.ACC.flagsCol))
	fmt.Printf("%d%2s%d%2s%d%2s%d", 0, SP, 1, SP, 0, SP, 1)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.ACC.flagsRow, core.ACC.flagsCol))
	fmt.Printf("%d%2s%d%2s%d%2s%d", 0, SP, 0, SP, 0, SP, 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.MEM.iaHexRow, core.MEM.memoryCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.iaHexRow, core.MEM.memoryCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.iaHexRow, core.MEM.memoryCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.MEM.daHexRow, core.MEM.memoryCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.daHexRow, core.MEM.memoryCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.daHexRow, core.MEM.memoryCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.daDecRow, core.MEM.memoryCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.daDecRow, core.MEM.memoryCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.daDecRow, core.MEM.memoryCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.MEM.drHexRow, core.MEM.memoryCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.drHexRow, core.MEM.memoryCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.drHexRow, core.MEM.memoryCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.drDecRow, core.MEM.memoryCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.drDecRow, core.MEM.memoryCol))
	fmt.Printf("%05d", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.drDecRow, core.MEM.memoryCol))
	fmt.Printf("%05d", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(core.MEM.irHexRow, core.MEM.memoryCol))
	fmt.Printf("%9s", SP)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.irHexRow, core.MEM.memoryCol))
	fmt.Printf("0x%04X", 65535)
	time.Sleep(core.SleepTime * time.Millisecond)
	fmt.Print(cursor.MoveTo(core.MEM.irHexRow, core.MEM.memoryCol))
	fmt.Printf("0x%04X", 0)
	time.Sleep(core.SleepTime * time.Millisecond)

	fmt.Print(cursor.MoveTo(2, 141))
	//fmt.Print(cursor.MoveTo(2, 136))
	fmt.Printf("%2s", SP)

	//core.clearIO()
	core.cursorHome()
	//core.setOutputHome()
	fmt.Print(cursor.Show())
}
