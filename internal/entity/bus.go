package entity

import (
	"time"
)

type Bus struct {
	Id int
	Number int
	MaxPassengers int
	CreatedIn time.Time
	UpdatedIn time.Time
}

func NewBus (id int, number int, maxPassengers int) (*Bus, error) {
	bus := &Bus{
		Id: id,
		Number: number,
		MaxPassengers: maxPassengers,
	}
	// err := bus.IsValid()
	// busRepository.Save(*bus)

	return bus, nil
}

// func (c *Bus) IsValid() error {
// 	if c.ID == "" {
// 		return errors.New("invalid id")
// 	}
// 	return nil
// }