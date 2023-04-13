package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type GetEmployeesTicketsUseCase struct {
	ctx                   context.Context
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository
}

func NewGetEmployeesTicketsUseCase(
	ctx context.Context,
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository,
) *GetEmployeesTicketsUseCase {
	return &GetEmployeesTicketsUseCase{
		ctx:                   ctx,
		ticketPGSQLRepository: ticketPGSQLRepository,
	}
}

func (u *GetEmployeesTicketsUseCase) Execute(
	command *getEmployeesTicketsCommand,
) (*[]dto.EmployeeTicket, error) {
	var output []dto.EmployeeTicket

	employeesTickets, err := u.ticketPGSQLRepository.GetEmployeesTickets(command.companyId)
	if err != nil {
		return &output, err
	}

	for _, empTicket := range *employeesTickets {
		output = append(output, dto.EmployeeTicket{
			Name: empTicket.Name,
		})
	}
	return &output, nil
}
