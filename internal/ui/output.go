package ui

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Success prints a success message in bold green.
func Success(msg string) {
	fmt.Printf("%s %s\n", ansi.Color("SUCCESS:", "green+b"), msg)
}

// Warning prints a warning message in bold yellow.
func Warning(msg string) {
	fmt.Printf("%s %s\n", ansi.Color("WARNING:", "yellow+b"), msg)
}

// Error prints an error message in bold red.
func Error(msg string) {
	fmt.Printf("%s %s\n", ansi.Color("ERROR:", "red+b"), msg)
}

// Info prints an info message in bold blue.
func Info(msg string) {
	fmt.Printf("%s %s\n", ansi.Color("INFO:", "blue+b"), msg)
}

// Analyzing prints a special message for the "check" command.
func Analyzing(architecture string) {
	fmt.Printf("🔍 %s **%s**...\n\n", ansi.Color("Analizando arquitectura", "cyan+b"), architecture)
}

// Fatal prints an error message and exits the program.
func Fatal(err error) {
	fmt.Printf("%s %v\n", ansi.Color("FATAL:", "red+b"), err)
	os.Exit(1)
}

// SuccessMsg returns a success message in bold green.
func SuccessMsg(msg string) string {
	return fmt.Sprintf("%s %s", ansi.Color("SUCCESS:", "green+b"), msg)
}

// WarningMsg returns a warning message in bold yellow.
func WarningMsg(msg string) string {
	return fmt.Sprintf("%s %s", ansi.Color("WARNING:", "yellow+b"), msg)
}

// ErrorMsg returns an error message in bold red.
func ErrorMsg(msg string) string {
	return fmt.Sprintf("%s %s", ansi.Color("ERROR:", "red+b"), msg)
}

// InfoMsg returns an info message in bold blue.
func InfoMsg(msg string) string {
	return fmt.Sprintf("%s %s", ansi.Color("INFO:", "blue+b"), msg)
}
