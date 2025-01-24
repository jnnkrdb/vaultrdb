package logging

import (
	"log/slog"
	"os"
	"strings"
)

var SLog *slog.Logger = slog.Default()

// init the default logging instance
// can be initialized with following levels:
//   - Debug
//   - Info
//   - Warn
//   - everything else is Error
func InitSLOG(sloglevel string) {

	var logLevel slog.Level
	switch strings.ToLower(sloglevel) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	default:
		logLevel = slog.LevelError
	}

	SLog = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: (logLevel == slog.LevelDebug),
		Level:     logLevel,
	}))

	SLog.Debug("initiated slog", "LogLevel", logLevel.String())
}
