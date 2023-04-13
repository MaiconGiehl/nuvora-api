package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type GetLastPurchasesUsecase struct {
	ctx                   context.Context
	ticketPGSQLRepository *entity.TicketPGSQLRepository
}

func NewGetLastPurchasesUsecase(
	ctx context.Context,
	ticketPGSQLRepository *entity.TicketPGSQLRepository,
) *GetLastPurchasesUsecase {
	return &GetLastPurchasesUsecase{
		ctx:                   ctx,
		ticketPGSQLRepository: ticketPGSQLRepository,
	}
}

func (u *GetLastPurchasesUsecase) Execute(
	command *getLastPurchasesCommand,
) (*[]dto.TicketOutputDTO, error) {
	var output []dto.TicketOutputDTO

	lastPurchases, err := u.ticketPGSQLRepository.GetLastPurchases(command.accountId)
	if err != nil {
		return &output, err
	}

	for purchase := range *lastPurchases {
		output = append(output, dto.TicketOutputDTO{
			ID: purchase,
		})
	}

	return &output, nil
}
