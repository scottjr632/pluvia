package logging

type noopLogger struct{}

func (l *noopLogger) Info(msg ...interface{})  {}
func (l *noopLogger) Error(msg ...interface{}) {}
func (l *noopLogger) Debug(msg ...interface{}) {}
func (l *noopLogger) Warn(msg ...interface{})  {}

func NewNoopLogger() *noopLogger {
	return &noopLogger{}
}

var _ Logger = (*noopLogger)(nil)
