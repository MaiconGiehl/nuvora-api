package dto

import "database/sql"

type CustomerAccountInputDTO struct {
  Username     	string
  Email   			string
  Password 			string
  Person  			PersonCustomerInputDTO
  DailyTickets 	sql.NullInt16
}

type CustomerAccountOutputDTO struct {
	ID 						int
  Username     	string
  Person  			PersonCustomerOutputDTO
	TicketsLeft 	sql.NullInt16
  DailyTickets 	sql.NullInt16
}

type LoginCustomerInputDTO struct {
	Email string
	Password string
}