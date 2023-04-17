package usecase

type getPurchasesCommand struct {
	accountId int
}

func With(
	accountId int,
) *getPurchasesCommand {
	return &getPurchasesCommand{
		accountId: accountId,
	}
}