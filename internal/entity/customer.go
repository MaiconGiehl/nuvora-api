package entity

import "time"

type Customer struct {
	Id 								int
	Name 							string
	Cpf 							int
	Phone 						int
	BirthDate					time.Time
	CompanyID 				int
	LastAccess 				time.Time
	CreatedAt 				time.Time
	UpdatedAt 				time.Time
}
