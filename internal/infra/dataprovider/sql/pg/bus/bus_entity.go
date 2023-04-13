package entity

import (
	"database/sql"
	"time"
)

type Bus struct {
	ID            int
	Number        int
	MaxPassengers int
	CompanyID     int
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
}
