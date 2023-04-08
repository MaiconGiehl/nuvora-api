package entity

import "database/sql"

type Customer struct {
	ID int
  Cpf int
  Name string
  Phone int 
  Birth_date sql.NullTime
  CompanyID int
  CreatedAt sql.NullTime
  UpdatedAt sql.NullTime 
}
