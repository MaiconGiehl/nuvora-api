package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type GetCustomerAccountUseCase struct {
	AccountRepository database.AccountRepository
}

func NewGetCustomerAccountUseCase(
	AccountRepository database.AccountRepository,
) *GetCustomerAccountUseCase {
	return &GetCustomerAccountUseCase{
		AccountRepository: AccountRepository,
	}
}

func (c *GetCustomerAccountUseCase) Execute(input *dto.LoginCustomerInputDTO) (*dto.CustomerAccountOutputDTO, error) {
	entity := entity.Account{
		Email: input.Email,
		Password: input.Password,
	}

	output, err := c.AccountRepository.GetCustomerAccount(&entity)
	if err != nil {
		return output, err
	}

	return output, nil
}