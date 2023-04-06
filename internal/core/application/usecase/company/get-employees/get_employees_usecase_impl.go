package usecase

import (
	"context"

	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
)

type GetEmployeesUseCase struct {
	ctx context.Context
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
}

func NewGetEmployeesUseCase(
	ctx context.Context,
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository,
) *GetEmployeesUseCase {
	return &GetEmployeesUseCase{
		ctx: ctx,
		customerPGSQLRepository: customerPGSQLRepository, 
	}	
}

func (u *GetEmployeesUseCase) Execute(
	command *GetEmployeesUseCase,
) error {

	return nil
}