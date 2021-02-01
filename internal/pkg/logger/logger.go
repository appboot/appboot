package logger

import (
	"github.com/fatih/color"
)

// LogE log error
func LogE(format string, a ...interface{}) {
	color.Red(format, a...)
}

// LogW log warning
func LogW(format string, a ...interface{}) {
	color.Yellow(format, a...)
}

// LogH log hint
func LogH(format string, a ...interface{}) {
	color.Green(format, a...)
}

// LogI log info
func LogI(format string, a ...interface{}) {
	color.Blue(format, a...)
}

// LogV log verbose
func LogV(format string, a ...interface{}) {
	color.White(format, a...)
}
