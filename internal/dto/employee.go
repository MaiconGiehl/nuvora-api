package dto

import "database/sql"

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
			Time 									sql.NullTime
			City 									string
		}
		Arrival struct {
			Time 									sql.NullTime
			City 									string
		}
	}
	Status 										string
}