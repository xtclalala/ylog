package ylog

type LogLevel uint16

const (
	DebugLevel LogLevel = 1 << iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)