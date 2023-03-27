package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type GetAllCompanyUseCase struct {
	CompanyRepository database.CompanyRepository
}

func NewGetAllCompanyUseCase(
	CompanyRepository database.CompanyRepository,
) *GetAllCompanyUseCase {
	return &GetAllCompanyUseCase{
		CompanyRepository: CompanyRepository,
	}
}

func (c *GetAllCompanyUseCase) Execute() (*[]dto.CompanyOutputDTO, error) {
	output, err := c.CompanyRepository.GetAll()
	if err != nil {
		return output, err
	}

	return output, nil
}
