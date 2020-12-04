package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

func NewLogger(filename,levels string) *zap.Logger{
	hook := lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10240,
		MaxAge:     30,
		MaxBackups: 7,
	}

	timeFormat := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "line_num",
		StacktraceKey:  "stack_trace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeFormat,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var level zapcore.Level
	switch strings.ToUpper(levels) {
	case "DEBUG":
		level = zap.DebugLevel
	case "WARN":
		level = zap.WarnLevel
	case "INFO":
		level = zap.InfoLevel
	case "ERROR":
		level = zap.ErrorLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	)

	caller := zap.AddCaller()
	development := zap.Development()
	return zap.New(core, caller, development)
}
