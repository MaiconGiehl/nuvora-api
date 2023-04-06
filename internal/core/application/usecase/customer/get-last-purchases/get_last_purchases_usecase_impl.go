package usecase

import (
	"context"

	entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type GetLastPurchasesUsecase struct {
	ctx context.Context
	ticketPGSQLRepository *entity.TicketPGSQLRepository
}

func NewGetLastPurchasesUsecase(
	ctx context.Context,
	ticketPGSQLRepository *entity.TicketPGSQLRepository,
) *GetLastPurchasesUsecase {
	return &GetLastPurchasesUsecase{
		ctx: ctx,
		ticketPGSQLRepository: ticketPGSQLRepository,
	}	
}

func (u *GetLastPurchasesUsecase) Execute(
	command *getLastPurchasesCommand,
) (*[]entity.Ticket, error) {

	output, err := u.ticketPGSQLRepository.GetLastPurchases(command.accountId)
	if err != nil {
		return output, err
	}

	return output, nil
}