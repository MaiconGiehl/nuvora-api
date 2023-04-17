package usecase

type getEmployeesCommand struct {
	companyId int
}

func With(
	companyId int,
) *getEmployeesCommand {
	return &getEmployeesCommand{
		companyId: companyId,
	}
}
