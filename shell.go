package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/ahmetalpbalkan/go-cursor"
	"github.com/mattn/go-tty"
)

//const homeColumn = 142
//const menuColumn = 140

// func (core *Config) clearIO() {
// 	fmt.Print(cursor.MoveTo(37, 108))
// 	fmt.Println("   ")
// 	fmt.Print(cursor.MoveTo(41, 108))
// 	fmt.Println("   ")
// }

func (core *Config) cursorHome() {
	fmt.Print(cursor.MoveTo(22, homeColumn))
}

// func (core *Config) setOutputHome() {
// 	core.OutRow = 37
// 	core.OutCol = 108
// }

func (core *Config) shell() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		fmt.Print(cursor.Show())
		fmt.Println("ctrl-c pressed!")
		close(quit)
	}()

	fmt.Print(cursor.Hide())
	fmt.Print(cursor.MoveTo(13, menuColumn))
	fmt.Println("[U,D] - TOGGLE DELAY")
	fmt.Print(cursor.MoveTo(14, menuColumn))
	fmt.Println("R - RESET CORE")

	fmt.Print(cursor.MoveTo(15, menuColumn))
	fmt.Println("A - ASSEMBLE ./source.asm")

	fmt.Print(cursor.MoveTo(16, menuColumn))
	fmt.Println("L - LOAD ./source.hex")

	fmt.Print(cursor.MoveTo(17, menuColumn))
	fmt.Println("I - STEP ONE INSTRUCTION")

	fmt.Print(cursor.MoveTo(18, menuColumn))
	fmt.Println("S - STEP ONE CYCLE")

	fmt.Print(cursor.MoveTo(19, menuColumn))
	fmt.Println("G - GO RUN PROGRAM")

	fmt.Print(cursor.MoveTo(20, menuColumn))
	fmt.Println("B - BREAK AT NOOP. S, I, G TO RESUME.")

	fmt.Print(cursor.MoveTo(21, menuColumn))
	fmt.Println("H - HALT")

	fmt.Print(cursor.MoveTo(22, menuColumn))
	fmt.Println("$ ")
	fmt.Print(cursor.MoveTo(22, homeColumn))
	fmt.Print(cursor.Show())

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	tickCount := 1

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		if r != 0 {
			fmt.Print(cursor.MoveTo(22, homeColumn))
			input := strings.ToUpper(string(r))

			fmt.Println(input)
			fmt.Print(cursor.MoveTo(22, homeColumn))
			fmt.Print(cursor.Show())

			if input == "H" {
				break
			} else if input == "U" && core.SleepTime < 10 {
				core.SleepTime += 1
			} else if input == "D" && core.SleepTime > 0 {
				core.SleepTime -= 1
			} else if input == "S" {
				core.DebugMode = true
				switch core.OperationClass {
				case "fetch":
					tickCount = core.fetchInstruction(tickCount)
				case "translate":
					tickCount = core.translateInstruction(tickCount)
				case "3Register":
					tickCount = core.threeRegister(tickCount)
				case "shift":
					tickCount = core.shift(tickCount)
				case "2Register":
					tickCount = core.twoRegister(tickCount)
				case "branch":
					tickCount = core.branch(tickCount)
				case "immediate":
					tickCount = core.immediate(tickCount)
				case "jump":
					tickCount = core.jump(tickCount)
				case "ldsr":
					tickCount = core.loadSR(tickCount)
				case "mvsr":
					tickCount = core.moveSR(tickCount)
				case "push":
					tickCount = core.push(tickCount)
				case "pop":
					tickCount = core.pop(tickCount)
				case "inout":
					tickCount = core.inOut(tickCount)
				case "NOOP":
					tickCount = core.noop()
				case "call":
					tickCount = core.call(tickCount)
				case "rtrn":
					tickCount = core.rtrn(tickCount)
				}
			} else if input == "R" {
				core.DebugMode = false
				core.clearPrintStack()
				core.resetCore()
				core.OperationClass = "fetch"
				core.CoreMemPoint = 0
				tickCount = 1
			} else if input == "I" {
				core.DebugMode = true
				core.runAll(input)
			} else if input == "G" {
				core.DebugMode = false
				core.clearPrintStack()
				core.runAll(input)
			} else if input == "B" {
				core.DebugMode = true
				core.runAll(input)
			} else if input == "A" {
				core.ASM.Assemble(homeColumn, core.CoreMemory[:], "assemble", core.DebugMap)
			} else if input == "L" {
				core.ASM.Assemble(homeColumn, core.CoreMemory[:], "loadCore", core.DebugMap)
			}
		}

	}
	fmt.Print(cursor.MoveTo(51, 0))
}

func (core *Config) runAll(key string) {

	var tickCount = 1

	for {

		switch core.OperationClass {
		case "fetch":
			tickCount = core.fetchInstruction(tickCount)
		case "translate":
			tickCount = core.translateInstruction(tickCount)
			if key == "I" {
				tickCount = 8
				return
			}
		case "3Register":
			tickCount = core.threeRegister(tickCount)
		case "shift":
			tickCount = core.shift(tickCount)
		case "2Register":
			tickCount = core.twoRegister(tickCount)
		case "branch":
			tickCount = core.branch(tickCount)
		case "immediate":
			tickCount = core.immediate(tickCount)
		case "jump":
			tickCount = core.jump(tickCount)
		case "ldsr":
			tickCount = core.loadSR(tickCount)
		case "mvsr":
			tickCount = core.moveSR(tickCount)
		case "push":
			tickCount = core.push(tickCount)
		case "pop":
			tickCount = core.pop(tickCount)
		case "inout":
			tickCount = core.inOut(tickCount)
		case "Halt":
			return
		case "NOOP":
			if key == "B" || key == "I" {
				return
			}
			tickCount = core.noop()
		case "call":
			tickCount = core.call(tickCount)
		case "rtrn":
			tickCount = core.rtrn(tickCount)
		}
	}
}
