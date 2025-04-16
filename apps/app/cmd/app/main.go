package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v3"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/karagenc/fj4echo"
)

func run(ctx context.Context, _ *cli.Command) error {
	l := setupLog(cfg.Logging)
	e := echo.New()
	e.JSONSerializer = fj4echo.New()
	e.Validator = NewValidator()

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg.Database.GetConfig())
	if err != nil {
		return fmt.Errorf("unable to create database connection pool: %w", err)
	}
	defer pool.Close()

	urls(e, cfg, l, pool)

	f := &fasthttp.Server{
		Handler: fasthttpadaptor.NewFastHTTPHandler(e),
	}

	bindAddress := fmt.Sprintf("%s:%d", cfg.Listen.Host, cfg.Listen.Port)
	slog.Info("starting delivery service", "host", cfg.Listen.Host, "port", cfg.Listen.Port)
	slog.Debug("application configuration", "cfg", cfg)
	return f.ListenAndServe(bindAddress)
}

func main() {
	if err := args.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
