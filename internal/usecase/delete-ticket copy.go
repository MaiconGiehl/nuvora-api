package usecase

import (
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type DeleteTravelUseCase struct {
	TravelRepository database.TravelRepository
}

func NewDeleteTravelUseCase(
	TravelRepository database.TravelRepository,
) *DeleteTravelUseCase {
	return &DeleteTravelUseCase{
		TravelRepository: TravelRepository,
	}
}

func (c *DeleteTravelUseCase) Execute(id int) error {
	entity := entity.Travel{
		ID:								id,
	}

	err := c.TravelRepository.Delete(&entity)
	if err != nil {
		return err
	}

	return nil
}