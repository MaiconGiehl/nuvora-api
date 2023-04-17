package dto

import "time"

type TicketOutputDTO struct {
	ID int
	StatusID int
	TravelID int
	CreatedAt time.Time
}

// func NewTicketOutputDTO(
// 	id int,
// 	status string,
// 	travelId int,
// 	travelPrice float64,
// 	companyId int,
// 	companyName string,
// 	departureTime time.Time,
// 	departureCity string,
// 	arrivalTime time.Time,
// 	arrivalCity string,
// ) *TicketOutputDTO {
// 	return &TicketOutputDTO{
// 		ID: id,
// 		Status: status,
// 		Travel: *NewTravelOutputDTO(
// 			travelId, 
// 			travelPrice, 
// 			companyId, 
// 			companyName,
// 		  departureTime,
// 			departureCity,
// 			arrivalTime,
// 			arrivalCity,
// 		),
// 	}
// }
