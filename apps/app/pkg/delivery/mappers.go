package delivery

import (
	"app/pkg/remapper"
)

var (
	spotmap = spotMapper{
		d2m: remapper.New(&SpotInDTO{}, &Spot{}),
		m2d: remapper.New(&Spot{}, &SpotOutDTO{}),
	}
)

type spotMapper struct {
	d2m remapper.Mapper
	m2d remapper.Mapper
}

func (m *spotMapper) ToModel(dto SpotInDTO) Spot {
	var model Spot
	m.d2m.Map(&dto, &model)
	return model
}

func (m *spotMapper) ToDTO(model Spot) SpotOutDTO {
	var dto SpotOutDTO
	m.m2d.Map(&model, &dto)
	return dto
}
