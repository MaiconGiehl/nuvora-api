package entity

import "time"

type Person struct {
	ID 									int
  Cep 								int
	PermissionLevel 		int
	CustomerID					int
	CompanyID 					int
  CreatedAt 					time.Time
  UpdatedAt 					time.Time
}