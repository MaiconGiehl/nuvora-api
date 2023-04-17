package usecase

type getPossibleTravelsCommand struct {
	accountID int
}

func With(
	accountID int,
) *getPossibleTravelsCommand {
	return &getPossibleTravelsCommand{
		accountID: accountID,
	}
}
