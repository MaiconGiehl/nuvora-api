package dto

import "time"

type CustomerInputDTO struct {
	Name 							string
	Cpf 							int
	Phone 						int
	BirthDate 				time.Time
	CompanyID  				int
}

type CustomerOutputDTO struct {
	ID 								int
	Name 							string
	Phone 						int
	CompanyID    			int
}