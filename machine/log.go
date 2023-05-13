package machine

import "fmt"

type Logger interface {
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

func DefaultLogger() Logger {
	return &defaultLogger{}
}

type defaultLogger struct{}

func (dl *defaultLogger) Errorf(template string, args ...interface{}) {
	fmt.Printf("[ Error ]     "+template, args)
}

func (dl *defaultLogger) Infof(template string, args ...interface{}) {
	fmt.Printf("[ Info ]     "+template, args)
}

func (dl *defaultLogger) Warnf(template string, args ...interface{}) {
	fmt.Printf("[ Warning! ]     "+template, args)
}

func (dl *defaultLogger) Fatalf(template string, args ...interface{}) {
	f := fmt.Sprintf("[ Fatal! ]     "+template, args)
	panic(f)
}

func (m *Machine) WithLogger(logger Logger) {
	m.logger = logger
}
