package entity

import "time"

type Travel struct {
	ID 								int
  Price 						float64
  AccountID 	int
  BusID 						int
  DepartureTime 		time.Time
  DepartureCityID 	int
  ArrivalTime 			time.Time
  ArrivalCityID 		int
  CreatedAt 				time.Time
  UpdatedAt 				time.Time
}