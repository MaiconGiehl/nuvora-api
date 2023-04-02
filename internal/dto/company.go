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


type CompanyLoginDTO struct {
	Email string
	Password string
}

type CompanyAccountInputDTO struct {
	Email   			string
  Password 			string
  Person  			struct {
		CityID 							int
		Company					  struct {
			FantasyName 			string
			Cnpj 							int
			Phone 						int
		}
	}
}

type CompanyAccountOutputDTO struct {
	ID 								int
  PermissionLevel  	int
	CityID  					int
}