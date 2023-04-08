package entity

import (
	"database/sql"
)

type Account struct {
	ID int
  Email string
  Password string
  PersonID int
  LastAccess sql.NullTime
  TicketsLeft int 
  DailyTickets int
  CreatedAt sql.NullTime
  UpdatedAt sql.NullTime
}