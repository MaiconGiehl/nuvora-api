package usecase

import (
	"context"

	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
)

type CreateEmployeeUseCase struct {
	ctx context.Context
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
	personPGSQLRepository *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
}

func NewCreateEmployeeUseCase(
	ctx context.Context,
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
) *CreateEmployeeUseCase {
	return &CreateEmployeeUseCase{
		ctx: ctx,
		customerPGSQLRepository: customerPGSQLRepository, 
		personPGSQLRepository: personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository, 
	}	
}

func (u *CreateEmployeeUseCase) Execute(
	command *createEmployeeCommand,
) error {

	return nil
}