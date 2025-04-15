package delivery

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type API struct {
	storage Storage
}

func NewAPI(pool *pgxpool.Pool) *API {
	return &API{NewStorage(pool)}
}

func (a *API) CreateSpot(c echo.Context) error {
	var req SpotInDTO
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	model, err := a.storage.CreateSpot(c.Request().Context(), spotmap.ToModel(req))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, spotmap.ToDTO(model))
}

func (a *API) GetSpot(c echo.Context) error {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id must be integer")
	}

	model, err := a.storage.GetSpot(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			slog.Warn("spot not found", "id", id)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return err
	}

	return c.JSON(http.StatusOK, spotmap.ToDTO(model))
}

func (a *API) CreateSpotItems(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}

func (a *API) ListSpotItems(c echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented)
}
