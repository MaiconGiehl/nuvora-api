package usecase

import "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"

type LoginAsCompanyUseCaseInterface interface {
	Execute(command *loginAsCompany) (*dto.CompanyAccountOutputDTO, error)
}