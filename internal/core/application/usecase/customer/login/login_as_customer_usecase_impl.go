package usecase

import (
	"context"

	entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
)

type LoginAsCustomerUseCase struct {
	ctx context.Context
	customerPGSQLRepository *entity.CustomerPGSQLRepository
}

func NewLoginAsCustomerUseCase(
	ctx context.Context,
	customerPGSQLRepository *entity.CustomerPGSQLRepository,
) *LoginAsCustomerUseCase {
	return &LoginAsCustomerUseCase{
		ctx:          					 ctx,
		customerPGSQLRepository: customerPGSQLRepository,
	}
}

func (u *LoginAsCustomerUseCase) Execute(command *loginAsCustomerCommand) (*entity.Customer, error) {
	var output *entity.Customer

	output, err := u.customerPGSQLRepository.Login(command.Email, command.Password)
	if err != nil {
		return output, err
	}

	return output, err
}
