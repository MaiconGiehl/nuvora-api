package dto

import "time"

type TravelInputDTO struct {
	Price     float64
	BusID     int
	Departure DepartureInputDTO
	Arrival   ArrivalInputDTO
}

type TravelOutputDTO struct {
	ID          int
	Price       float64
	CompanyID   int
	CompanyName string
	Departure   struct {
		Time time.Time
		City string
	}
	Arrival struct {
		Time time.Time
		City string
	}
}

func NewTravelOutputDTO(
	id int,
	price float64,
	companyId int,
	companyName string,
	departureTime time.Time,
	departureCity string,
	arrivalTime time.Time,
	arrivalCity string,
) *TravelOutputDTO {
	return &TravelOutputDTO{
		ID:          id,
		Price:       price,
		CompanyID:   companyId,
		CompanyName: companyName,
		Departure:   *NewDepartureOutputDTO(departureTime, departureCity),
		Arrival:     *NewArrivalOutputDTO(arrivalTime, arrivalCity),
	}
}
