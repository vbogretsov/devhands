package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/wait/:ms", WaitCPU)
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	f := &fasthttp.Server{
		Handler: fasthttpadaptor.NewFastHTTPHandler(e),
	}

	if err := f.ListenAndServe(":8000"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
