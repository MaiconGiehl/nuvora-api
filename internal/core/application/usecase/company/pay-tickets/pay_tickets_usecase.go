package usecase

type PayTicketsUseCaseInterface interface {
	Execute(command *payTicketsCommand) (string, error)
}
