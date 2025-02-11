package ctxlog_test

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/weirdgiraffe/ctxlog"
)

func Example() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   a.Key,
					Value: slog.TimeValue(time.Date(2025, 2, 11, 0, 0, 0, 0, time.UTC)),
				}
			}
			return a
		},
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	log := ctxlog.NewSlog(logger)
	bgCtx := context.Background()
	log = log.With("example", "main")
	ctx := ctxlog.Embed(bgCtx, log)
	foo(ctx)

	// Output:
	// time=2025-02-11T00:00:00.000Z level=INFO msg="Hello, World!" example=main foo=foo
}

func foo(ctx context.Context) {
	log := ctxlog.From(ctx)
	ctx = ctxlog.Embed(ctx, log.With("foo", "foo"))
	bar(ctx)
}

func bar(ctx context.Context) {
	log := ctxlog.From(ctx)
	log.Info("Hello, World!")
}
