package dto

import "time"

type TravelInputDTO struct {
	Price     float64
	BusID     int
	Departure DepartureInputDTO
	Arrival   ArrivalInputDTO
}

type TravelOutputDTO struct {
	ID        int
	Price     float64
	AccountID int
	Departure struct {
		Time   time.Time
		CityID int
	}
	Arrival struct {
		Time   time.Time
		CityID int
	}
}

func NewTravelOutputDTO(
	id int,
	price float64,
	accountID int,
	departureTime time.Time,
	departureCityId int,
	arrivalTime time.Time,
	arrivalCityId int,
) *TravelOutputDTO {
	return &TravelOutputDTO{
		ID:        id,
		Price:     price,
		AccountID: accountID,
		Departure: struct {
			Time   time.Time
			CityID int
		}{
			Time:   departureTime,
			CityID: departureCityId,
		},
		Arrival: struct {
			Time   time.Time
			CityID int
		}{
			Time:   arrivalTime,
			CityID: arrivalCityId,
		},
	}
}
