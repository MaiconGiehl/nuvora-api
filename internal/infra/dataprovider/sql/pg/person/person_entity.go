package entity

import "database/sql"

type Person struct {
	ID int
  PermissionLevel int
  CustomerID sql.NullInt64
  CompanyID sql.NullInt64
  CityID int
  CreatedAt sql.NullTime
  UpdatedAt sql.NullTime
}
