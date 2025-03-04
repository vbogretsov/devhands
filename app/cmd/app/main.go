package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	slogecho "github.com/samber/slog-echo"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	e := echo.New()

	e.Use(slogecho.New(l))
	e.Use(middleware.Recover())
	e.Use(Metrics())

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	fs := &fasthttp.Server{
		Handler: fasthttpadaptor.NewFastHTTPHandler(e),
	}

	if err := fs.ListenAndServe(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
