package entity

import "time"

type Person struct {
	ID 									int
  CityID 							int
	PermissionLevel 		int
	CustomerID					int
	CompanyID 					int
  CreatedAt 					time.Time
  UpdatedAt 					time.Time
}