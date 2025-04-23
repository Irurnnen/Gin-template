// New creates and returns a new instance of a zap.Logger based on the provided
// log level and debug mode. The function initializes the logger in development,
// example, or production mode depending on the Debug flag. It also adjusts the
// logging level based on the specified LogLevel parameter.
//
// Parameters:
//   - LogLevel: A string representing the desired logging level. Supported values
//     are "debug", "info", "warn", and "error". If an unsupported value is provided,
//     the default level is set to "info".
//   - Debug: A boolean flag indicating whether to enable debug mode. If true, the
//     logger is initialized in example mode; otherwise, it is initialized in production mode.
//
// Returns:
//   - *zap.Logger: A pointer to the configured zap.Logger instance.
//
// Note: The logger's Sync method is deferred within the function, which may not
// be effective as the logger is returned to the caller. Ensure to call Sync on
// the returned logger instance in the calling code to flush any buffered log entries.
package logger

import "go.uber.org/zap"

func New(LogLevel string, Debug bool) *zap.Logger {
	logger, _ := zap.NewDevelopment()
	if Debug {
		logger = zap.NewExample()
	} else {
		logger, _ = zap.NewProduction()
	}

	switch LogLevel {
	case "debug":
		logger = logger.WithOptions(zap.IncreaseLevel(zap.DebugLevel))
	case "info":
		logger = logger.WithOptions(zap.IncreaseLevel(zap.InfoLevel))
	case "warn":
		logger = logger.WithOptions(zap.IncreaseLevel(zap.WarnLevel))
	case "error":
		logger = logger.WithOptions(zap.IncreaseLevel(zap.ErrorLevel))
	default:
		logger = logger.WithOptions(zap.IncreaseLevel(zap.InfoLevel))
	}
	defer logger.Sync()
	return logger
}
