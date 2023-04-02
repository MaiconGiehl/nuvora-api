package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
)


type GetAllEmployeesTicketsUseCase struct {
	CompanyRepository database.CompanyRepository
}

func NewGetAllEmployeesTicketsUseCase(
	CompanyRepository database.CompanyRepository,
) *GetAllEmployeesTicketsUseCase {
	return &GetAllEmployeesTicketsUseCase{
		CompanyRepository: CompanyRepository,
	}
}

func (c *GetAllEmployeesTicketsUseCase) Execute(id int) (*[]dto.EmployeesTicketsOutputDTO, error) {
	entity := entity.Company{
		ID:								id,
	}

	output, err := c.CompanyRepository.GetLastMonthTickets(&entity)
	if err != nil {
		return output, err
	}

	return output, nil
}
