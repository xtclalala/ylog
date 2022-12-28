package ylog
type HookLevel map[LogLevel][]HookFn

type HookFn func(entry *Entry)

type Hook func(entry *Entry) error
