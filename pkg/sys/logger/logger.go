package logger

import (
	"context"
	"io"
	"os"

	"golang.org/x/exp/slog"
)

const (
	envDisable = "disable"
	envLocal   = "local"
	envDev     = "dev"
	envProd    = "prod"
)

type loggerKey struct{}

// AssignLogger прокидывает логгер в контекст
func AssignLogger(ctx context.Context, logger *slog.Logger) context.Context {
	ctx = context.WithValue(ctx, loggerKey{}, logger)

	return ctx
}

// GetLogger получает логгер из контекста
func GetLogger(ctx context.Context) *slog.Logger {
	return ctx.Value(loggerKey{}).(*slog.Logger)
}

// SetupLogger создает объект логгера на основе типа окружения. Для локали и разработки текст, для прода json
func SetupLogger(env string) *slog.Logger {
	var lg *slog.Logger

	switch env {
	case envDisable:
		lg = slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case envLocal:
		lg = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		lg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		lg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return lg
}
