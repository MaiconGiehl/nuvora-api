package entity

import "time"

type Ticket struct {
	ID 								int
	ClienteID 				int
	Price							float64
	Status 						int
	DepartureCity 		string
	DepartureTime 		time.Time
	DestinyCity 			string
	DestinyTime 			time.Time
	BusID 						int
}