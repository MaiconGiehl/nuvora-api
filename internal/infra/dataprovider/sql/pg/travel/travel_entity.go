package entity

import (
	"database/sql"
)

type Travel struct {
	ID int
	Email sql.NullString
	LastAccess sql.NullTime
	TicketsLeft sql.NullInt16
	PermissionLevel sql.NullInt16
	Company sql.NullString
	City sql.NullString
	Name sql.NullString
}
