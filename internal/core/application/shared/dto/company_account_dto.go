package dto

type CompanyAccountOutputDTO struct {
	ID              int
	FantasyName     string
	PermissionLevel int
	AccessToken string
}

func NewCompanyOutputDTO(
	id int,
	fantasyName string,
	permissionLevel int,
) *CompanyAccountOutputDTO {
	return &CompanyAccountOutputDTO{
		ID:              id,
		FantasyName:     fantasyName,
		PermissionLevel: permissionLevel,
	}
}
