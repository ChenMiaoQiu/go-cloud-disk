package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/ChenMiaoQiu/go-cloud-disk/conf"
)

const (
	LevelError = iota
	LevelWarning
	LevelInformational
	LevelDebug
)

// logger logger
type Logger struct {
	level int
}

var logger *Logger

// Println print log msg with time
func (ll *Logger) Println(msg string) {
	log.Println(msg)
}

// Panic print panic error
func (ll *Logger) Panic(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[Panic] "+format, v...)
	ll.Println(msg)
	os.Exit(0)
}

// Error print err
func (ll *Logger) Error(format string, v ...interface{}) {
	if LevelError > ll.level {
		return
	}
	msg := fmt.Sprintf("[Error] "+format, v...)
	ll.Println(msg)
}

// Panic print Warning
func (ll *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > ll.level {
		return
	}
	msg := fmt.Sprintf("[Warning] "+format, v...)
	ll.Println(msg)
}

// Info print tips info
func (ll *Logger) Info(format string, v ...interface{}) {
	if LevelInformational > ll.level {
		return
	}
	msg := fmt.Sprintf("[Info] "+format, v...)
	ll.Println(msg)
}

// Debug print any msg
func (ll *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > ll.level {
		return
	}
	msg := fmt.Sprintf("[Debug] "+format, v...)
	ll.Println(msg)
}

// BuildLogger build loger by level
func BuildLogger() {
	level := conf.LogLevel
	intLevel := LevelError
	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInformational
	case "debug":
		intLevel = LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	logger = &l
}

// Log return logger
func Log() *Logger {
	if logger == nil {
		l := Logger{
			level: LevelDebug,
		}
		logger = &l
	}
	return logger
}
