package authz

import (
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	RoleAdmin    = 1
	RoleClient   = 2
	RoleCouriers = 4
	RoleAll      = RoleAdmin | RoleClient | RoleCouriers
)

type Claims struct {
	Roles int `json:"roles"`
	jwt.RegisteredClaims
}

func RequiresRole(roles int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(Claims)

			if claims.Roles&roles == 0 {
				slog.Warn("unauthorized access", "sub", claims.Subject, "role", claims.Roles)
				return echo.NewHTTPError(http.StatusForbidden)
			}

			return next(c)
		}
	}
}
