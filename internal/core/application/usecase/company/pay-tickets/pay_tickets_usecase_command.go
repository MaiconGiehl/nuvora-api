package usecase

type payTicketsCommand struct {
	companyAccountID int
}

func With(
	companyAccountId int,
) *payTicketsCommand {
	return &payTicketsCommand{
		companyAccountID: companyAccountId,
	}
}
