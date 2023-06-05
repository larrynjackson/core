package main

import (
	"fmt"

	"github.com/ahmetalpbalkan/go-cursor"
)

func (core *Config) printDoc(address int) {
	fmt.Print(cursor.Hide())

	fmt.Print(cursor.MoveTo(46, 140))
	fmt.Print("                                                                                 ")

	fmt.Print(cursor.MoveTo(46, 140))
	fmt.Printf(core.DebugMap[address])

	core.cursorHome()
	fmt.Print(cursor.Show())

}
