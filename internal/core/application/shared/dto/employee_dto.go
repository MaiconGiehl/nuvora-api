package dto

type EmployeeOutputDTO struct {
	Cpf int
	Name string
	Email string
	Phone int
}

func NewEmployeeOutputDTO(
	cpf int,
	name string,
	email string,
	phone int,
) *EmployeeOutputDTO {
	return &EmployeeOutputDTO{
		Cpf: cpf, 
		Name: name, 
		Email: email, 
		Phone: phone, 
	}
}