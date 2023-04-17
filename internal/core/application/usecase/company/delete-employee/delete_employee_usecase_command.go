package usecase

type deleteEmployeeCommand struct {
	companyID int
	employeeID int
}

func With(
	companyID int,
	employeeID int,
) *deleteEmployeeCommand {
	return &deleteEmployeeCommand{
		 companyID: companyID,
		 employeeID: employeeID,
	}
}