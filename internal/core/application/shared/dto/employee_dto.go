package dto

type EmployeeOutputDTO struct {
	Cpf int
	Name string
	Phone int
}

func NewEmployeeOutputDTO(
	cpf int,
	name string,
	phone int,
) *EmployeeOutputDTO {
	return &EmployeeOutputDTO{
		Cpf: cpf, 
		Name: name, 
		Phone: phone, 
	}
}
