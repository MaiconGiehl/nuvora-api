package usecase

type DeleteTravelUseCaseInterface interface {
	Execute(command *deleteTravelCommand) error
}