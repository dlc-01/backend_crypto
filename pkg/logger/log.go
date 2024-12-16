package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	DevelopMode = "develop"
	ProductMode = "product"
)

type ConfLogger struct {
	AppMod string `env:"APP_ENV" envDefault:"develop" json:"AppMod"`
	File   struct {
		Directory  string `env:"LOGGER_DIRECTORY" envDefault:"temp/logs/" json:"LoggerDirectory"`
		MaxSize    int    `env:"LOGGER_FILE_MAX_SIZE" envDefault:"1" json:"LoggerMaxSize"`
		MaxBackups int    `env:"LOGGER_FILE_MAX_BACKUPS" envDefault:"1" json:"LoggerMaxBackups"`
		MaxAge     int    `env:"LOGGER_FILE_MAX_AGE" envDefault:"1" json:"LoggerMaxAge"`
		Compress   bool   `env:"LOGGER_FILE_COMPRESS" envDefault:"true" json:"LoggerCompress"`
	}
}

type Logger struct {
	zap.SugaredLogger
}

// InitLogger — initialization function of zap logger.
func Initialize(cfg ConfLogger) (*Logger, error) {
	switch cfg.AppMod {
	case DevelopMode:
		logger, err := zap.NewDevelopment()
		if err != nil {
			return nil, fmt.Errorf("cannot init logger: %w", err)
		}
		lgr := &Logger{
			SugaredLogger: *logger.Sugar(),
		}
		return lgr, nil

	case ProductMode:
		loggerConf := zap.NewProductionEncoderConfig()
		loggerConf.EncodeTime = zapcore.ISO8601TimeEncoder
		fileEncoder := zapcore.NewJSONEncoder(loggerConf)
		defaultLogLevel := zapcore.DebugLevel
		dir := fmt.Sprintf("%v/", cfg.File.Directory)
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return nil, err
		}
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%v/%v.log", dir, time.Now().Format("2022-02-24")),
			MaxSize:    cfg.File.MaxSize,
			MaxBackups: cfg.File.MaxBackups,
			MaxAge:     cfg.File.MaxAge,
			Compress:   cfg.File.Compress,
		})
		core := zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
		lgr := &Logger{
			SugaredLogger: *zap.New(zapcore.NewTee(core)).Sugar(),
		}
		return lgr, nil
	default:
		return nil, fmt.Errorf("projectError while: creating logger, not supported mode")
	}
}
