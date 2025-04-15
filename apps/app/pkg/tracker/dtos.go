package tracker

type ListDTO[T any] struct {
	Items []T `json:"items"`
}

type LocationDTO struct {
	UserID string  `json:"user_id"`
	Lat    float32 `json:"lat"`
	Lng    float32 `json:"lng"`
}
