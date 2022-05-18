package logger

import (
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Module = fx.Provide(NewLogger)

const (
	LoggerTimeKey    = "time"
	LoggerTimeFormat = "2006-01-02 15:04:05"
)

func FileLogHook(vp *viper.Viper) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   vp.GetString("LOG.PATH"),
		MaxSize:    vp.GetInt("LOG.MAXSIZE"),
		MaxAge:     vp.GetInt("LOG.MAXAGE"),
		MaxBackups: vp.GetInt("LOG.MAXBACKUPS"),
		Compress:   vp.GetBool("LOG.COMPRESS"),
	}
}

// Load Encoder Config
func NewProductionEncoderConfig() zapcore.EncoderConfig {
	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.TimeKey = LoggerTimeKey
	EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format(LoggerTimeFormat))
	}
	return EncoderConfig
}

func NewLogger(vp *viper.Viper) *zap.Logger {

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// cores: Maybe Add Kafka Log Hook, cores shuold be slice
	var cores []zapcore.Core

	if vp.GetBool("LOG.DEBUG") {
		// Development Stdout
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		consoleDebugging := zapcore.Lock(os.Stdout)
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	} else {
		// Other (Test、Production) LogFile
		fileEncoder := zapcore.NewJSONEncoder(NewProductionEncoderConfig())
		writerHook := zapcore.AddSync(FileLogHook(vp))
		cores = append(cores, zapcore.NewCore(fileEncoder, writerHook, highPriority))
	}

	return zap.New(zapcore.NewTee(cores...)).WithOptions(zap.AddCaller())
}
