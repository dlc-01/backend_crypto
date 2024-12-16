package logger

type ILogger interface {
	Fatalf(format string, opts ...any)
	Errorf(format string, opts ...any)
	Infof(format string, opts ...any)
	Warnf(format string, opts ...any)
	Panicf(format string, opts ...any)
	Info(msg string)
}
