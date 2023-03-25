package dto

import "time"

type CompanyInputDTO struct {
  ID 							int
  Phone 					int
  Cnpj 						int
  SocialReason 		string
  FantasyName 		string
  CompanyTypeId 	int
}

type CompanyOutputDTO struct {
  Cep 						int
  UserID 					int
  CompanyID 			int
  CreatedAt 			time.Time
  UpdatedAt 			time.Time
}