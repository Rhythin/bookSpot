package customlogger

import (
	// Using standard log for initial setup messages, not the global Zap logger yet
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger initializes and sets up a global Zap logger
// It replaces the global Zap logger and redirects Go's standard log to Zap.
func InitLogger() (logger *zap.Logger, err error) {

	var zapCfg zap.Config // nolint

	// TODO: for now we are using the development config
	zapCfg = zap.NewDevelopmentConfig()

	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// Build the logger
	logger, err = zapCfg.Build()
	if err != nil {
		log.Fatalln("failed to create logger", err)
		return logger, err
	}

	// Replace Zap's global logger
	zap.ReplaceGlobals(logger)

	// Redirect Go's standard log messages to Zap.
	zap.RedirectStdLog(logger)

	return logger, err
}

// S returns a SugaredLogger instance
// this will be later used to handling logging along with trace and span
func S() *zap.SugaredLogger {
	return zap.S()
}
