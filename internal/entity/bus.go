package entity

import (
	"time"
)

type Bus struct {
	Id 									int
	Number 							int
	MaxPassengers 			int
	CompanyID 					int
	CreatedAt 					time.Time
	UpdatedAt 					time.Time
}