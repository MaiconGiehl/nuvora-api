package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
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
