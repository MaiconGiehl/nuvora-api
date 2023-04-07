package usecase

import "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"

type GetEmployeesUseCaseInterface interface {
	Execute(command *getEmployeesCommand) (*[]dto.EmployeeOutputDTO, error)
}