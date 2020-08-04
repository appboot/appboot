package logger

import (
	"fmt"
	"log"
)

// Color list
const (
	Red = uint8(iota + 31)
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// LogE log error
func LogE(v ...interface{}) {
	LogWrap(Red, v...)
}

// LogW log warning
func LogW(v ...interface{}) {
	LogWrap(Yellow, v...)
}

// LogI log info
func LogI(v ...interface{}) {
	LogWrap(Blue, v...)
}

// LogV log verbose
func LogV(v ...interface{}) {
	LogWrap(White, v...)
}

// LogWrap log with color and wrap
func LogWrap(color uint8, v ...interface{}) {
	Log(color, v, "\n")
}

// Log log with color
func Log(color uint8, v ...interface{}) {
	tmp := fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, v)
	if len(tmp) > 0 {
		log.Print(tmp)
	}
}
