package dto

type CompanyAccountOutputDTO struct {
	ID int
	Email string
	PermissionLevel int
	CityName string
	Cnpj int
	SocialReason string
	FantasyName string
	Phone int
	CompanyTypeId int
	AccessToken string
}

func NewCompanyOutputDTO(
	id int,
	email string,
	permissionLevel int,
	cityName string,
	cnpj int,
	socialReason string,
	fantasyName string,
	phone int,
	companyTypeId int,
) *CompanyAccountOutputDTO {
	return &CompanyAccountOutputDTO{
		ID: id,
		Email: email,
		PermissionLevel: permissionLevel,
		CityName: cityName,
		Cnpj: cnpj,
		SocialReason: socialReason,
		FantasyName: fantasyName,
		Phone: phone,
		CompanyTypeId: companyTypeId,
	}
}

func (dto *CompanyAccountOutputDTO) SetAccessToken(
	accessToken string,
) *CompanyAccountOutputDTO {
	dto.AccessToken = accessToken
	return dto
}
