package dto

type CompanyAccountOutputDTO struct {
	ID int
	Name string
	PermissionLevel int
}

func NewCompanyOutputDTO(
	id int,
	name string,
	permissionLevel int,
) *CompanyAccountOutputDTO {
	return &CompanyAccountOutputDTO{
		ID: id,
		Name: name,
		PermissionLevel: permissionLevel,
	}
}