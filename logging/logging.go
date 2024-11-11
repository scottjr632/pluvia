package logging

type Logger interface {
	Info(msg ...interface{})
	Debug(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
}
