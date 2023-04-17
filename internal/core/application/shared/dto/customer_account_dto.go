package dto

import (
	"time"
)

type CustomerAccountOutputDTO struct {
	ID int
	Email string
	TicketsLeft int
	PermissionLevel int
	Cpf int
	Name string
	Phone int
	AccessToken string
}

func NewCustomerAccountOutputDTO(
	id int,
	email string,
	ticketsLeft int,
	permissionLevel int,
	cpf int,
	name string,
	phone int,
) *CustomerAccountOutputDTO {
	return &CustomerAccountOutputDTO{
		ID: id, 
		Email: email, 
		TicketsLeft: ticketsLeft, 
		PermissionLevel: permissionLevel, 
		Cpf: cpf, 
		Name: name, 
		Phone: phone, 
	}
}

func (dto *CustomerAccountOutputDTO) SetAccessToken(accessToken string) *CustomerAccountOutputDTO {
	dto.AccessToken = accessToken
	return dto
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
