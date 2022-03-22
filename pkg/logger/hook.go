package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

func FileLogHook(cfg *FileLogConfig) *lumberjack.Logger {
	hook := &lumberjack.Logger{
		Filename:   cfg.Filepath,
		MaxSize:    cfg.FileMaxSize,
		MaxAge:     cfg.FileMaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   cfg.Compress,
	}
	return hook
}

// TODO KafKaLog
type KafKaLog struct{}

func KafkaLogHook(cfg *kafkaLogConfig) *KafKaLog {
	return &KafKaLog{}
}

func (k *KafKaLog) Write(data []byte) (n int, err error) {
	return
}
