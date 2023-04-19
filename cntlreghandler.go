package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) cntlRegHandler(reg int, cmd string) {

	switch cmd {
	case "show":
		switch reg {
		case int(DA):
			fmt.Print(cursor.MoveTo(core.DA.hexRow, core.DA.valueCol))
			fmt.Printf("0x%04X", core.RegValues[DA])

			fmt.Print(cursor.MoveTo(core.DA.decRow, core.DA.valueCol))
			fmt.Printf("%05d", core.RegValues[DA])
		case int(DR):
			fmt.Print(cursor.MoveTo(core.DR.hexRow, core.DR.valueCol))
			fmt.Printf("0x%04X", core.RegValues[DR])

			fmt.Print(cursor.MoveTo(core.DR.decRow, core.DR.valueCol))
			fmt.Printf("%05d", core.RegValues[DR])
		case int(IA):
			fmt.Print(cursor.MoveTo(core.IA.hexRow, core.IA.valueCol))
			fmt.Printf("0x%04X", core.RegValues[IA])

			fmt.Print(cursor.MoveTo(core.IA.decRow, core.IA.valueCol))
			fmt.Printf("%05d", core.RegValues[IA])
		case int(IR):
			fmt.Print(cursor.MoveTo(core.IR.hexRow, core.IR.valueCol))
			fmt.Printf("0x%04X", core.RegValues[IR])
		}
	case "closeout":
		switch reg {
		case int(DA):
			fmt.Print(cursor.MoveTo(core.DA.gateRow, core.DA.outgateCol))
			fmt.Printf("%s", closeGate)
			fmt.Print(cursor.MoveTo(core.DA.directRow, core.DA.directCol))
			fmt.Printf("%s", OUTDIRECTION)
		case int(DR):
			fmt.Print(cursor.MoveTo(core.DR.gateRow, core.DR.outgateCol))
			fmt.Printf("%s", closeGate)
		case int(IA):
			fmt.Print(cursor.MoveTo(core.IA.gateRow, core.IA.outgateCol))
			fmt.Printf("%s", closeGate)
		case int(IR):
			fmt.Print(cursor.MoveTo(core.IR.gateRow, core.IR.outgateCol))
			fmt.Printf("%s", closeGate)
			fmt.Print(cursor.MoveTo(core.IR.directRow, core.IR.directCol))
			fmt.Printf("%s", INDIRECTION)
		}
	case "closein":
		switch reg {
		case int(DA):
			fmt.Print(cursor.MoveTo(core.DA.gateRow, core.DA.ingateCol))
			fmt.Printf("%s", closeGate)
			fmt.Print(cursor.MoveTo(core.DA.directRow, core.DA.directCol))
			fmt.Printf("%s", OUTDIRECTION)
		case int(DR):
			fmt.Print(cursor.MoveTo(core.DR.gateRow, core.DR.ingateCol))
			fmt.Printf("%s", closeGate)
		case int(IA):
			fmt.Print(cursor.MoveTo(core.IA.gateRow, core.IA.ingateCol))
			fmt.Printf("%s", closeGate)
		case int(IR):
			fmt.Print(cursor.MoveTo(core.IR.gateRow, core.IR.ingateCol))
			fmt.Printf("%s", closeGate)
			fmt.Print(cursor.MoveTo(core.IR.directRow, core.IR.directCol))
			fmt.Printf("%s", INDIRECTION)
		}
	case "openout":
		switch reg {
		case int(DA):
			fmt.Print(cursor.MoveTo(core.DA.gateRow, core.DA.outgateCol))
			fmt.Printf("%s", openGate)
			fmt.Print(cursor.MoveTo(core.DA.directRow, core.DA.directCol))
			fmt.Printf("%3s", SP)
		case int(DR):
			fmt.Print(cursor.MoveTo(core.DR.gateRow, core.DR.outgateCol))
			fmt.Printf("%s", openGate)
			fmt.Print(cursor.MoveTo(core.DR.directRow, core.DR.directCol))
			fmt.Printf("%3s", SP)
		case int(IA):
			fmt.Print(cursor.MoveTo(core.IA.gateRow, core.IA.outgateCol))
			fmt.Printf("%s", openGate)
			fmt.Print(cursor.MoveTo(core.IA.directRow, core.IA.directCol))
			fmt.Printf("%3s", SP)
		case int(IR):
			fmt.Print(cursor.MoveTo(core.IR.gateRow, core.IR.outgateCol))
			fmt.Printf("%s", openGate)
			fmt.Print(cursor.MoveTo(core.IR.directRow, core.IR.directCol))
			fmt.Printf("%3s", SP)
		}
	case "openin":
		switch reg {
		case int(DA):
			fmt.Print(cursor.MoveTo(core.DA.gateRow, core.DA.ingateCol))
			fmt.Printf("%s", openGate)
			fmt.Print(cursor.MoveTo(core.DA.directRow, core.DA.directCol))
			fmt.Printf("%3s", SP)
		case int(DR):
			fmt.Print(cursor.MoveTo(core.DR.gateRow, core.DR.ingateCol))
			fmt.Printf("%s", openGate)
			fmt.Print(cursor.MoveTo(core.DR.directRow, core.DR.directCol))
			fmt.Printf("%3s", SP)
		case int(IA):
			fmt.Print(cursor.MoveTo(core.IA.gateRow, core.IA.ingateCol))
			fmt.Printf("%s", openGate)
			fmt.Print(cursor.MoveTo(core.IA.directRow, core.IA.directCol))
			fmt.Printf("%3s", SP)
		case int(IR):
			fmt.Print(cursor.MoveTo(core.IR.gateRow, core.IR.ingateCol))
			fmt.Printf("%s", openGate)
			fmt.Print(cursor.MoveTo(core.IR.directRow, core.IR.directCol))
			fmt.Printf("%3s", SP)
		}
	case "indirection":
		switch reg {
		case int(DR):
			fmt.Print(cursor.MoveTo(core.DR.directRow, core.DR.directCol))
			fmt.Printf("%s", INDIRECTION)
		case int(IA):
			fmt.Print(cursor.MoveTo(core.IA.directRow, core.IA.directCol))
			fmt.Printf("%s", INDIRECTION)
		}

	case "outdirection":
		switch reg {
		case int(DR):
			fmt.Print(cursor.MoveTo(core.DR.directRow, core.DR.directCol))
			fmt.Printf("%s", OUTDIRECTION)
		case int(IA):
			fmt.Print(cursor.MoveTo(core.IA.directRow, core.IA.directCol))
			fmt.Printf("%s", OUTDIRECTION)

		}
	}
}
