package entity

import (
	"time"
)

type Bus struct {
	Id 									int
	Number 							int
	MaxPassengers 			int
	CreatedIn 					time.Time
	UpdatedIn 					time.Time
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

// func (e *Bus) isValid() bool {
	
// }