package entity

import "time"

type Ticket struct {
	ID 									int
  AccountID						int
  Status 							int
  TravelID 						int
  CreatedAt 					time.Time
  UpdatedAt 					time.Time
}