package usecase

type GetEmployeesUseCaseInterface interface {
	Execute(command *getEmployeesCommand) error
}