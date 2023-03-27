package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type GetTravelsByDestinyUseCase struct {
	TravelRepository database.TravelRepository
}

func NewGetAllTravelsByDestinyUseCase(
	TravelRepository database.TravelRepository,
) *GetTravelsByDestinyUseCase {
	return &GetTravelsByDestinyUseCase{
		TravelRepository: TravelRepository,
	}
}

func (c *GetTravelsByDestinyUseCase) Execute(input *dto.TravelInputDTO) (*[]dto.TravelOutputDTO, error) {
	entity := entity.Travel{
		ArrivalCityID:  input.ArrivalCityID,
		DepartureCityID: input.DepartureCityID,
	}

	output, err := c.TravelRepository.GetAllByDestiny(&entity)
	if err != nil {
		return output, err
	}

	return output, nil
}
