package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type GetEmployeesTicketsUseCase struct {
	ctx                    context.Context
	logger                 logger.Logger
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository
	personPGSQLRepository  *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
	ticketPGSQLRepository  *ticket_entity.TicketPGSQLRepository
}

func NewGetEmployeesTicketsUseCase(
	ctx context.Context,
	logger logger.Logger,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository,
) *GetEmployeesTicketsUseCase {
	return &GetEmployeesTicketsUseCase{
		ctx:                    ctx,
		logger:                 logger,
		companyPGSQLRepository: companyPGSQLRepository,
		personPGSQLRepository:  personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
		ticketPGSQLRepository:  ticketPGSQLRepository,
	}
}

func (u *GetEmployeesTicketsUseCase) Execute(
	command *getEmployeesTicketsCommand,
) (*[]dto.EmployeeTicket, error) {
	var output []dto.EmployeeTicket

	companyAccount, err := u.accountPGSQLRepository.FindAccountByID(command.companyId)
	if err != nil {
		return &output, err
	}

	companyPerson, err := u.personPGSQLRepository.FindPersonByID(companyAccount.PersonID)
	if err != nil {
		return &output, err
	}

	company, err := u.companyPGSQLRepository.FindCompanyByID(int(companyPerson.CompanyID.Int64))
	if err != nil {
		return &output, err
	}

	tickets, err := u.ticketPGSQLRepository.GetEmployeesTickets(company.ID)
	if err != nil {
		return &output, err
	}

	for _, ticket := range tickets {
		output = append(output, dto.EmployeeTicket{
			TicketOutputDTO: dto.TicketOutputDTO{
				ID:        ticket.ID,
				StatusID:  ticket.StatusID,
				TravelID:  ticket.TravelID,
				CreatedAt: ticket.CreatedAt,
			},
		})
	}
	return &output, nil
}
