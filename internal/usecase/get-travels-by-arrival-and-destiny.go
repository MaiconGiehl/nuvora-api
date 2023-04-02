package usecase

import (
	"github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/entity"
	"github.com/maicongiehl/nuvera-api/internal/infra/database"
)


type GetTravelsByArrivalAndDeparture struct {
	TravelRepository database.TravelRepository
}

func NewGetTravelsByArrivalAndDeparture(
	TravelRepository database.TravelRepository,
) *GetTravelsByArrivalAndDeparture {
	return &GetTravelsByArrivalAndDeparture{
		TravelRepository: TravelRepository,
	}
}

func (c *GetTravelsByArrivalAndDeparture) Execute(input *dto.TravelInputDTO) (*[]dto.TravelOutputDTO, error) {
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
