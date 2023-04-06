package usecase

import (
	"github.com/maicongiehl/nuvora-api/internal/core/application/usecase/shared/dto"
)

type LoginAsCustomerUseCaseInterface interface {
	Execute(command *loginAsCustomerCommand) (*dto.CustomerAccountOutputDTO, error)
}