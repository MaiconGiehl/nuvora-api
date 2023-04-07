package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
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
	command *getEmployeesCommand,
) (*[]dto.EmployeeOutputDTO, error) {
	var output []dto.EmployeeOutputDTO

	employees, err := u.customerPGSQLRepository.GetCustomersByCompanyID(command.companyId)
	if err != nil {
		return &output, err
	}

	for _, employee := range *employees {
		output = append(output, dto.EmployeeOutputDTO{
			Name: employee.Name,
		})
	}

	return &output, nil
}