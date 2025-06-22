package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// Initialize sets up the logger
func Initialize(environment string) {
    var config zap.Config
    
    if environment == "production" {
        config = zap.NewProductionConfig()
        config.EncoderConfig.TimeKey = "timestamp"
        config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    } else {
        config = zap.NewDevelopmentConfig()
        config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    }
    
    var err error
    log, err = config.Build()
    if err != nil {
        panic(err)
    }
}

// GetLogger returns the global logger instance
func GetLogger() *zap.Logger {
    if log == nil {
        // Default to development logger if not initialized
        Initialize("development")
    }
    return log
}

// Info logs an info message
func Info(message string, fields ...zap.Field) {
    GetLogger().Info(message, fields...)
}

// Error logs an error message
func Error(message string, fields ...zap.Field) {
    GetLogger().Error(message, fields...)
}

// Fatal logs a fatal message and exits
func Fatal(message string, fields ...zap.Field) {
    GetLogger().Fatal(message, fields...)
}

// Debug logs a debug message
func Debug(message string, fields ...zap.Field) {
    GetLogger().Debug(message, fields...)
}

// With returns a logger with additional fields
func With(fields ...zap.Field) *zap.Logger {
    return GetLogger().With(fields...)
}