package logger

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

var once = sync.Once{}

type Config struct {
	Filename       string `koanf:"filename"`
	LocalTime      bool   `koanf:"local_time"`
	MaxSize        int    `koanf:"max_size"`
	MaxBackups     int    `koanf:"max_backups"`
	MaxAge         int    `koanf:"max_age"`
	StdoutLogLevel string `koanf:"stdout_log_level"`
	WriterLogLevel string `koanf:"writer_log_level"`
}

func Start(cfg Config) {
	once.Do(func() {
		Logger = zap.Must(zap.NewProduction())

		PEConfig := zap.NewProductionEncoderConfig()
		PEConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		defaultEncoder := zapcore.NewJSONEncoder(PEConfig)

		fmt.Printf("%+v\n", cfg)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Filename,
			LocalTime:  cfg.LocalTime,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
		})

		stdOutWriter := zapcore.AddSync(os.Stdout)

		writerLogLevel := zapcore.InfoLevel
		stdOutLogLevel := writerLogLevel

		switch cfg.StdoutLogLevel {
		case "debug":
			stdOutLogLevel = zapcore.DebugLevel
		case "error":
			stdOutLogLevel = zapcore.ErrorLevel
		case "info":
			stdOutLogLevel = zapcore.InfoLevel
		case "fatal":
			stdOutLogLevel = zapcore.FatalLevel
		case "warning":
			stdOutLogLevel = zapcore.WarnLevel
		}

		switch cfg.WriterLogLevel {
		case "debug":
			writerLogLevel = zapcore.DebugLevel
		case "error":
			writerLogLevel = zapcore.ErrorLevel
		case "info":
			writerLogLevel = zapcore.InfoLevel
		case "fatal":
			writerLogLevel = zapcore.FatalLevel
		case "warning":
			writerLogLevel = zapcore.WarnLevel
		}

		core := zapcore.NewTee(
			zapcore.NewCore(defaultEncoder, stdOutWriter, stdOutLogLevel),
			zapcore.NewCore(defaultEncoder, writer, writerLogLevel),
		)
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})
}
