package ylog

import "bytes"
type TextFormatter struct {
}

func (s *TextFormatter) Build(buffer *bytes.Buffer, key string, value string) *bytes.Buffer {

	buffer.WriteString(key)
	buffer.WriteString("=")
	buffer.WriteString(value)
	return buffer
}

func (s *TextFormatter) NewSerial(buffer *bytes.Buffer, entry *Entry) ([]byte, error) {
	for key, value := range entry.Data {
		s.Build(buffer, key, value+" ")
	}
	s.Build(buffer, "msg", entry.msg+"\n")

	return buffer.Bytes(), nil
}
