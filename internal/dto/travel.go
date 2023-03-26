package dto

import (
	"time"
)

type TravelInputDTO struct {
  Price 						float64
  AccountID 				int
  BusID 						int
  DepartureTime 		time.Time
  DepartureCityID 	int
  ArrivalTime 			time.Time
  ArrivalCityID 		int
}

type TravelOutputDTO struct {
	ID 								int
  Price 						float64
  FantasyName       string
  BusNumber      		int
  MaxPassengers 		int
  DepartureTime 		time.Time
  DepartureCity 	  string
  ArrivalTime 			time.Time
  ArrivalCity   		string
}