package entity

import (
	"database/sql"
	"time"
)

type Account struct {
	ID           int
	Email        string
	Password     string
	PersonID     int
	LastAccess   time.Time
	TicketsLeft  sql.NullInt64
	DailyTickets int
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
}
