package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	city_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/city"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
)

type DeleteEmployeeUseCase struct {
	ctx context.Context
	logger logger.Logger
	cityPGSQLRepository *city_entity.CityPGSQLRepository
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
	personPGSQLRepository *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
}

func NewDeleteEmployeeUseCase(
	ctx context.Context,
	logger logger.Logger,
	cityPGSQLRepository *city_entity.CityPGSQLRepository,
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
) *DeleteEmployeeUseCase {
	return &DeleteEmployeeUseCase{
		ctx: ctx,
		logger: logger,
		cityPGSQLRepository: cityPGSQLRepository,
		customerPGSQLRepository: customerPGSQLRepository,
		personPGSQLRepository: personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
	}
}

func (u *DeleteEmployeeUseCase) Execute(
	command *deleteEmployeeCommand,
) error {
	customerAccount, err := u.accountPGSQLRepository.FindAccountByID(command.employeeID)
	if err != nil {
		u.logger.Errorf("DeleteEmployeeUseCase.Execute: Unable to get account, %s", err.Error())
		return err
	}

	
	customerPerson, err := u.personPGSQLRepository.FindPersonByID(customerAccount.PersonID)
	if err != nil {
		u.logger.Errorf("DeleteEmployeeUseCase.Execute: Unable to get person, %s", err.Error())
		return err
	}
	
	
	customer, err := u.customerPGSQLRepository.FindCustomerByID(int(customerPerson.CustomerID.Int64))
	if err != nil {
		u.logger.Errorf("DeleteEmployeeUseCase.Execute: Unable to get customer, %s", err.Error())
		return err
	}
	
	err = u.customerPGSQLRepository.DeleteCustomerByID(customer.ID, command.companyID)
	if err != nil {
		return err
	}
	
	err = u.personPGSQLRepository.DeletePersonByID(customerAccount.PersonID)
	if err != nil {
		return err
	}

	err = u.accountPGSQLRepository.DeleteAccountByID(customerAccount.ID)

	return err
}
