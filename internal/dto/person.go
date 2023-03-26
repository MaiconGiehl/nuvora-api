package dto

type PersonCustomerInputDTO struct {
  Cep 								int
	PermissionLevel 		int
	Customer					  CustomerInputDTO
}

type PersonCompanyInputDTO struct {
  Cep 								int
	PermissionLevel 		int
	Company 					  CompanyInputDTO
}

type PersonCustomerOutputDTO struct {
	ID 									int
  Cep 								int
	PermissionLevel 		int
	Customer					  CustomerOutputDTO
}

type PersonCompanyOutputDTO struct {
	ID 									int
  Cep 								int
	PermissionLevel 		int
	Customer					  CompanyOutputDTO
}