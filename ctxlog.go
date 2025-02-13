package ctxlog

import (
	"context"
)

type Level int

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

type Logger interface {
	With(args ...any) Logger
	WithGroup(name string) Logger

	Log(ctx context.Context, level Level, msg string, args ...any)

	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type loggerCtxKey struct{}

// Embed adds the logger to the provided context
func Embed(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

// From returns the logger from the provided context
func From(ctx context.Context) Logger {
	if v := ctx.Value(loggerCtxKey{}); v != nil {
		return v.(Logger)
	}
	return discardLogger{}
}

type discardLogger struct{}

func (l discardLogger) With(args ...any) Logger                                     { return l }
func (l discardLogger) WithGroup(name string) Logger                                { return l }
func (discardLogger) Log(ctx context.Context, level Level, msg string, args ...any) {}
func (discardLogger) Debug(msg string, args ...any)                                 {}
func (discardLogger) Info(msg string, args ...any)                                  {}
func (discardLogger) Warn(msg string, args ...any)                                  {}
func (discardLogger) Error(msg string, args ...any)                                 {}
