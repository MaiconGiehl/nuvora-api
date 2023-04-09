package usecase

type buyTicketCommand struct {
	AccountID int
	TravelID int
}

func With(
	accountId int,
	travelId int,
) *buyTicketCommand {
	return &buyTicketCommand{
		AccountID: accountId,
		TravelID: travelId,
	}
}