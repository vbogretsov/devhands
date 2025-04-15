package tracker

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrNotFound = errors.New("not found")

type Storage interface {
	SetLocation(context.Context, LocationDTO) error
	GetLocation(context.Context, string) (LocationDTO, error)
}

type API struct {
	storage Storage
}

func NewAPI(storage Storage) *API {
	return &API{storage}
}

func (a *API) SetLocation(c echo.Context) error {
	var dto LocationDTO
	if err := c.Bind(&dto); err != nil {
		return err
	}

	if err := c.Validate(dto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := a.storage.SetLocation(c.Request().Context(), dto); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

type LocationQuery struct {
	UserID string `query:"user_id"`
}

func (a *API) GetLocation(c echo.Context) error {
	var query LocationQuery
	if err := c.Bind(&query); err != nil {
		return err
	}

	res, err := a.storage.GetLocation(c.Request().Context(), query.UserID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			slog.Warn("location not found", "user", query.UserID)
			return echo.NewHTTPError(http.StatusNotFound)
		}
	}

	return c.JSON(http.StatusOK, ListDTO[LocationDTO]{Items: []LocationDTO{res}})
}
