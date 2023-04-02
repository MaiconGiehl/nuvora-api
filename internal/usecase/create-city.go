package usecase

import (
	"github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/entity"
	"github.com/maicongiehl/nuvera-api/internal/infra/database"
)


type CreateCityUseCase struct {
	CityRepository database.CityRepository
}

func NewCreateCityUseCase(
	CityRepository database.CityRepository,
) *CreateCityUseCase {
	return &CreateCityUseCase{
		CityRepository: CityRepository,
	}
}

func (c *CreateCityUseCase) Execute(input *dto.CityInputDTO) error {
	entity := entity.City{
		Name: input.Name,	
	}

	err := c.CityRepository.Save(&entity)
	if err != nil {
		return err
	}

	return nil
}
