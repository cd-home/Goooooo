package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		Name string
		Cfg *FileLogConfig
	}{
		{
			Name: "console",
			Cfg: &FileLogConfig{
				Debug:       true,
				FilePath:    "./testdata/log.log",
				FileMaxSize: 500,
				FileMaxAge:  30,
				MaxBackups:  10,
				Compress:    true,
			},
		},
		{
			Name: "LogFile",
			Cfg: &FileLogConfig{
				Debug:       false,
				FilePath:    "./testdata/log.log",
				FileMaxSize: 500,
				FileMaxAge:  30,
				MaxBackups:  10,
				Compress:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			logger := New(tt.Cfg)
			logger.Info("Testing", zap.String("type", tt.Name))
		})
	}
}

func BenchmarkLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
