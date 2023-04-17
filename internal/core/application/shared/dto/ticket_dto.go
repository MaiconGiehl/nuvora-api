package dto

import "time"

type TicketOutputDTO struct {
	ID int
	StatusID int
	TravelOutputDTO
	CreatedAt time.Time
}

func NewTicketOutputDTO(
	id int,
	status string,
	travelId int,
	travelPrice float64,
	busId int,
	busNumber int,
	busMaxPassengers,
	companyId int,
	companyName string,
	departureTime time.Time,
	departureCity string,
	arrivalTime time.Time,
	arrivalCity string,
) *TicketOutputDTO {
	return &TicketOutputDTO{
		ID: id,
		TravelOutputDTO: *NewTravelOutputDTO(
			travelId,
			travelPrice,
			busId,
			busNumber,
			busMaxPassengers,
			companyId,
			departureTime,
			departureCity,
			arrivalTime,
			arrivalCity,
		),
	}
}