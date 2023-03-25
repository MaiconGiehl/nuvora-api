package entity

import "time"

type Person struct {
	ID 						int
  Cep 					int
  UserID 				int
  CompanyID 		int
  CreatedAt 		time.Time
  UpdatedAt 		time.Time
}