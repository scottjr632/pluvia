package logging

type noopLogger struct{}

func (l *noopLogger) Info(msg string)  {}
func (l *noopLogger) Error(msg string) {}

func NewNoopLogger() *noopLogger {
	return &noopLogger{}
}

var _ Logger = (*noopLogger)(nil)
