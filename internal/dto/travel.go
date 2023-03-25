package dto

import "time"

type Travel struct {
	ID 								int
  Price 						float64
  CompanyAccountId 	int
  BusID 						int
  DepartureTime 		time.Time
  DepartureCityId 	int
  ArrivalTime 			time.Time
  ArrivalCityId 		int
  CreatedAt 				time.Time
  UpdatedAt 				time.Time
}