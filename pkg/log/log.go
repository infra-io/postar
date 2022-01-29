package log

import (
	"github.com/FishGoddess/logit"
	"github.com/avinoplan/postar/configs"
)

var (
	// globalLogger is the logger using in global.
	globalLogger *logit.Logger
)

// Initialize initializes log package with config.
func Initialize(c *configs.Config) error {
	options := logit.Options()
	globalLogger = logit.NewLogger(
		options.WithCallerDepth(4),
	)
	return nil
}

// Debug returns a Log with debug level if debug level is enabled.
func Debug(msg string, args ...interface{}) *logit.Log {
	return globalLogger.Debug(msg, args)
}

// Info returns a Log with info level if info level is enabled.
func Info(msg string, args ...interface{}) *logit.Log {
	return globalLogger.Info(msg, args)
}

// Warn returns a Log with warn level if warn level is enabled.
func Warn(msg string, args ...interface{}) *logit.Log {
	return globalLogger.Warn(msg, args)
}

// Error adds an entry which key is string and value is error type to l.
func Error(err error, msg string, args ...interface{}) *logit.Log {
	return globalLogger.Error(msg, args).Error("err", err)
}

func Close() error {
	return globalLogger.Close()
}
