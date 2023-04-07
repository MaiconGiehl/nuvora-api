package entity

import "time"

type Ticket struct {
	ID int
}

type EmployeeTravelTicket struct {
	Name string
	Email string
	Price float64
	DepartureTime time.Time
	DepartureCity string
	ArrivalTime time.Time
	ArrivalCity string
}