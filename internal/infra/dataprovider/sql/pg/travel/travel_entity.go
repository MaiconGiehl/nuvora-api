package entity

import (
	"database/sql"
	"time"
)

type Travel struct {
	ID        int
	Price     float64
	CompanyID string
	Bus       struct {
		Number        int
		MaxPassengers int
	}
	Departure struct {
		Time     time.Time
		CityName string
	}
	Arrival struct {
		Time     time.Time
		CityName string
	}
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
