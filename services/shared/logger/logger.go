package logger

import (
	// Using standard log for initial setup messages, not the global Zap logger yet
	"go.uber.org/zap"
)

// InitLogger initializes and sets up a global Zap logger
// It replaces the global Zap logger and redirects Go's standard log to Zap.
func InitLogger() (logger *zap.Logger) {

	var zapCfg zap.Config // nolint

	// TODO: for now we are using the development config
	zapCfg = zap.NewDevelopmentConfig()

	// Build the logger
	logger = zap.Must(zapCfg.Build())

	// Replace Zap's global logger
	zap.ReplaceGlobals(logger)

	// Redirect Go's standard log messages to Zap.
	zap.RedirectStdLog(logger)

	return logger
}
