package entity

import "time"

type Travel struct {
	ID int
  Price float64
  CompanyID string
  CompanyFantasyName string
	Bus struct {
		Number int
		MaxPassengers int
	}
	Departure struct {
		Time time.Time
		CityName string
	}
	Arrival struct {
		Time time.Time
		CityName string
	}
}
