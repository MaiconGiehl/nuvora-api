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

func NewBus (id int, number int, maxPassengers int) (*Bus, error) {
	bus := &Bus{
		Id: 							id,
		Number: 					number,
		MaxPassengers: 		maxPassengers,
	}

	// if !bus.isValid() {
	// 	return bus, errors.New("Invalid input")
	// }

	return bus, nil
}