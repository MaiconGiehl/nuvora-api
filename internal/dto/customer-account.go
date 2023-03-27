package dto

type CustomerAccountInputDTO struct {
  Username     	string
  Email   			string
  Password 			string
  Person  			PersonCustomerInputDTO
  DailyTickets 	int
}

type CustomerAccountOutputDTO struct {
	ID 								int
  Username     			string
  PermissionLevel  	int
	City 							string
	TicketsLeft 			int
  DailyTickets 			int
}

type LoginCustomerInputDTO struct {
	Email string
	Password string
}