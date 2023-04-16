package entity

import (
	"database/sql"
	"time"
)

type Travel struct {
	ID int
  Price float64
  AccountID int
	BusID int
	Status int
	Departure struct {
		Time time.Time
		CityID int
	}
	Arrival struct {
		Time time.Time
		CityID int
	}
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

