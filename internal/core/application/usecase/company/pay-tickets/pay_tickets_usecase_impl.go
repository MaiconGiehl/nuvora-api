package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
)

type PayTicketsUseCase struct {
	ctx context.Context
	logger logger.Logger
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository
	personPGSQLRepository *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository
}

func NewPayTicketsUseCase(
	ctx context.Context,
	logger logger.Logger,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository,
) *PayTicketsUseCase {
	return &PayTicketsUseCase{
		ctx: ctx,
		logger: logger,
		companyPGSQLRepository: companyPGSQLRepository,
		personPGSQLRepository: personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
		ticketPGSQLRepository: ticketPGSQLRepository,
	}
}

func (u *PayTicketsUseCase) Execute(
	command *payTicketsCommand) (string, error) {

 	companyAccount, err :=	u.accountPGSQLRepository.FindAccountByID(command.companyAccountID)
	if err != nil {
		return "invalid account", err
	}

	person, err := u.personPGSQLRepository.FindPersonByID(companyAccount.PersonID)
	if err != nil {
		return "", err
	}

	company, err := u.companyPGSQLRepository.FindCompanyByID(int(person.CompanyID.Int64))
	if err != nil {
		return "", err
	}

	rowsAffected, err := u.ticketPGSQLRepository.UpdateTicketsStatusByCompanyID(company.ID)
	if err != nil {
		return "", err
	}

	return rowsAffected, nil 
}
