package usecase

type getAllBusCommand struct {
	travelCompanyID int
}

func With(
	travelCompanyId int,
) *getAllBusCommand {
	return &getAllBusCommand{
		travelCompanyID: travelCompanyId,
	}
}
