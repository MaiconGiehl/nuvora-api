package dto

import (
	"time"
)

type EmployeeOutputDTO struct {
	ID 											string
	Name 										string
}

type EmployeesTicketsOutputDTO struct {
	Name 											string
	Cpf 											int
	Travel struct {
		Price 									float64
		Departure struct {
			Time 									time.Time
			City 									string
		}
		Arrival struct {
			Time 									time.Time
			City 									string
		}
	}
	Status 										string
}