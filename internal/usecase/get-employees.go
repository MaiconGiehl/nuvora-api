package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
)


type GetEmployees struct {
	CompanyRepository database.CompanyRepository
}

func NewGetEmployees(
	CompanyRepository database.CompanyRepository,
) *GetEmployees {
	return &GetEmployees{
		CompanyRepository: CompanyRepository,
	}
}

func (c *GetEmployees) Execute(id int) (*[]dto.EmployeeOutputDTO, error) {
	entity := entity.Company{
		ID:								id,
	}

	output, err := c.CompanyRepository.GetEmployees(&entity)
	if err != nil {
		return output, err
	}

	return output, nil
}