package entity

import (
	"database/sql"
	"time"
)

type Ticket struct {
	ID        int
	AccountID int
	StatusID  int
	TravelID  int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type EmployeeTravelTicket struct {
	Name          string
	Email         string
	Price         float64
	DepartureTime time.Time
	DepartureCity string
	ArrivalTime   time.Time
	ArrivalCity   string
}
