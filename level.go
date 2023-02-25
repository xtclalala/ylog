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
		return "info"
	case 2:
		return "warning"
	case 4:
		return "error"
	case 8:
		return "fatal"
	case 16:
		return "panic"
	case 32:
		return "debug"
	default:
		return "debug"
	}
}

func SetLogLevel(level LogLevel) {
	defaultLogger.SetLogLevel(level)
}
