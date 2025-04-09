package delivery

import "time"

type SpotInDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Raiting     float32 `json:"raiting"`
	Country     string  `json:"country"`
	State       string  `json:"state"`
	City        string  `json:"city"`
	Lat         float32 `json:"lat"`
	Lng         float32 `json:"lng"`
}

type SpotOutDTO struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Raiting     float32   `json:"raiting"`
	Country     string    `json:"country"`
	State       string    `json:"state"`
	City        string    `json:"city"`
	Lat         float32   `json:"lat"`
	Lng         float32   `json:"lng"`
}

type ItemInDTO struct {
	ID          int     `json:"id"`
	SpotID      string  `json:"spot_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Raiting     float32 `json:"raiting"`
}

type ItemOutDTO struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	SpotID      string    `json:"spot_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Raiting     float32   `json:"raiting"`
}
