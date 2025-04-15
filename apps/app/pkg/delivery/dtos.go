package delivery

import "time"

type SpotInDTO struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Country     string  `json:"country" validate:"iso3166_1_alpha2"`
	State       string  `json:"state"`
	City        string  `json:"city" validate:"required"`
	ID          int     `json:"id" validate:"required,gte=0"`
	Raiting     float32 `json:"raiting" validate:"gte=0"`
	Lat         float32 `json:"lat" validate:"required,latitude"`
	Lng         float32 `json:"lng" validate:"required,longitude"`
}

type SpotOutDTO struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Country     string    `json:"country"`
	State       string    `json:"state"`
	City        string    `json:"city"`
	ID          int       `json:"id"`
	Raiting     float32   `json:"raiting"`
	Lat         float32   `json:"lat"`
	Lng         float32   `json:"lng"`
}

type ItemInDTO struct {
	SpotID      string  `json:"spot_id" validate:"required,gte=0"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	ID          int     `json:"id" validate:"required,gte=0"`
	Raiting     float32 `json:"raiting" validate:"gte=0"`
}

type ItemOutDTO struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	SpotID      string    `json:"spot_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
	Raiting     float32   `json:"raiting"`
}
