package delivery

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("not found")

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) Storage {
	return Storage{pool}
}

func (s *Storage) CreateSpot(ctx context.Context, spot Spot) (Spot, error) {
	query := `
	INSERT INTO sm_spot (
		id, created_at, updated_at, name, description,
		country, state, city, lat, lng, raiting
	) VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, $8, $9, $10, $11
	)
	`

	spot.CreatedAt = time.Now()
	spot.UpdatedAt = time.Now()

	con, err := s.pool.Acquire(ctx)
	if err != nil {
		return spot, fmt.Errorf("failed to acuire connection from pool: %w", err)
	}
	defer con.Release()

	_, err = con.Exec(
		ctx,
		query,
		spot.ID,
		spot.CreatedAt,
		spot.UpdatedAt,
		spot.Name,
		spot.Description,
		spot.Country,
		spot.State,
		spot.City,
		spot.Lat,
		spot.Lng,
		spot.Raiting,
	)

	return spot, nil
}

func (s *Storage) GetSpot(ctx context.Context, id int) (Spot, error) {
	query := `
	SELECT
		id,
		created_at,
		updated_at,
		name,
		description,
		raiting,
		country,
		state,
		city,
		lat,
		lng
	FROM sm_spot WHERE id = $1
	`
	var result Spot

	con, err := s.pool.Acquire(ctx)
	if err != nil {
		return result, err
	}
	defer con.Release()

	// TODO: Use struct's "db" tags.
	err = con.QueryRow(ctx, query, id).Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Name,
		&result.Description,
		&result.Raiting,
		&result.Country,
		&result.State,
		&result.City,
		&result.Lat,
		&result.Lng,
	)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return result, ErrNotFound
	}
	return result, err
}
