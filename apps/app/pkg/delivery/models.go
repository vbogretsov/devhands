package delivery

import "time"

type Spot struct {
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Country     string    `db:"country"`
	State       string    `db:"state"`
	City        string    `db:"city"`
	ID          int       `db:"id"`
	Raiting     float32   `db:"raiting"`
	Lat         float32   `db:"lat"`
	Lng         float32   `db:"lng"`
}

type Item struct {
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	SpotID      string    `db:"spot_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	ID          int       `db:"id"`
	Raiting     float32   `db:"raiting"`
}
