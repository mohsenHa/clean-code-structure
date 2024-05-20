package logger

import (
	"clean-code-structure/config"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

var Logger *zap.Logger

var once = sync.Once{}

type Config struct {
	Filename   string `json:"filename"`
	LocalTime  bool   `json:"local_time"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
}

func init() {
	once.Do(func() {
		cfg := config.Load("config.yml")
		fmt.Printf("cfg: %+v\n", cfg)

		Logger, _ = zap.NewProduction()

		PEConfig := zap.NewProductionEncoderConfig()
		PEConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		defaultEncoder := zapcore.NewJSONEncoder(PEConfig)

		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Logger.Filename,
			LocalTime:  cfg.Logger.LocalTime,
			MaxSize:    cfg.Logger.MaxSize,
			MaxBackups: cfg.Logger.MaxBackups,
			MaxAge:     cfg.Logger.MaxAge,
		})

		stdOutWriter := zapcore.AddSync(os.Stdout)
		defaultLogLevel := zapcore.InfoLevel
		core := zapcore.NewTee(
			zapcore.NewCore(defaultEncoder, writer, defaultLogLevel),
			zapcore.NewCore(defaultEncoder, stdOutWriter, zap.InfoLevel),
		)
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})
}
