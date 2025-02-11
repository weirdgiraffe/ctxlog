[![Go Reference](https://pkg.go.dev/badge/github.com/weirdgiraffe/ctxlog.svg)](https://pkg.go.dev/github.com/weirdgiraffe/ctxlog)

# Ctxlog

Small package to embed/fetch logger into the context

Usage:

```go
import (
    "context"
	"github.com/weirdgiraffe/ctxlog"
)

func foo(ctx context.Context) {
    log := ctxlog.From(ctx)
    log.Info("hello")
    bar(ctxlog.Embed(ctx, log.WithField("foo", "foo"))
}

func bar(ctx context.Context) {
    log := ctxlog.From(ctx)
    log.Info("world")
}
```

Right now only `slog` is supported, which could be inited as:

```go
log := ctxlog.NewSlog(slog.Default())
```

But feel free to add whichever logger implementation which fullfils the `Logger`
interface.
