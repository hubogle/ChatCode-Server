package middleware

import (
	"log/slog"
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
