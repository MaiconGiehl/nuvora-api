package usecase

import (
	"context"

	entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type GetLastPurchases struct {
	ctx context.Context
	ticketPGSQLRepository *entity.TicketPGSQLRepository
}

func NewGetLastPurchases(
	ctx context.Context,
	ticketPGSQLRepository *entity.TicketPGSQLRepository,
) *GetLastPurchases {
	return &GetLastPurchases{
		ctx: ctx,
		ticketPGSQLRepository: ticketPGSQLRepository,
	}	
}

func (u *GetLastPurchases) Execute(
	command *getLastPurchasesCommand,
) (*[]entity.Ticket, error) {
	var output []entity.Ticket

	return &output, nil
}