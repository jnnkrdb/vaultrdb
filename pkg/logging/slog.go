package logging

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

var Default *slog.Logger = GetLogger(true, "debug")

// get a logging instance
// can be initialized with following levels:
//   - Debug
//   - Info
//   - Warn
//   - Error (default)
func GetLogger(isJsonFormat bool, sloglevel string) *slog.Logger {

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

	handlerOpts := &slog.HandlerOptions{
		AddSource: (logLevel == slog.LevelDebug),
		Level:     logLevel,
	}

	if isJsonFormat {
		return slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))
	} else {
		return slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))
	}
}

// get a logger from context
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(slog.Logger{}).(*slog.Logger); ok {
		return logger
	} else {
		return Default
	}
}

// insert a logger into a given context
func IntoContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, slog.Logger{}, logger)
}
