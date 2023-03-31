package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
)


type DeleteCityUseCase struct {
	CityRepository database.CityRepository
}

func NewDeleteCityUseCase(
	CityRepository database.CityRepository,
) *DeleteCityUseCase {
	return &DeleteCityUseCase{
		CityRepository: CityRepository,
	}
}

func (c *DeleteCityUseCase) Execute(id int) error {
	entity := entity.City{
		ID:								id,
	}

	err := c.CityRepository.Delete(&entity)
	if err != nil {
		return err
	}

	return nil
}
