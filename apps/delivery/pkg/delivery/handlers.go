package delivery

import (
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"app/pkg/authz"
)

type API struct {
	pool *pgxpool.Pool
}

func NewAPI(pool *pgxpool.Pool) *API {
	return &API{pool: pool}
}

func (a *API) CreateSpot(c echo.Context) error {
	var dto SpotInDTO
	if err := c.Bind(&dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(&dto); err != nil {
		return err
	}
	return nil
}

func (a *API) ListSpots(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	slog.Info("")
	return nil
}

func (a *API) CreateSpotItem(c echo.Context) error {
	return nil
}

func (a *API) ListSpotItems(c echo.Context) error {
	return nil
}
