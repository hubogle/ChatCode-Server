package log

import (
	"log/slog"
	"os"
	"time"
)

func NewSlog(format string, level int) {
	opts := slog.HandlerOptions{
		Level: slog.Level(level),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				a.Value = slog.StringValue(t.Format(time.DateTime))
				return a
			}

			return a
		},
	}
	var h slog.Handler
	if format == "json" {
		h = slog.NewJSONHandler(os.Stderr, &opts)
	} else {
		h = slog.NewTextHandler(os.Stderr, &opts)
	}

	slog.SetDefault(slog.New(h))
}
