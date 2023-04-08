package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type BuyTicketUsecase struct {
	ctx context.Context
	ticketPGSQLRepository *entity.TicketPGSQLRepository
}

func NewBuyTicketUsecase(
	ctx context.Context,
	ticketPGSQLRepository *entity.TicketPGSQLRepository,
) *BuyTicketUsecase {
	return &BuyTicketUsecase{
		ctx: ctx,
		ticketPGSQLRepository: ticketPGSQLRepository,
	}	
}

func (u *BuyTicketUsecase) Execute(
	command *buyTicketCommand,
) (*[]dto.TicketOutputDTO, error) {
	var output []dto.TicketOutputDTO

	err := u.ticketPGSQLRepository.CreateTicket()
	if err != nil {
		return &output, err
	}

	return &output, nil
}