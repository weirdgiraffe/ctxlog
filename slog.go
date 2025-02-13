package ctxlog

import (
	"context"
	"log/slog"
)

func NewSlog(l *slog.Logger) Logger {
	return &slogLogger{impl: l}
}

type slogLogger struct {
	impl *slog.Logger
}

func (l *slogLogger) With(args ...any) Logger      { return &slogLogger{impl: l.impl.With(args...)} }
func (l *slogLogger) WithGroup(name string) Logger { return &slogLogger{impl: l.impl.WithGroup(name)} }

func (l *slogLogger) Log(ctx context.Context, level Level, msg string, args ...any) {
	l.impl.Log(ctx, slog.Level(level), msg, args...)
}

func (l *slogLogger) Debug(msg string, args ...any) { l.impl.Debug(msg, args...) }
func (l *slogLogger) Info(msg string, args ...any)  { l.impl.Info(msg, args...) }
func (l *slogLogger) Warn(msg string, args ...any)  { l.impl.Warn(msg, args...) }
func (l *slogLogger) Error(msg string, args ...any) { l.impl.Error(msg, args...) }
