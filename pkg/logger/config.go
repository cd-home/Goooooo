package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FileLogConfig struct {
	Level       string `json:"level"`
	Filepath    string `json:"filepath"`
	FileMaxSize int    `json:"fileMaxSize"`
	FileMaxAge  int    `json:"fileMaxAge"`
	MaxBackups  int    `json:"maxBackups"`
	Compress    bool   `json:"compress"`
}

type kafkaLogConfig struct {
	Addr        []string `json:"addr"`
	Enable      bool     `json:"enable"`
	Topics      string   `json:"topics"`
	RequiredAck int      `json:"requiredack"`
	Mode        string   `json:"mode"`
}

func NewEmptyFileLogger(cfg *FileLogConfig) *Logger {
	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})
	logger := &Logger{
		enabled: true,
		high:    highPriority,
		low:     lowPriority,
		cfg:     cfg,
	}
	// Default Open
	logger.WithOptions(WithEnable(true))
	return logger
}
