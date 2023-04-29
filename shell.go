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

func (core *Config) clearIO() {
	fmt.Print(cursor.MoveTo(37, 108))
	fmt.Println("   ")
	fmt.Print(cursor.MoveTo(41, 108))
	fmt.Println("   ")
}

func (core *Config) cursorHome() {
	fmt.Print(cursor.MoveTo(22, 137))
}

func (core *Config) setOutputHome() {
	core.OutRow = 37
	core.OutCol = 108
}

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
	fmt.Print(cursor.MoveTo(13, 135))
	fmt.Println("[U,D] - toggle delay")
	fmt.Print(cursor.MoveTo(14, 135))
	fmt.Println("R - reset core")

	fmt.Print(cursor.MoveTo(15, 135))
	fmt.Println("A - assemble ./source.asm")

	fmt.Print(cursor.MoveTo(16, 135))
	fmt.Println("L - load ./source.hex")

	fmt.Print(cursor.MoveTo(17, 135))
	fmt.Println("I - step one Instruction")

	fmt.Print(cursor.MoveTo(18, 135))
	fmt.Println("S - step one cycle")

	fmt.Print(cursor.MoveTo(19, 135))
	fmt.Println("G - go run program")

	fmt.Print(cursor.MoveTo(20, 135))
	fmt.Println("B - break at NOOP for debug. S or G to resume.")

	fmt.Print(cursor.MoveTo(21, 135))
	fmt.Println("H - halt")

	fmt.Print(cursor.MoveTo(22, 135))
	fmt.Println("$ ")
	fmt.Print(cursor.MoveTo(22, 137))
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
			fmt.Print(cursor.MoveTo(22, 137))
			input := strings.ToUpper(string(r))

			fmt.Println(input)
			fmt.Print(cursor.MoveTo(22, 137))
			fmt.Print(cursor.Show())

			if input == "H" {
				break
			} else if input == "U" && core.SleepTime < 10 {
				core.SleepTime += 1
			} else if input == "D" && core.SleepTime > 0 {
				core.SleepTime -= 1
			} else if input == "S" {
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
				core.resetCore()
				core.OperationClass = "fetch"
				core.CoreMemPoint = 0
				tickCount = 1
			} else if input == "I" {
				core.runAll(input)
			} else if input == "G" {
				core.runAll(input)
			} else if input == "B" {
				core.runAll(input)
			} else if input == "A" {
				core.ASM.Assemble(core.CoreMemory[:], "assemble")
			} else if input == "L" {
				core.ASM.Assemble(core.CoreMemory[:], "loadCore")
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
