package usecase

type CreateEmployeeUseCaseInterface interface {
	Execute(command *createEmployeeCommand) error
}
