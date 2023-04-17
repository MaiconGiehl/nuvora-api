package usecase

type deleteTravelCommand struct {
	travelCompanyID int
	travelID int
}

func With(
	travelCompanyId int,
	travelId int,
	) *deleteTravelCommand {
	return &deleteTravelCommand{
		travelCompanyID: travelCompanyId,
		travelID: travelId,
	}
}