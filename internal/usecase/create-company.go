package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
)


type CreateCompanyUseCase struct {
	CompanyRepository database.CompanyRepository
}

func NewCreateCompanyUseCase(
	CompanyRepository database.CompanyRepository,
) *CreateCompanyUseCase {
	return &CreateCompanyUseCase{
		CompanyRepository: CompanyRepository,
	}
}

func (c *CreateCompanyUseCase) Execute(input *dto.CompanyInputDTO) error {
	entity := entity.Company{
		Phone: 					input.Phone,
  	Cnpj: 					input.Cnpj,
  	SocialReason: 	input.SocialReason,
  	FantasyName: 		input.FantasyName,
  	CompanyTypeId: 	input.CompanyTypeId,
	}

	err := c.CompanyRepository.Save(&entity)
	if err != nil {
		return err
	}

	return nil
}
