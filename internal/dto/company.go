package dto

type CompanyInputDTO struct {
  Phone 					int
  Cnpj 						int
  SocialReason 		string
  FantasyName 		string
  CompanyTypeId 	int
}

type CompanyOutputDTO struct {
	ID int
  Phone int
  FantasyName string
  CompanyType string
}