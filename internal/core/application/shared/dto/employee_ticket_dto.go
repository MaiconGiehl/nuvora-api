package dto

import "time"

type EmployeeTicket struct {
	ID int
	Status string
	Account struct {
		Email string
	}
	Travel struct {
		Price float64
		Account struct {
			SocialReason string
		}
		Departure struct {
			Time time.Time
			City string
		}
		Arrival struct {
			Time time.Time
			City string
		}
	}
	CreatedAt time.Time
}

func NewEmployeeTicket(
	id int,
	status string,
	accountEmail string,
	price float64,
	socialReason string,
	departureTime time.Time,
	departureCity string,
	arrivalTime time.Time,
	arrivalCity string,
	createdAt time.Time,
) *EmployeeTicket {
	return &EmployeeTicket{
		ID: id,
		Status: status,
		Account: struct{Email string}{Email: accountEmail},
		Travel: struct{Price float64; Account struct{SocialReason string}; Departure struct{Time time.Time; City string}; Arrival struct{Time time.Time; City string}}{
			Price: price,
			Account: struct{SocialReason string}{
				SocialReason: socialReason,
			},
			Departure: struct{Time time.Time; City string}{
				Time: departureTime,
				City: departureCity,
			},
			Arrival: struct{Time time.Time; City string}{
				Time: arrivalTime,
				City: arrivalCity,
			},
		},
		CreatedAt: createdAt,
	}
}
