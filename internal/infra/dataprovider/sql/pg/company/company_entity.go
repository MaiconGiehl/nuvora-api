package entity

import (
	"database/sql"
	"time"
)

type Company struct {
	ID            int
	Cnpj          int
	SocialReason  string
	FantasyName   sql.NullString
	Phone         sql.NullInt64
	CompanyTypeID int
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
}
