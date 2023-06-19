package assembler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ahmetalpbalkan/go-cursor"
)

type Assembler struct {
	asmFileName    string
	docFileName    string
	hexFileName    string
	docFile        *os.File
	hexFile        *os.File
	labelMap       map[string]uint16
	opCodeMap      map[string]uint16
	byteMap        map[string]uint16
	regOneMap      map[string]uint16
	regTwoMap      map[string]uint16
	regThreeMap    map[string]uint16
	shiftMap       map[string]uint16
	flagMap        map[string]uint16
	opcClassMap    map[string]string
	DebugMap       map[int]string
	strConstMap    map[string][]byte
	strConstAdrMap map[string]uint16
	adrMemPointer  uint16
	CoreMemory     []uint16
	lineCount      int
	homeColunm     int
}

func (asm *Assembler) Assemble(homeColumn int, CoreMemory []uint16, asmCmd string, CoreDebugMap map[int]string) bool {

	asm.homeColunm = homeColumn
	asm.CoreMemory = CoreMemory
	asm.DebugMap = CoreDebugMap

	asm.asmFileName = "source.asm"
	asm.docFileName = "source.doc"
	asm.hexFileName = "source.hex"

	asm.regTwoMap = make(map[string]uint16)
	asm.regThreeMap = make(map[string]uint16)
	asm.labelMap = make(map[string]uint16)
	asm.opCodeMap = make(map[string]uint16)
	asm.byteMap = make(map[string]uint16)
	asm.regOneMap = make(map[string]uint16)
	asm.opcClassMap = make(map[string]string)

	asm.shiftMap = make(map[string]uint16)
	asm.flagMap = make(map[string]uint16)
	asm.strConstMap = make(map[string][]byte)
	asm.strConstAdrMap = make(map[string]uint16)

	asm.adrMemPointer = 0

	asm.loadOpCodeMap()
	asm.loadRegOneMap()
	asm.loadRegTwoMap()
	asm.loadRegThreeMap()
	asm.loadShiftMap()
	asm.loadFlagMap()
	asm.loadOpcClassMap()

	if asmCmd == "assemble" {
		if !asm.passOne() {
			return false
		}
		return asm.passTwo()
	}
	if asmCmd == "loadCore" {
		return asm.loadCoreMemFromFile()
	}
	return true
}

func (asm *Assembler) loadCoreMemFromFile() bool {

	hexFile, err := os.Open(asm.hexFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer hexFile.Close()
	asm.adrMemPointer = 0

	fileScanner := bufio.NewScanner(hexFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {

		someString := fileScanner.Text()
		asm.lineCount++
		words := strings.Fields(someString)

		parseValue, err := strconv.ParseUint(words[0], 16, 16)
		if err != nil {
			fmt.Print(cursor.Hide())
			fmt.Print(cursor.MoveTo(45, 140))
			fmt.Print(strings.ToUpper("Failed to load hexFile into CoreMemory"))
			fmt.Print(cursor.MoveTo(46, 140))
			fmt.Print(strings.ToUpper(err.Error()))
			fmt.Print(cursor.MoveTo(22, asm.homeColunm))
			fmt.Print(cursor.Show())
			return false
		}

		asm.CoreMemory[asm.adrMemPointer] = uint16(parseValue)
		asm.adrMemPointer++
	}

	fmt.Print(cursor.Hide())
	fmt.Print(cursor.MoveTo(45, 140))
	fmt.Print("                                                          ")
	fmt.Print(cursor.MoveTo(46, 140))
	fmt.Print("                                                                                 ")
	fmt.Print(cursor.Show())
	fmt.Print(cursor.MoveTo(22, asm.homeColunm))

	return true
}

func (asm *Assembler) outputErrorMsg(statement []string, msg string) {

	fmt.Print(cursor.Hide())
	fmt.Print(cursor.MoveTo(45, 140))
	fmt.Print("                                                          ")
	fmt.Print(cursor.MoveTo(46, 140))
	fmt.Print("                                                                                 ")

	fmt.Print(cursor.MoveTo(45, 140))
	fmt.Print(statement)
	fmt.Print(cursor.MoveTo(46, 140))
	fmt.Printf(strings.ToUpper(msg)+" %d", asm.lineCount)

	fmt.Print(cursor.MoveTo(22, asm.homeColunm))
	fmt.Print(cursor.Show())
}

func (asm *Assembler) passTwo() bool {

	asmFile, err := os.Open(asm.asmFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	docFile, err := os.Create(asm.docFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	asm.docFile = docFile
	hexFile, err := os.Create(asm.hexFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	asm.hexFile = hexFile
	defer asmFile.Close()
	defer docFile.Close()
	defer hexFile.Close()

	asm.docFile.WriteString("\n")
	asm.docFile.WriteString("\n")
	asm.docFile.WriteString("\n")
	asm.docFile.WriteString("\n")
	for key, value := range asm.byteMap {
		tmpString := adjustStringLength(key, 15, "prepend")
		tmpString = adjustStringLength(tmpString, 40, "append")
		asm.docFile.WriteString(tmpString + "   DBYTE   ")
		asm.docFile.WriteString(fmt.Sprintf("%d  ", value))
		asm.docFile.WriteString("\n")
	}
	asm.docFile.WriteString("\n")
	asm.docFile.WriteString("\n")

	for key, value := range asm.labelMap {
		tmpString := adjustStringLength(key, 15, "prepend")
		tmpString = adjustStringLength(tmpString, 40, "append")
		asm.docFile.WriteString(tmpString + "   DBYTE   ")
		asm.docFile.WriteString(fmt.Sprintf("%d  ", value))
		asm.docFile.WriteString("\n")
	}
	asm.docFile.WriteString("\n")
	asm.docFile.WriteString("\n")

	for key, value := range asm.strConstMap {
		tmpString := adjustStringLength(key, 15, "prepend")
		tmpString = adjustStringLength(tmpString, 40, "append")
		asm.docFile.WriteString(tmpString + "   DSTRING ")
		asm.docFile.WriteString(fmt.Sprintf("%s  ", value))
		asm.docFile.WriteString("\n")
	}
	asm.docFile.WriteString("\n")
	asm.docFile.WriteString("\n")

	asm.lineCount = 0
	asm.adrMemPointer = 0

	fileScanner := bufio.NewScanner(asmFile)
	fileScanner.Split(bufio.ScanLines)

	var okPassTwo bool = true

	for fileScanner.Scan() {

		errorMessage := ""

		someString := fileScanner.Text()
		asm.lineCount++
		words := strings.Fields(someString)

		var statement []string

		for w := range words {
			if strings.HasPrefix(words[w], ";") {
				break
			} else {
				statement = append(statement, strings.ToUpper(words[w]))
			}
		}

		if len(statement) > 0 {

			var opcIndex int

			if strings.HasSuffix(statement[0], ":") {
				opcIndex = 1
			} else {
				opcIndex = 0
			}

			opcClass, ok := asm.opcClassMap[statement[opcIndex]]
			if ok {
				switch opcClass {
				case "0args":
					if !asm.opCodeNoArg(statement, opcIndex, &errorMessage) {
						okPassTwo = false
						asm.outputErrorMsg(statement, errorMessage)
					}
				case "1args":
					if !asm.opCodeOneArg(statement, opcIndex, &errorMessage) {
						okPassTwo = false
						asm.outputErrorMsg(statement, errorMessage)
					}
				case "2args":
					if !asm.opCodeTwoArg(statement, opcIndex, &errorMessage) {
						okPassTwo = false
						asm.outputErrorMsg(statement, errorMessage)
					}
				case "3args":
					if !asm.opCodeThreeArg(statement, opcIndex, &errorMessage) {
						okPassTwo = false
						asm.outputErrorMsg(statement, errorMessage)
					}
				case "skip":
					// do nothing
				default:
					okPassTwo = false
					asm.outputErrorMsg(statement, "Invalid statement. line ")
				}
			} else {
				okPassTwo = false
				asm.outputErrorMsg(statement, "Invalid statement. line ")
			}

		}

	}
	if okPassTwo {

		for key, value := range asm.strConstMap {
			strConstAdr := asm.strConstAdrMap[key]

			for idx := 0; idx < len(value); idx++ {
				asm.CoreMemory[strConstAdr] = uint16(value[idx])
				strConstAdr++
			}
			asm.CoreMemory[strConstAdr] = 0
		}

		fmt.Print(cursor.Hide())
		fmt.Print(cursor.MoveTo(45, 140))
		fmt.Print("                                                          ")
		fmt.Print(cursor.MoveTo(46, 140))
		fmt.Print("                                                                                 ")
		fmt.Print(cursor.Show())
		fmt.Print(cursor.MoveTo(22, asm.homeColunm))
	}
	return okPassTwo
}

func adjustStringLength(code string, length int, direction string) string {
	var retString string = ""
	if direction == "append" {
		retString = code
		for idx := 0; idx < length-len(code); idx++ {
			retString = retString + " "
		}
	} else if direction == "prepend" {
		for idx := 0; idx < length; idx++ {
			retString = retString + " "
		}
		retString = retString + code
	}

	return retString
}

func (asm *Assembler) writeToDocFile(statement []string, opCodeHex uint16) {
	asm.docFile.WriteString(fmt.Sprintf("%05d  ", asm.adrMemPointer))
	asm.docFile.WriteString(fmt.Sprintf("0x%04X  ", opCodeHex))

	var debugString string = fmt.Sprintf("%05d  ", asm.adrMemPointer)
	debugString += fmt.Sprintf("0x%04X  ", opCodeHex)

	if strings.HasSuffix(statement[0], ":") {
		for s := range statement {
			if s == 0 {
				asm.docFile.WriteString(adjustStringLength(statement[s], 18, "append"))
				debugString += adjustStringLength(statement[s], 18, "append")
			} else if s == 1 {
				asm.docFile.WriteString(adjustStringLength(statement[s], 10, "append"))
				debugString += adjustStringLength(statement[s], 10, "append")
			} else {
				asm.docFile.WriteString(statement[s] + " ")
				debugString += statement[s] + " "
			}
		}
	} else {
		for s := range statement {
			if s == 0 {
				tmpString := adjustStringLength(statement[s], 18, "prepend")
				tmpString = adjustStringLength(tmpString, 28, "append")

				asm.docFile.WriteString(tmpString)
				debugString += tmpString
			} else {
				asm.docFile.WriteString(statement[s] + " ")
				debugString += statement[s] + " "
			}
		}
	}

	asm.DebugMap[int(asm.adrMemPointer)] = debugString
	asm.docFile.WriteString("\n")
}

func (asm *Assembler) opCodeThreeArg(statement []string, opCodeIdx int, msg *string) bool {

	if len(statement) < 4 {
		*msg = "too few arguments. line "
		return false
	}
	switch statement[opCodeIdx] {
	case "SHL":
		fallthrough
	case "SHR":
		opCodeHex, okOpc := asm.regTwoMap[statement[opCodeIdx]]
		regOneHex, okRegOne := asm.regOneMap[statement[opCodeIdx+1]]
		regTwoHex, okRegTwo := asm.regTwoMap[statement[opCodeIdx+2]]
		shiftValue, okShift := asm.shiftMap[statement[opCodeIdx+3]]

		if okOpc && okRegOne && okRegTwo && okShift {
			opCodeHex = opCodeHex + regOneHex + regTwoHex + shiftValue
			asm.CoreMemory[asm.adrMemPointer] = opCodeHex

			asm.hexFile.WriteString(fmt.Sprintf("%04X  ", opCodeHex) + "\n")

			asm.writeToDocFile(statement, opCodeHex)
			asm.adrMemPointer++
		} else {
			if !okRegOne || !okRegTwo {
				*msg = "invalid register name: R0-R7. line "
			} else if !okShift {
				*msg = "Invalid shift value: SH(L/R) R0 R1 V(0-15). line "
			}
			return false
		}

	default:
		opCodeHex, okOpc := asm.regThreeMap[statement[opCodeIdx]]
		regOneHex, okRegOne := asm.regOneMap[statement[opCodeIdx+1]]
		regTwoHex, okRegTwo := asm.regTwoMap[statement[opCodeIdx+2]]
		regThreeHex, okRegThree := asm.regThreeMap[statement[opCodeIdx+3]]

		if okOpc && okRegOne && okRegTwo && okRegThree {
			opCodeHex = opCodeHex + regOneHex + regTwoHex + regThreeHex
			asm.CoreMemory[asm.adrMemPointer] = opCodeHex

			asm.hexFile.WriteString(fmt.Sprintf("%04X  ", opCodeHex) + "\n")

			asm.writeToDocFile(statement, opCodeHex)
			asm.adrMemPointer++
		} else {
			if !okRegOne || !okRegTwo || !okRegThree {
				*msg = "invalid register name: R0-R7. line "
			}
			return false
		}

	}

	return true
}

func (asm *Assembler) opCodeTwoArg(statement []string, opCodeIdx int, msg *string) bool {

	if len(statement) < 3 {
		*msg = "too few arguments. line "
		return false
	}
	switch statement[opCodeIdx] {
	case "NOT":
		fallthrough
	case "CMP":
		opCodeHex, okOpc := asm.regTwoMap[statement[opCodeIdx]]
		regOneHex, okRegOne := asm.regOneMap[statement[opCodeIdx+1]]
		regTwoHex, okRegTwo := asm.regTwoMap[statement[opCodeIdx+2]]
		if okOpc && okRegOne && okRegTwo {
			opCodeHex = opCodeHex + regOneHex + regTwoHex
			asm.CoreMemory[asm.adrMemPointer] = opCodeHex

			asm.hexFile.WriteString(fmt.Sprintf("%04X  ", opCodeHex) + "\n")

			asm.writeToDocFile(statement, opCodeHex)
			asm.adrMemPointer++
		} else {
			if !okRegOne || !okRegTwo {
				*msg = "invalid register name: R0-R7. line "
			}
			return false
		}
	case "BFLAG":
		opCodeHex, okOpc := asm.regOneMap[statement[opCodeIdx]]
		regHex, okReg := asm.regOneMap[statement[opCodeIdx+1]]
		flagHex, okFlag := asm.flagMap[statement[opCodeIdx+2]]
		if okOpc && okReg && okFlag {
			opCodeHex = opCodeHex + regHex + flagHex
			asm.CoreMemory[asm.adrMemPointer] = opCodeHex

			asm.hexFile.WriteString(fmt.Sprintf("%04X  ", opCodeHex) + "\n")

			asm.writeToDocFile(statement, opCodeHex)
			asm.adrMemPointer++
		} else {
			if !okReg {
				*msg = "invalid register name: R0-R7. line "
			} else if !okFlag {
				*msg = "invalid branch flag: (GT, LT, NE, EQ). line "
			}
			return false
		}
	case "ADDI":
		fallthrough
	case "SUBI":
		fallthrough
	case "ANDI":
		fallthrough
	case "XORI":
		fallthrough
	case "ORI":
		fallthrough
	case "LDI":
		fallthrough
	case "CMPI":
		opCodeHex, okOpc := asm.regOneMap[statement[opCodeIdx]]
		regHex, okReg := asm.regOneMap[statement[opCodeIdx+1]]
		if !okOpc || !okReg {
			*msg = "invalid register name: R0-R7. line "
			return false
		}

		var value uint16
		var okValue bool = false

		labelValue, okLabel := asm.labelMap[statement[opCodeIdx+2]]
		byteValue, okByte := asm.byteMap[statement[opCodeIdx+2]]
		if okLabel {
			value = labelValue
			okValue = true
		} else if okByte {
			value = byteValue
			okValue = true
		} else {
			parseValue, err := strconv.ParseUint(statement[opCodeIdx+2], 10, 16)
			if err != nil {
				*msg = "invalid argument: (Label name/DBYTE name/Value). line "
				return false
			}
			value = uint16(parseValue)
			okValue = true
		}

		if value > 255 {
			*msg = "illegal value: (0-255). line "
			return false
		}

		if okOpc && okReg && okValue {
			opCodeHex = opCodeHex + regHex + uint16(value)
			asm.CoreMemory[asm.adrMemPointer] = opCodeHex

			asm.hexFile.WriteString(fmt.Sprintf("%04X  ", opCodeHex) + "\n")

			asm.writeToDocFile(statement, opCodeHex)
			asm.adrMemPointer++
		} else {
			*msg = "invalid register name: R0-R7. line "
			return false
		}
	}

	return true
}

func (asm *Assembler) opCodeOneArg(statement []string, opCodeIdx int, msg *string) bool {

	if len(statement) < 2 {
		*msg = "too few arguments. line "
		return false
	}
	opCodeHex, okOpc := asm.regOneMap[statement[opCodeIdx]]
	regHex, okReg := asm.regOneMap[statement[opCodeIdx+1]]
	if okOpc && okReg {
		opCodeHex = opCodeHex + regHex
		asm.CoreMemory[asm.adrMemPointer] = opCodeHex

		asm.hexFile.WriteString(fmt.Sprintf("%04X  ", opCodeHex) + "\n")

		asm.writeToDocFile(statement, opCodeHex)
		asm.adrMemPointer++
	} else {
		*msg = "invalid register name: R0-R7. line "
		return false
	}

	return true
}

func (asm *Assembler) opCodeNoArg(statement []string, opCodeIdx int, msg *string) bool {

	opCodeHex, ok := asm.opCodeMap[statement[opCodeIdx]]
	if ok {
		asm.CoreMemory[asm.adrMemPointer] = opCodeHex

		asm.hexFile.WriteString(fmt.Sprintf("%04X  ", opCodeHex) + "\n")

		asm.writeToDocFile(statement, opCodeHex)
		asm.adrMemPointer++
	} else {
		*msg = "invalid statement. line "
		return false
	}
	return true
}

func (asm *Assembler) passOne() bool {

	readFile, err := os.Open(asm.asmFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	asm.adrMemPointer = 0
	asm.lineCount = 0
	okPassOne := true

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {

		someString := fileScanner.Text()
		asm.lineCount++
		words := strings.Fields(someString)

		var statement []string

		for w := range words {
			if strings.HasPrefix(words[w], ";") {
				break
			} else {
				statement = append(statement, strings.ToUpper(words[w]))
			}
		}
		statementLength := len(statement)

		if statementLength == 1 {
			_, ok := asm.opCodeMap[statement[0]]
			if ok {
				asm.adrMemPointer++
			} else {
				okPassOne = false
				asm.outputErrorMsg(statement, "Invalid statement. line ")
			}

		} else if statementLength >= 2 && statement[1] == "DSTRING" {

			strconst := strings.Join(statement[2:], " ")
			asm.strConstMap[statement[0]] = []byte(strconst)
			strConstLen := len(strconst)

			if strConstLen == 0 {
				okPassOne = false
				asm.outputErrorMsg(statement, "DSTRING is empty. line  ")
			}

		} else if statementLength >= 2 && statementLength <= 5 {

			if strings.HasSuffix(statement[0], ":") {
				_, ok := asm.opCodeMap[statement[1]]
				if ok {

					labHi := statement[0] + "HI"
					adrHi := asm.adrMemPointer >> 8
					labLo := statement[0] + "LO"
					adrLo := asm.adrMemPointer & 0x00FF
					asm.labelMap[labHi] = adrHi
					asm.labelMap[labLo] = adrLo
					asm.adrMemPointer++
				} else if statement[1] == "DBYTE" && statementLength == 3 {
					byteValue, err := strconv.ParseUint(statement[2], 10, 16)
					if err != nil {
						okPassOne = false
						asm.outputErrorMsg(statement, "number conversion error. line  ")
					} else if byteValue > 255 {
						okPassOne = false
						asm.outputErrorMsg(statement, "DBYTE value > 255. line  ")
					}
					key := strings.TrimSuffix(statement[0], ":")
					asm.byteMap[key] = uint16(byteValue)

				} else {
					okPassOne = false
					asm.outputErrorMsg(statement, "Invalid statement. line  ")
				}
			} else {
				_, ok := asm.opCodeMap[statement[0]]
				if ok {
					asm.adrMemPointer++
				} else {
					okPassOne = false
					asm.outputErrorMsg(statement, "Invalid statement. line  ")
				}
			}
		} else if statementLength > 5 {
			okPassOne = false
			asm.outputErrorMsg(statement, "Invalid statement. line  ")
		}
	}
	readFile.Close()

	for key, byteArray := range asm.strConstMap {

		labHi := key + "HI"
		adrHi := asm.adrMemPointer >> 8
		labLo := key + "LO"
		adrLo := asm.adrMemPointer & 0x00FF
		asm.labelMap[labHi] = adrHi
		asm.labelMap[labLo] = adrLo

		asm.strConstAdrMap[key] = asm.adrMemPointer

		strConstLen := len(byteArray)
		strConstLen++

		asm.adrMemPointer += uint16(strConstLen)
		asm.adrMemPointer++
	}

	return okPassOne
}

func (asm *Assembler) loadFlagMap() {
	asm.flagMap["EQ"] = 16
	asm.flagMap["NE"] = 32
	asm.flagMap["LT"] = 64
	asm.flagMap["GT"] = 128
}

func (asm *Assembler) loadShiftMap() {

	asm.shiftMap["0"] = 0
	asm.shiftMap["1"] = 2
	asm.shiftMap["2"] = 4
	asm.shiftMap["3"] = 6
	asm.shiftMap["4"] = 8
	asm.shiftMap["5"] = 10
	asm.shiftMap["6"] = 12
	asm.shiftMap["7"] = 14
	asm.shiftMap["8"] = 16
	asm.shiftMap["9"] = 18
	asm.shiftMap["10"] = 20
	asm.shiftMap["11"] = 22
	asm.shiftMap["12"] = 24
	asm.shiftMap["13"] = 26
	asm.shiftMap["14"] = 28
	asm.shiftMap["15"] = 30
}

func (asm *Assembler) loadRegThreeMap() {

	asm.regThreeMap["R0"] = 0
	asm.regThreeMap["R1"] = 4
	asm.regThreeMap["R2"] = 8
	asm.regThreeMap["R3"] = 12
	asm.regThreeMap["R4"] = 16
	asm.regThreeMap["R5"] = 20
	asm.regThreeMap["R6"] = 24
	asm.regThreeMap["R7"] = 28

	asm.regThreeMap["ADD"] = 2048
	asm.regThreeMap["SUB"] = 4096
	asm.regThreeMap["AND"] = 6144
	asm.regThreeMap["XOR"] = 8192
	asm.regThreeMap["OR"] = 10240
	asm.regThreeMap["LDW"] = 12288
	asm.regThreeMap["STW"] = 14336

}

func (asm *Assembler) loadRegTwoMap() {

	asm.regTwoMap["R0"] = 0
	asm.regTwoMap["R1"] = 32
	asm.regTwoMap["R2"] = 64
	asm.regTwoMap["R3"] = 96
	asm.regTwoMap["R4"] = 128
	asm.regTwoMap["R5"] = 160
	asm.regTwoMap["R6"] = 192
	asm.regTwoMap["R7"] = 224

	asm.regTwoMap["SHL"] = 16384
	asm.regTwoMap["SHR"] = 18432
	asm.regTwoMap["NOT"] = 20480
	asm.regTwoMap["CMP"] = 22528

}

func (asm *Assembler) loadRegOneMap() {

	asm.regOneMap["R0"] = 0
	asm.regOneMap["R1"] = 256
	asm.regOneMap["R2"] = 512
	asm.regOneMap["R3"] = 768
	asm.regOneMap["R4"] = 1024
	asm.regOneMap["R5"] = 1280
	asm.regOneMap["R6"] = 1536
	asm.regOneMap["R7"] = 1792

	asm.regOneMap["BFLAG"] = 24576
	asm.regOneMap["ADDI"] = 26624
	asm.regOneMap["SUBI"] = 28672
	asm.regOneMap["ANDI"] = 30720
	asm.regOneMap["XORI"] = 32768
	asm.regOneMap["ORI"] = 34816
	asm.regOneMap["LDI"] = 36864
	asm.regOneMap["CMPI"] = 38912
	asm.regOneMap["JUMP"] = 40960
	asm.regOneMap["PUSH"] = 43008
	asm.regOneMap["POP"] = 45056
	asm.regOneMap["OUT"] = 47104
	asm.regOneMap["IN"] = 49152
	asm.regOneMap["LDSR"] = 51200
	asm.regOneMap["CALL"] = 55296
	asm.regOneMap["MVSR"] = 59392
}

func (asm *Assembler) loadOpcClassMap() {
	asm.opcClassMap["HALT"] = "0args"
	asm.opcClassMap["ADD"] = "3args"
	asm.opcClassMap["SUB"] = "3args"
	asm.opcClassMap["AND"] = "3args"
	asm.opcClassMap["XOR"] = "3args"
	asm.opcClassMap["OR"] = "3args"
	asm.opcClassMap["LDW"] = "3args"
	asm.opcClassMap["STW"] = "3args"
	asm.opcClassMap["SHL"] = "3args"
	asm.opcClassMap["SHR"] = "3args"
	asm.opcClassMap["NOT"] = "2args"
	asm.opcClassMap["CMP"] = "2args"
	asm.opcClassMap["BFLAG"] = "2args"
	asm.opcClassMap["ADDI"] = "2args"
	asm.opcClassMap["SUBI"] = "2args"
	asm.opcClassMap["ANDI"] = "2args"
	asm.opcClassMap["XORI"] = "2args"
	asm.opcClassMap["ORI"] = "2args"
	asm.opcClassMap["LDI"] = "2args"
	asm.opcClassMap["CMPI"] = "2args"
	asm.opcClassMap["JUMP"] = "1args"
	asm.opcClassMap["PUSH"] = "1args"
	asm.opcClassMap["POP"] = "1args"
	asm.opcClassMap["OUT"] = "1args"
	asm.opcClassMap["IN"] = "1args"
	asm.opcClassMap["LDSR"] = "1args"
	asm.opcClassMap["MVSR"] = "1args"
	asm.opcClassMap["NOOP"] = "0args"
	asm.opcClassMap["CALL"] = "1args"
	asm.opcClassMap["RTRN"] = "0args"
	asm.opcClassMap["DBYTE"] = "skip"
	asm.opcClassMap["DSTRING"] = "skip"

}

func (asm *Assembler) loadOpCodeMap() {

	asm.opCodeMap["HALT"] = 0
	asm.opCodeMap["ADD"] = 2048
	asm.opCodeMap["SUB"] = 4096
	asm.opCodeMap["AND"] = 6144
	asm.opCodeMap["XOR"] = 8192
	asm.opCodeMap["OR"] = 10240
	asm.opCodeMap["LDW"] = 12288
	asm.opCodeMap["STW"] = 14336
	asm.opCodeMap["SHL"] = 16384
	asm.opCodeMap["SHR"] = 18432
	asm.opCodeMap["NOT"] = 20480
	asm.opCodeMap["CMP"] = 22528
	asm.opCodeMap["BFLAG"] = 24576
	asm.opCodeMap["ADDI"] = 26624
	asm.opCodeMap["SUBI"] = 28672
	asm.opCodeMap["ANDI"] = 30720
	asm.opCodeMap["XORI"] = 32768
	asm.opCodeMap["ORI"] = 34816
	asm.opCodeMap["LDI"] = 36864
	asm.opCodeMap["CMPI"] = 38912
	asm.opCodeMap["JUMP"] = 40960
	asm.opCodeMap["PUSH"] = 43008
	asm.opCodeMap["POP"] = 45056
	asm.opCodeMap["OUT"] = 47104
	asm.opCodeMap["IN"] = 49152
	asm.opCodeMap["LDSR"] = 51200
	asm.opCodeMap["NOOP"] = 53248
	asm.opCodeMap["CALL"] = 55296
	asm.opCodeMap["RTRN"] = 57344
	asm.opCodeMap["MVSR"] = 59392

}
