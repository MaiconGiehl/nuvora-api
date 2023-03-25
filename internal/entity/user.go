package entity

import "time"

type User struct {
	Id 								int
	Name 							string
	Cpf 							int
	Email							string
	Cep 							int
	LastAccess 				time.Time
	CreatedAt 				time.Time
	UpdatedAt 				time.Time
}