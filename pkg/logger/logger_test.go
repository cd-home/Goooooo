package logger

import (
	"testing"
)

func TestZapLogger(t *testing.T) {
	cfg := &FileLogConfig{
		Level:       "DEBUG",
		Filepath:    "./log.log",
		FileMaxSize: 5000,
		FileMaxAge:  5000,
		MaxBackups:  10,
		Compress:    true,
	}
	logger := NewEmptyFileLogger(cfg)
	logger.Build().Info("Write")
	logger.WithOptions(WithEnable(false)).Build().Info("False Write")
}

func BenchmarkZapLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
