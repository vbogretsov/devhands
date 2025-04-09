package delivery

import "time"

type Spot struct {
	ID          int
	CreateAt    time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Raiting     float32
	Country     string
	State       string
	City        string
	Lat         float32
	Lng         float32
}

type Item struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SpotID      string
	Name        string
	Description string
	Raiting     float32
}
