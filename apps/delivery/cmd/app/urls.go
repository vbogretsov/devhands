package main

import (
	"app/pkg/authz"
	"app/pkg/delivery"
	"app/pkg/healthz"
	"app/pkg/waitcpu"
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	slogecho "github.com/samber/slog-echo"
	"github.com/jackc/pgx/v5/pgxpool"
)

func urls(e *echo.Echo, cfg Conf, logger *slog.Logger, pgpool *pgxpool.Pool) {
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())
	e.Use(Metrics())

	e.GET("/healthz", healthz.Healthz)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/wait/:ms", waitcpu.WaitCPU)

	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(authz.Claims)
		},
		SigningKey: []byte(cfg.JWTKey),
	}

	deliveryAPI := delivery.NewAPI(pgpool)

	r := e.Group("/api")
	r.Use(echojwt.WithConfig(jwtConfig))
	r.GET("/spots", deliveryAPI.ListSpots, authz.RequiresRole(authz.RoleAll))
	r.POST("/spots", deliveryAPI.CreateSpot, authz.RequiresRole(authz.RoleAdmin))
	r.GET("/spots/:id/items", deliveryAPI.ListSpotItems, authz.RequiresRole(authz.RoleAll))
	r.POST("/spots/-/items", deliveryAPI.CreateSpotItem, authz.RequiresRole(authz.RoleAdmin))
}
