package usecase

import (
	"errors"
	"time"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
)

type createTravelCommand struct {
	CompanyID       int
	Price           float64
	BusID           int
	DepartureTime   time.Time
	DepartureCityID int
	ArrivalTime     time.Time
	ArrivalCityID   int
}

func With(
	companyID int,
	travel dto.TravelInputDTO,
) (*createTravelCommand, error) {

	dptTime, err := time.Parse("2006-01-02T00:00:00", travel.Departure.Time)
	if err != nil {
		return nil, errors.New("departure time must be in time format")
	}

	arvTime, err := time.Parse("2006-01-02T00:00:00", travel.Arrival.Time)
	if err != nil {
		return nil, errors.New("arrival time must be in time format")
	}

	return &createTravelCommand{
		CompanyID:       companyID,
		Price:           travel.Price,
		BusID:           travel.BusID,
		DepartureTime:   dptTime,
		DepartureCityID: travel.Departure.CityID,
		ArrivalTime:     arvTime,
		ArrivalCityID:   travel.Arrival.CityID,
	}, nil
}
