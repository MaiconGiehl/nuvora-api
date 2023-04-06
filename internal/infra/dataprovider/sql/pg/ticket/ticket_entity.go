package entity

import (
	"database/sql"
)

type Ticket struct {
	ID int
	Email sql.NullString
	LastAccess sql.NullTime
	TicketsLeft sql.NullInt16
	PermissionLevel sql.NullInt16
	Company sql.NullString
	City sql.NullString
	Name sql.NullString
}
