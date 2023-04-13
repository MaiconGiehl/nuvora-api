package usecase

import "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"

type GetEmployeesTicketsUseCaseInterface interface {
	Execute(command *getEmployeesTicketsCommand) (*dto.EmployeeTicket, error)
}
