package usecase

import (
	"context"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
)

type LoginAsCustomerUseCase struct {
	ctx                     context.Context
	logger                  logger.Logger
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
	personPGSQLRepository   *person_entity.PersonPGSQLRepository
	accountPGSQLRepository  *account_entity.AccountPGSQLRepository
}

func NewLoginAsCustomerUseCase(
	ctx context.Context,
	logger logger.Logger,
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
) *LoginAsCustomerUseCase {
	return &LoginAsCustomerUseCase{
		ctx:                     ctx,
		logger:                  logger,
		customerPGSQLRepository: customerPGSQLRepository,
		personPGSQLRepository:   personPGSQLRepository,
		accountPGSQLRepository:  accountPGSQLRepository,
	}
}

func (u *LoginAsCustomerUseCase) Execute(command *loginAsCustomerCommand) (*dto.CustomerAccountOutputDTO, error) {
	var output *dto.CustomerAccountOutputDTO

	customerAccount, err := u.accountPGSQLRepository.Login(command.Email, command.Password)
	if err != nil {
		u.logger.Errorf("LoginAsCustomerUseCase.Execute: Unable to login in account, %s", err.Error())
		return output, err
	}

	customerPerson, _ := u.personPGSQLRepository.FindPersonByID(customerAccount.PersonID)

	if customerPerson.CompanyID.Valid {
		err = errors.New("invalid credentials")
		u.logger.Errorf("LoginAsCustomerUseCase.Execute: Unable to login, %s", err.Error())
		return output, err
	}

	customer, _ := u.customerPGSQLRepository.FindCustomerByID(int(customerPerson.CustomerID.Int64))

	output = dto.NewCustomerAccountOutputDTO(
		customerAccount.ID,
		customerAccount.Email,
		int(customerAccount.TicketsLeft.Int64),
		customerPerson.PermissionLevel,
		customer.Cpf,
		customer.Name,
		customer.Phone,
	)

	return output, err
}
