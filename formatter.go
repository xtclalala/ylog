package ylog

import "bytes"
type Formatter interface {
	NewSerial(*bytes.Buffer, *Entry) ([]byte, error)
	Build(*bytes.Buffer, string, string) *bytes.Buffer
}
