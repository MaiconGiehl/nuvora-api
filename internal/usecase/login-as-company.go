package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
)


type GetCompanyAccountUseCase struct {
	AccountRepository database.AccountRepository
}

func NewGetCompanyAccountUseCase(
	AccountRepository database.AccountRepository,
) *GetCompanyAccountUseCase {
	return &GetCompanyAccountUseCase{
		AccountRepository: AccountRepository,
	}
}

func (c *GetCompanyAccountUseCase) Execute(input *dto.CompanyLoginDTO) (*dto.CompanyAccountOutputDTO, error) {
	entity := entity.Account{
		Email: input.Email,
		Password: input.Password,
	}

	output, err := c.AccountRepository.GetCompanyAccount(&entity)
	if err != nil {
		return output, err
	}

	return output, nil
}
