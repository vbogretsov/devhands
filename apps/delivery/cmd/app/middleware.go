package main

import (
	"net/http"
	// "syscall"
	"time"
	// "log/slog"

	"github.com/labstack/echo/v4"

	"app/pkg/metrics"
)

/* func getrusage() (*syscall.Rusage, error) {
	var res syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &res); err != nil {
		slog.Warn("getrusage call failed", "err", err)
		return nil, err
	}
	return &res, nil
} */

func Metrics() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			duration := time.Since(start).Seconds()

			status := c.Response().Status
			method := c.Request().Method
			path := c.Path()

			metrics.RequestsTotal.WithLabelValues(method, path, http.StatusText(status)).Inc()
			metrics.RequestDuration.WithLabelValues(method, path).Observe(duration)

			return err
		}
	}
}
