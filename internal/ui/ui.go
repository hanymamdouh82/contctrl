// internal/ui/ui.go
package ui

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	confirm = color.New(color.FgGreen)
	info    = color.New(color.FgCyan)
	fatal   = color.New(color.FgRed, color.Bold)
	dim     = color.New(color.Faint)
)

func Confirm(label, value string) {
	fmt.Printf("\r")
	confirm.Printf("✔ ")
	fmt.Printf("%s: ", label)
	info.Printf("%s\n", value)
}

func Section(label string) {
	fmt.Println()
	dim.Printf("─── %s ", label)
	fmt.Println()
}

func Fatal(msg string) {
	fatal.Printf("✖ %s\n", msg)
}
