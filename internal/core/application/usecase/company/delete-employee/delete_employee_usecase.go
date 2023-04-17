package usecase

type DeleteEmployeeUseCaseInterface interface {
	Execute(command *deleteEmployeeCommand) error
}