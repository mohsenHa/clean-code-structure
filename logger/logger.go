package logger

import (
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

var once = sync.Once{}

type Config struct {
	Filename       string `json:"filename"`
	LocalTime      bool   `json:"local_time"`
	MaxSize        int    `json:"max_size"`
	MaxBackups     int    `json:"max_backups"`
	MaxAge         int    `json:"max_age"`
	StdoutLogLevel string `koanf:"stdout_log_level"`
}

func Start(cfg Config) {
	once.Do(func() {
		lg, err := zap.NewProduction()
		if err != nil {
			log.Fatal(err)
		}

		Logger = lg

		PEConfig := zap.NewProductionEncoderConfig()
		PEConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		defaultEncoder := zapcore.NewJSONEncoder(PEConfig)

		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Filename,
			LocalTime:  cfg.LocalTime,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
		})

		stdOutWriter := zapcore.AddSync(os.Stdout)
		defaultLogLevel := zapcore.InfoLevel
		stdOutLogLevel := defaultLogLevel
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
		core := zapcore.NewTee(
			zapcore.NewCore(defaultEncoder, writer, defaultLogLevel),
			zapcore.NewCore(defaultEncoder, stdOutWriter, stdOutLogLevel),
		)
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})
}
