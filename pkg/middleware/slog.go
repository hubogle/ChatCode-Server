package middleware

import (
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GinSlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start).Seconds()
		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				slog.Error(e)
			}
		} else {
			slog.Info("",
				slog.String("method", c.Request.Method),
				slog.String("host", c.Request.Host),
				slog.String("path", path),
				slog.String("query", query),
				slog.String("proto", c.Request.Proto),
				slog.String("client_ip", c.ClientIP()),
				slog.Float64("latency", latency))
		}
	}
}

func RecoverySlog(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					slog.Error(
						"[Recovery from panic]",
						slog.Any("error", err),
						slog.String("request", string(httpRequest)),
					)

					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					slog.Error(
						"[Recovery from panic]",
						slog.Time("time", time.Now()),
						slog.Any("error", err),
						slog.String("request", string(httpRequest)),
						slog.String("stack", string(debug.Stack())),
					)
				} else {
					slog.Error(
						"[Recovery from panic]",
						slog.Time("time", time.Now()),
						slog.Any("error", err),
						slog.String("request", string(httpRequest)),
					)
				}

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
