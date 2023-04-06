package usecase

import (
	"context"

	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type GetEmployeesTicketsUseCase struct {
	ctx context.Context
	customerPGSQLRepository *ticket_entity.TicketPGSQLRepository
}

func NewGetEmployeesUseCase(
	ctx context.Context,
	customerPGSQLRepository *ticket_entity.TicketPGSQLRepository,
) *GetEmployeesTicketsUseCase {
	return &GetEmployeesTicketsUseCase{
		ctx: ctx,
		customerPGSQLRepository: customerPGSQLRepository, 
	}	
}

func (u *GetEmployeesTicketsUseCase) Execute(
	command *GetEmployeesTicketsUseCase,
) error {

	return nil
}