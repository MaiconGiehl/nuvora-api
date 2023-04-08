package entity

import "database/sql"

type Person struct {
	ID int
  PermissionLevel int
  CustomerID int
  CompanyID int
  CityID int
  CreatedAt sql.NullTime
  UpdatedAt sql.NullTime
}
