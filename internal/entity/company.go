package entity

import (
	"time"
)

type Company struct {
	ID int
  Phone int
  Cnpj int
  SocialReason string
  FantasyName string
  CompanyTypeId int
  CreatedAt time.Time
  UpdatedAt time.Time
}