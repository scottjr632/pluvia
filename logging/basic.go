package logging

import (
	"github.com/fatih/color"
)

const (
	debugPrefix = "[DEBUG]: "
	infoPrefix  = "[INFO]: "
	warnPrefix  = "[WARN]: "
	errorPrefix = "[ERROR]: "
)

var (
	debugColor = color.New(color.FgWhite)
	infoColor  = color.New(color.FgHiWhite)
	warnColor  = color.New(color.FgYellow)
	errorColor = color.New(color.FgRed)
)

type BasicLogger struct{}

func NewBasicLogger() *BasicLogger {
	return &BasicLogger{}
}

func (l *BasicLogger) Info(msg ...interface{}) {
	infoColor.Println(append([]interface{}{infoPrefix}, msg...)...)
	println()
}

func (l *BasicLogger) Error(msg ...interface{}) {
	errorColor.Println(append([]interface{}{errorPrefix}, msg...)...)
	println()
}

func (l *BasicLogger) Debug(msg ...interface{}) {
	debugColor.Print(append([]interface{}{debugPrefix}, msg...)...)
	println()
}

func (l *BasicLogger) Warn(msg ...interface{}) {
	warnColor.Println(append([]interface{}{warnPrefix}, msg...))
	println()
}

var _ Logger = (*BasicLogger)(nil)
