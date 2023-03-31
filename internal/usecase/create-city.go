package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
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
