package zlog

import (
	"os"
	"runtime"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// log file path
	logFilePath = "device.log"
	// log file maxSize(MB)
	maxSize = 100
	// the max number of log file
	maxBackups = 700
	// the max day of log file
	maxAge = 365

	timeLayout = "2006-01-02 15:04:05"
)

var logger *zap.Logger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeLayout)

	// encoder := zapcore.NewConsoleEncoder(encoderConfig)
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, getFileLogWriter(), zapcore.InfoLevel),
	)
	logger = zap.New(core)
}

func Info(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	logger.Info(message, append(fields, callerFields...)...)
}

func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	logger.Debug(message, append(fields, callerFields...)...)
}

func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	logger.Warn(message, append(fields, callerFields...)...)
}

func Error(message error, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	logger.Error(message.Error(), append(fields, callerFields...)...)
}

func Panic(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	logger.Panic(message, append(fields, callerFields...)...)
}

func Fatal(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	logger.Fatal(message, append(fields, callerFields...)...)
}

// Log cutting
func getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func getCallerInfoForLog() (callerFields []zap.Field) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}
	funcName := runtime.FuncForPC(pc).Name()

	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))

	return
}
