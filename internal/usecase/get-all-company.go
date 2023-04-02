package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
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
