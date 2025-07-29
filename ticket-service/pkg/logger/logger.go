package logger

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/quangdat385/holiday-ticket/ticket-service/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.Config) *LoggerZap {
	logLevel := config.Logger.LogLevel
	fmt.Println("Setting config level", config.Logger.MaxSize)
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.Logger.FileLogName,
		MaxSize:    config.Logger.MaxSize, // megabytes
		MaxBackups: config.Logger.MaxBackups,
		MaxAge:     config.Logger.MaxAge,   //days
		Compress:   config.Logger.ComPress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)
	return &LoggerZap{
		zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}
}
func getEncoderLog() zapcore.Encoder {
	endcodeConfig := zap.NewProductionEncoderConfig()

	endcodeConfig.TimeKey = "time"

	endcodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	endcodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(endcodeConfig)
}
