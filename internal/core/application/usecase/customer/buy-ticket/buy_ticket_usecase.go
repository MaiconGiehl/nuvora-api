package usecase

type BuyTicketUseCaseInterface interface {
	Execute(command *buyTicketCommand) error
}