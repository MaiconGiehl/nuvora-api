package usecase

type CreateTravelUseCaseInterface interface {
	Execute(command *createTravelCommand) error
}