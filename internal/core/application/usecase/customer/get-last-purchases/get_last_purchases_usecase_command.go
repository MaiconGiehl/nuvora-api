package usecase

type getLastPurchasesCommand struct {
	accountId int
}

func With(
	accountId int,
) *getLastPurchasesCommand {
	return &getLastPurchasesCommand{
		accountId: accountId,
	}
}
