package dto

import (
	"time"
)

type CustomerAccountOutputDTO struct {
	ID              int
	Name            string
	PermissionLevel int
	TicketsLeft     int64
}

func NewCustomerAccountOutputDTO(
	id int,
	name string,
	permissionLevel int,
	ticketsLeft int64,
) *CustomerAccountOutputDTO {
	return &CustomerAccountOutputDTO{
		ID:              id,
		Name:            name,
		PermissionLevel: permissionLevel,
		TicketsLeft:     ticketsLeft,
	}
}

type CustomerAccountInputDTO struct {
	Email        string
	Password     string
	DailyTickets int
	CityID       int
	Cpf          int
	Name         string
	Phone        int
	BirthDate    time.Time
	CompanyID    int
}
