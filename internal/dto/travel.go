package dto

import "time"

type TravelInputDTO struct {
  Price 						float64
  CompanyAccountId 	int
  BusID 						int
  DepartureTime 		time.Time
  DepartureCityId 	int
  ArrivalTime 			time.Time
  ArrivalCityId 		int
}

type TravelOutputDTO struct {
	ID 								int
  Price 						float64
  Company         	CompanyOutputDTO
  Bus 			   			BusOutputDTO
  DepartureTime 		time.Time
  DepartureCity 	  CityOutputDTO
  ArrivalTime 			time.Time
  ArrivalCity   		CityOutputDTO
}