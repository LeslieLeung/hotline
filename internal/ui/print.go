package ui

import (
	"github.com/gookit/color"
	"os"
)

var (
	Debug bool // Enable debug output
)

func Errorf(format string, a ...interface{}) {
	color.Danger.Printf(format+"\n", a...)
}

func ErrorfExit(format string, a ...interface{}) {
	color.Danger.Printf(format+"\n", a...)
	os.Exit(1)
}

func Printf(format string, a ...interface{}) {
	color.Success.Printf(format+"\n", a...)
}

func Debugf(format string, a ...interface{}) {
	if Debug {
		color.Info.Printf("[DEBUG]"+format+"\n", a...)
	}
}
