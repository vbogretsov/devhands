package remapper

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SpotInDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`

	ID      int     `json:"id"`
	Raiting float32 `json:"raiting"`
	Lat     float32 `json:"lat"`
	Lng     float32 `json:"lng"`
}

type Spot struct {
	CreateAt  time.Time
	UpdatedAt time.Time

	Name        string
	Description string
	Country     string
	State       string
	City        string

	ID      int
	Raiting float32
	Lat     float32
	Lng     float32
}

func TestMap(t *testing.T) {
	d := SpotInDTO{
		ID:          1,
		Name:        "Test Spot",
		Description: "A beautiful place",
		Raiting:     4.5,
		Country:     "US",
		State:       "CA",
		City:        "San Francisco",
		Lat:         37.7749,
		Lng:         -122.4194,
	}
	m := Spot{}
	mapper := New(&d, &m)
	mapper.Map(&d, &m)

	assert.Equal(t, d.ID, m.ID, "ID should match")
	assert.Equal(t, d.Name, m.Name, "Name should match")
	assert.Equal(t, d.Description, m.Description, "Description should match")
	assert.Equal(t, d.Raiting, m.Raiting, "Raiting should match")
	assert.Equal(t, d.Country, m.Country, "Country should match")
	assert.Equal(t, d.State, m.State, "State should match")
	assert.Equal(t, d.City, m.City, "City should match")
	assert.Equal(t, d.Lat, m.Lat, "Lat should match")
	assert.Equal(t, d.Lng, m.Lng, "Lng should match")
}

func BenchmarkDirectAssignment(b *testing.B) {
	src := SpotInDTO{
		ID:          1,
		Name:        "Test Spot",
		Description: "A beautiful place",
		Raiting:     4.5,
		Country:     "US",
		State:       "CA",
		City:        "San Francisco",
		Lat:         37.7749,
		Lng:         -122.4194,
	}
	var dst Spot

	var result *Spot

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst.ID = src.ID
		dst.Name = src.Name
		dst.Description = src.Description
		dst.Raiting = src.Raiting
		dst.Country = src.Country
		dst.State = src.State
		dst.City = src.City
		dst.Lat = src.Lat
		dst.Lng = src.Lng

		result = &dst
	}

	runtime.KeepAlive(result)
}

func BenchmarkUnsafeMapper(b *testing.B) {
	src := SpotInDTO{
		ID:          1,
		Name:        "Test Spot",
		Description: "A beautiful place",
		Raiting:     4.5,
		Country:     "US",
		State:       "CA",
		City:        "San Francisco",
		Lat:         37.7749,
		Lng:         -122.4194,
	}
	var dst Spot

	mapper := New(&src, &dst)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mapper.Map(&src, &dst)
	}
}
