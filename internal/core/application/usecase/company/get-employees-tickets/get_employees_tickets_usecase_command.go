package usecase

type getEmployeesTicketsCommand struct {
	companyId int
}

func With(
	companyId int,
) *getEmployeesTicketsCommand {
	return &getEmployeesTicketsCommand{
		companyId: companyId,
	}
}