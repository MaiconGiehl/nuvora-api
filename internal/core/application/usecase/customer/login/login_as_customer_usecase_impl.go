package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/usecase/shared/dto"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
)

type LoginAsCustomerUseCase struct {
	ctx context.Context
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
	personPGSQLRepository *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
}

func NewLoginAsCustomerUseCase(
	ctx context.Context,
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
) *LoginAsCustomerUseCase {
	return &LoginAsCustomerUseCase{
		ctx:          					 ctx,
		customerPGSQLRepository: customerPGSQLRepository,
		personPGSQLRepository: personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
	}
}

func (u *LoginAsCustomerUseCase) Execute(command *loginAsCustomerCommand) (*dto.CustomerAccountOutputDTO, error) {
	var output *dto.CustomerAccountOutputDTO

	customerAccount, err := u.accountPGSQLRepository.LoginAsCustomer(command.Email, command.Password)
	if err != nil {
		return output, err
	}

	customerPerson, err := u.personPGSQLRepository.GetPersonByAccount(customerAccount.PersonID)
	if err != nil {
		return output, err
	}

	customer, err := u.customerPGSQLRepository.GetCustomerByPerson(customerPerson.ID)
	if err != nil {
		return output, err
	}

	output = dto.NewCustomerAccountOutputDTO(
		customerAccount.ID,
		customer.Name,
		customerPerson.PermissionLevel,
		customerAccount.TicketsLeft,
	)

	return output, err
}
