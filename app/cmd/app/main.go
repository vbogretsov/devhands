package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	slogecho "github.com/samber/slog-echo"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"app/pkg/handlers"
)

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	e := echo.New()

	e.Use(slogecho.New(l))
	e.Use(middleware.Recover())
	e.Use(Metrics())

	e.GET("/api/hello", handlers.Hello)
	e.GET("/api/wait/:ms", handlers.WaitCPU)

	e.GET("/api/metrics", echo.WrapHandler(promhttp.Handler()))

	fs := &fasthttp.Server{
		Handler: fasthttpadaptor.NewFastHTTPHandler(e),
	}

	if err := fs.ListenAndServe(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
