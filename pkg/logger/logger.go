package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	enabled bool
	Log     *zap.Logger
	cfg     *FileLogConfig
	high    zap.LevelEnablerFunc
	low     zap.LevelEnablerFunc
}

func (s *Logger) WithOptions(opts ...option) *Logger {
	c := s.clone()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

func WithEnable(enabled bool) option {
	return optionFunc(func(s *Logger) {
		s.enabled = enabled
	})
}

func (s *Logger) clone() *Logger {
	copy := *s
	return &copy
}

// Build zap.Logger
func (log *Logger) Build() *Logger {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileWriter := zapcore.AddSync(FileLogHook(log.cfg))
	var priority zap.LevelEnablerFunc
	if log.cfg.Level == DebugLevel {
		priority = log.low
	} else {
		priority = log.high
	}
	core := zapcore.NewTee(zapcore.NewCore(consoleEncoder, fileWriter, priority))
	log.Log = zap.New(core).WithOptions(zap.AddCaller())
	return log
}

func (log *Logger) Info(msg string, fields ...zap.Field) {
	if !log.enabled {
		return
	}
	if ce := log.Log.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *Logger) Debug(msg string, fields ...zap.Field) {
	if !log.enabled {
		return
	}
	if ce := log.Log.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *Logger) Warn(msg string, fields ...zap.Field) {
	if !log.enabled {
		return
	}
	if ce := log.Log.Check(zap.WarnLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *Logger) Error(msg string, fields ...zap.Field) {
	if !log.enabled {
		return
	}
	if ce := log.Log.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *Logger) DPanic(msg string, fields ...zap.Field) {
	if !log.enabled {
		return
	}
	if ce := log.Log.Check(zap.DPanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *Logger) Panic(msg string, fields ...zap.Field) {
	if !log.enabled {
		return
	}
	if ce := log.Log.Check(zap.PanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (log *Logger) Fatal(msg string, fields ...zap.Field) {
	if !log.enabled {
		return
	}
	if ce := log.Log.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}
