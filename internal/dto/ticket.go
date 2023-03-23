package dto

import "time"

type TicketInputDTO struct {
	ClienteID 				int
	Price							float64
	Status 						int
	DepartureCity 		string
	DepartureTime 		time.Time
	DestinyCity 			string
	DestinyTime 			time.Time
	BusID 						int
}

type TicketOutputDTO struct {
	ID 								int
	Price							float64
	Status 						int
	DepartureCity 		string
	DepartureTime 		time.Time
	DestinyCity 			string
	DestinyTime 			time.Time
	BusNumber 				int
}