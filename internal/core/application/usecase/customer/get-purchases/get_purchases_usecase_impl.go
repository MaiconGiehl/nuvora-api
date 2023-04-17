package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type GetPurchasesUsecase struct {
	ctx                   context.Context
	ticketPGSQLRepository *entity.TicketPGSQLRepository
}

func NewGetPurchasesUsecase(
	ctx context.Context,
	ticketPGSQLRepository *entity.TicketPGSQLRepository,
) *GetPurchasesUsecase {
	return &GetPurchasesUsecase{
		ctx:                   ctx,
		ticketPGSQLRepository: ticketPGSQLRepository,
	}
}

func (u *GetPurchasesUsecase) Execute(
	command *getPurchasesCommand,
) (*[]dto.TicketOutputDTO, error) {
	var output []dto.TicketOutputDTO

	tickets, err := u.ticketPGSQLRepository.GetPurchases(command.accountId)
	if err != nil {
		return &output, err
	}

	for _, ticket := range *tickets {
		output = append(output, dto.TicketOutputDTO{
			ID:        ticket.ID,
			StatusID:  ticket.StatusID,
			TravelID:  ticket.TravelID,
			CreatedAt: ticket.CreatedAt,
		})
	}

	return &output, nil
}
