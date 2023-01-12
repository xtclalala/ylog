package ylog

type LogLevel uint16

const (
	InfoLevel LogLevel = 1 << iota
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	DebugLevel
)

func (s LogLevel) String() string {
	switch s {
	case 1:
		return "debug"
	case 2:
		return "info"
	case 4:
		return "warning"
	case 8:
		return "error"
	case 16:
		return "fatal"
	case 32:
		return "panic"
	default:
		return "debug"
	}
}

func SetLogLevel(level LogLevel) {
	defaultLogger.SetLogLevel(level)
}
