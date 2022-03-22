package logger

const (
	DebugLevel = "DEBUG"
	InfoLevel  = "INFO"
	WarnLevel  = "WARNING"
	ErrorLevel = "ERROR"
)

type option interface {
	apply(*Logger)
}

type optionFunc func(*Logger)

func (f optionFunc) apply(s *Logger) {
	f(s)
}
