package dto

import "time"

type TravelInputDTO struct {
  Price float64
  BusID int
  Departure DepartureInputDTO
  Arrival ArrivalInputDTO
}

type TravelOutputDTO struct {
	ID int
	Price float64
	*BusOutputDTO
	AccountID int
	Departure struct {
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
	busId int,
	busNumber int,
	busMaxPassengers,
	accountID int,
	departureTime time.Time,
	departureCity string,
	arrivalTime time.Time,
	arrivalCity string,
) *TravelOutputDTO {
	return &TravelOutputDTO{
		ID: id,
		Price: price,
		BusOutputDTO: NewBusOutputDTO(busId, busNumber, busMaxPassengers, accountID),
		AccountID: accountID,
		Departure: struct{Time time.Time; City string}{
			Time: departureTime, 
			City: departureCity,
		},
		Arrival: struct{Time time.Time; City string}{
			Time: arrivalTime, 
			City: arrivalCity,
		},
	}
}