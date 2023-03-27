package entity

import (
	"time"
)

type Bus struct {
	Id 									int
	Number 							int
	MaxPassengers 			int
	CreatedAt 					time.Time
	UpdatedAt 					time.Time
}