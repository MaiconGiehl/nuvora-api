package dto

type CustomerAccountOutputDTO struct {
	ID int
	Name string
	PermissionLevel int
	TicketsLeft int
}

func NewCustomerAccountOutputDTO(
	id int,
	name string,
	permissionLevel int,
	ticketsLeft int,
) *CustomerAccountOutputDTO {
	return &CustomerAccountOutputDTO{
		ID: id,
		Name: name,
		PermissionLevel: permissionLevel,
		TicketsLeft: ticketsLeft,
	}
}