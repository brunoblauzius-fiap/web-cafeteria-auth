package logging

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(s string, fields ...Field)
	Info(s string, fields ...Field)
	Error(s string, fields ...Field)
	Warn(s string, fields ...Field)
	Fatal(s string, fields ...Field)
	With(fields ...Field) Logger
}

type Log struct {
	log *zap.Logger
}

type Field = zap.Field

var (
	String    = zap.String
	Error     = zap.Error
	Int       = zap.Int
	Int32p    = zap.Int32p
	Int32     = zap.Int32
	Int64     = zap.Int64
	Durationp = zap.Durationp
	Any       = zap.Any
	Float64   = zap.Float64
)

func New() (*Log, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)

	zapLogger, err := config.Build(zap.AddCallerSkip(1))

	if err != nil {
		return nil, fmt.Errorf("error building logging: %w", err)
	}

	return &Log{
		log: zapLogger,
	}, nil
}

func (l *Log) Flush() {
	err := l.log.Sync()
	if err != nil {
		fmt.Println("Error flushing logging", err)
	}
}

func (l *Log) Debug(s string, fields ...Field) {
	l.log.Debug(s, fields...)
}

func (l *Log) Info(s string, fields ...Field) {
	l.log.Info(s, fields...)
}

func (l *Log) Error(s string, fields ...Field) {
	l.log.Error(s, fields...)
}

func (l *Log) Warn(s string, fields ...Field) {
	l.log.Warn(s, fields...)
}

func (l *Log) Fatal(s string, fields ...Field) {
	l.log.Fatal(s, fields...)
}

func (l *Log) With(fields ...Field) Logger {
	zapLogger := l.log.With(fields...)
	return &Log{
		log: zapLogger,
	}
}
