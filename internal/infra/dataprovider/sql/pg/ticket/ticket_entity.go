package entity

import (
	"database/sql"
	"time"
)

type Ticket struct {
	ID int
	AccountID int
	StatusID int
	TravelID int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}