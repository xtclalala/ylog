package ylog

type HookLevel map[LogLevel][]HookFn

type HookFn func(entry *Entry)

type Hook func(entry *Entry) error

func AddHook(fn Hook) {
	defaultLogger.AddHook(fn)
}

func AddHookLevel(level LogLevel, fns []HookFn) {
	defaultLogger.AddHookLevel(level, fns)
}
