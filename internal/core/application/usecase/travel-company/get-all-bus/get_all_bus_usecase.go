package usecase

type GetAllBusUseCaseInterface interface {
	Execute(command *getAllBusCommand) error
}