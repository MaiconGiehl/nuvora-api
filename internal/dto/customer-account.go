package dto

import "time"

type CustomerAccountInputDTO struct {
  Email   			string
  Password 			string
  Person  			struct {
		CityID 							int
		Customer					  struct {
			Name 							string
			Cpf 							int
			Phone 						int
			BirthDate 				time.Time
		}
	}
  DailyTickets 	int
}

type CustomerAccountOutputDTO struct {
	ID 								int
  PermissionLevel  	int
	CityID  					int
	TicketsLeft 			int
}

type LoginCustomerInputDTO struct {
	Email string
	Password string
}