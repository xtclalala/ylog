package ylog

type Entry struct {
	Data map[string]string
	msg  string
	err  error
}

func (s *Entry) WithFields(f map[string]string) {
	var data = s.Data
	var nData = make(map[string]string, len(data)+len(f))
	for k, v := range data {
		nData[k] = v
	}
	for k, v := range f {
		nData[k] = v
	}
	s.Data = nData
}

func (s *Entry) WithField(key, value string) {
	var f = make(map[string]string, 1)
	f[key] = value
	s.WithFields(f)
}

func (s *Entry) Clear() {
	s.Data = nil
	s.msg = ""
	s.err = nil
}

func WithField(key, value string) {
	defaultLogger.WithField(key, value)
}

func WithFields(f map[string]string) {
	defaultLogger.WithFields(f)
}
