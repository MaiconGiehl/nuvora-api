package usecase

type GetEmployeesTicketsUseCaseInterface interface {
	Execute(command *getEmployeesTicketsCommand) error
}