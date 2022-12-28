package ylog

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger struct {

	// The logger output modes. example: file or os
	Outs []io.Writer

	// Used to Out before
	HookLevel HookLevel
	Hooks     []Hook

	// The log level the logger should log it.
	Level LogLevel

	// Used to sync writing to the log.
	mx *sync.Mutex

	// All log Formatter before logged to Out.The included
	// formatters are `JsonFormatter` and `TextFormatter` for
	// which TextFormatter is default.
	formatter Formatter

	// The Buffer Pool used to format the log.
	bup *BufferPool

	Ctx context.Context

	enp *EntryPool

	ExitFn ExitFn

	entry *Entry
}

type ExitFn func(int)

func New() *Logger {
	return &Logger{
		Outs:      []io.Writer{os.Stderr},
		Level:     InfoLevel,
		mx:        &sync.Mutex{},
		Ctx:       context.Background(),
		enp:       newEnP(),
		bup:       defaultBufferPool,
		ExitFn:    os.Exit,
		formatter: new(TextFormatter),
		// todo
	}
}

func (s *Logger) getEntry() *Entry {
	var e = s.enp.Get()
	return e
}

func (s *Logger) putEntry(entry *Entry) {
	s.enp.Put(entry)
}

func (s *Logger) WithField(key string, value string) *Logger {
	var f = make(map[string]string, 1)
	f[key] = value
	return s.WithFields(f)
}

func (s *Logger) WithFields(fields map[string]string) *Logger {
	var e = s.getEntry()
	e.Data = fields
	return &Logger{
		Outs:      s.Outs,
		HookLevel: s.HookLevel,
		Hooks:     s.Hooks,
		Level:     s.Level,
		mx:        s.mx,
		formatter: s.formatter,
		bup:       s.bup,
		Ctx:       s.Ctx,
		enp:       s.enp,
		ExitFn:    s.ExitFn,
		entry:     e,
	}
}

func (s *Logger) log(level LogLevel, msg string) {
	var (
		err error
		e   *Entry
	)
	e = s.entry
	if e == nil {
		e = s.getEntry()
	}

	for _, hook := range s.Hooks {
		err = hook(e)
		if err != nil {
			msg += err.Error()
		}
	}

	for _, levelFn := range s.HookLevel[level] {
		levelFn(e)
	}
	var (
		b   []byte
		buf *bytes.Buffer
	)
	e.msg = msg
	buf = s.bup.Get()
	b, err = s.formatter.NewSerial(buf, e)
	buf.Reset()
	s.bup.Put(buf)
	s.putEntry(e)
	s.entry = nil
	for _, out := range s.Outs {
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		_, err = out.Write(b)
	}

	if level == PanicLevel || level == FatalLevel {
		s.ExitFn(1)
	}
}

func (s *Logger) Debug(msg string) {
	s.log(DebugLevel, msg)
}
func (s *Logger) Debugf(format string, args ...any) {
	s.log(DebugLevel, fmt.Sprintf(format, args...))
}

func (s *Logger) Info(msg string) {
	s.log(InfoLevel, msg)
}
func (s *Logger) Infof(format string, args ...any) {
	s.log(InfoLevel, fmt.Sprintf(format, args...))
}

func (s *Logger) Warn(msg string) {
	s.log(WarnLevel, msg)
}
func (s *Logger) Warnf(format string, args ...any) {
	s.log(WarnLevel, fmt.Sprintf(format, args...))
}

func (s *Logger) Error(msg string) {
	s.log(ErrorLevel, msg)
}
func (s *Logger) Errorf(format string, args ...any) {
	s.log(ErrorLevel, fmt.Sprintf(format, args...))
}

func (s *Logger) Fatal(msg string) {
	s.log(FatalLevel, msg)
}
func (s *Logger) Fatalf(format string, args ...any) {
	s.log(FatalLevel, fmt.Sprintf(format, args...))
}

func (s *Logger) Panic(msg string) {
	s.log(PanicLevel, msg)
}
func (s *Logger) Panicf(format string, args ...any) {
	s.log(PanicLevel, fmt.Sprintf(format, args...))
}

var defaultLogger *Logger

func init() {
	defaultLogger = New()
}

func Debug(msg string) {
	log(DebugLevel, msg)
}
func Debugf(format string, args ...any) {
	log(DebugLevel, fmt.Sprintf(format, args...))
}

func Info(msg string) {
	log(InfoLevel, msg)
}
func Infof(format string, args ...any) {
	log(InfoLevel, fmt.Sprintf(format, args...))
}

func Warn(msg string) {
	log(WarnLevel, msg)
}
func Warnf(format string, args ...any) {
	log(WarnLevel, fmt.Sprintf(format, args...))
}

func Error(msg string) {
	log(ErrorLevel, msg)
}
func Errorf(format string, args ...any) {
	log(ErrorLevel, fmt.Sprintf(format, args...))
}

func Fatal(msg string) {
	log(FatalLevel, msg)
}
func Fatalf(format string, args ...any) {
	log(FatalLevel, fmt.Sprintf(format, args...))
}

func Panic(msg string) {
	log(PanicLevel, msg)
}
func Panicf(format string, args ...any) {
	log(PanicLevel, fmt.Sprintf(format, args...))
}

func log(level LogLevel, msg string) {
	defaultLogger.log(level, msg)
}

func Demo() {
	l := New()
	l.WithField("key1", "value").Info("this is testing")
	l.WithFields(map[string]string{"key2": "value"}).Info("this is testing")
	l.WithField("key3", "value").Warn("this is testing")
	l.WithField("key4", "value").Fatal("this is testing")
}
