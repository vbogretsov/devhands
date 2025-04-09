package delivery

func spotDTO2Model(dto SpotInDTO) Spot {
	return Spot{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Raiting:     dto.Raiting,
		Country:     dto.Country,
		State:       dto.State,
		City:        dto.City,
		Lat:         dto.Lat,
		Lng:         dto.Lng,
	}
}

func spotModel2DTO(model Spot) SpotOutDTO {
	return SpotOutDTO{
		ID:          model.ID,
		CreatedAt:   model.CreateAt,
		UpdatedAt:   model.UpdatedAt,
		Name:        model.Name,
		Description: model.Description,
		Raiting:     model.Raiting,
		Country:     model.Country,
		State:       model.State,
		City:        model.City,
		Lat:         model.Lat,
		Lng:         model.Lng,
	}
}

func itemDTO2Model(dto ItemInDTO) Item {
	return Item{
		ID:          dto.ID,
		SpotID:      dto.SpotID,
		Name:        dto.Name,
		Description: dto.Description,
		Raiting:     dto.Raiting,
	}
}

func itemModel2DTO(model ItemOutDTO) ItemOutDTO {
	return ItemOutDTO{
		ID:          model.ID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		SpotID:      model.SpotID,
		Name:        model.Name,
		Description: model.Description,
		Raiting:     model.Raiting,
	}
}
