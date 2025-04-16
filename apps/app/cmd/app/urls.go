package main

import (
	"log/slog"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	slogecho "github.com/samber/slog-echo"

	"app/pkg/authz"
	"app/pkg/delivery"
	"app/pkg/healthz"
	"app/pkg/tracker"
	"app/pkg/waitcpu"
)

func urls(e *echo.Echo, cfg Conf, logger *slog.Logger, pgpool *pgxpool.Pool) {
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(Metrics())

	e.GET("/healthz", healthz.Healthz)
	e.GET("/wait/:ms", waitcpu.WaitCPU)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(authz.Claims)
		},
		SigningKey: []byte(cfg.JWTKey),
	}

	// TODO: Make app more close to courriers trackign PoC and adjust roles.

	deliveryAPI := delivery.NewAPI(pgpool)

	d := e.Group("/api/delivery")
	d.Use(echojwt.WithConfig(jwtConfig))
	d.GET("/spots/:id", deliveryAPI.GetSpot, authz.RequiresRole(authz.RoleAll))
	d.POST("/spots", deliveryAPI.CreateSpot, authz.RequiresRole(authz.RoleAdmin))
	// d.GET("/spots/:id/items", deliveryAPI.ListSpotItems, authz.RequiresRole(authz.RoleAll))
	// d.POST("/spots/-/items", deliveryAPI.CreateSpotItems, authz.RequiresRole(authz.RoleAdmin))

	trackerAPI := tracker.NewAPI(tracker.NewRedisStorage(redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       int(cfg.Redis.Database),
		PoolSize: int(cfg.Redis.PoolSize),
	})))

	t := e.Group("/api/tracker")
	t.Use(echojwt.WithConfig(jwtConfig))
	t.POST("/locations", trackerAPI.SetLocation, authz.RequiresRole(authz.RoleAdmin))
	t.GET("/locations", trackerAPI.GetLocation, authz.RequiresRole(authz.RoleAdmin))
}
