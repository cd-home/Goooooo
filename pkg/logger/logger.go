package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type FileLogConfig struct {
	Debug       bool   `json:"debug"`
	FilePath    string `json:"filePath"`
	FileMaxSize int    `json:"fileMaxSize"`
	FileMaxAge  int    `json:"fileMaxAge"`
	MaxBackups  int    `json:"maxBackups"`
	Compress    bool   `json:"compress"`
}

func FileLogHook(cfg *FileLogConfig) *lumberjack.Logger {
	hook := &lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.FileMaxSize,
		MaxAge:     cfg.FileMaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
	}
	return hook
}

// Load Config
func NewProductionEncoderConfig() zapcore.EncoderConfig {
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.TimeKey = "time"
	EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}
	return EncoderConfig
}

func New(cfg *FileLogConfig) *zap.Logger {

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	var cores []zapcore.Core

	if cfg.Debug {
		// Development
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleDebugging := zapcore.Lock(os.Stdout)
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	} else {
		// Production
		fileEncoder := zapcore.NewJSONEncoder(NewProductionEncoderConfig())
		writerHook := zapcore.AddSync(FileLogHook(cfg))
		cores = append(cores, zapcore.NewCore(fileEncoder, writerHook, highPriority))
	}

	return zap.New(zapcore.NewTee(cores...)).WithOptions(zap.AddCaller())
}
