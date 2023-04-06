package usecase

import entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"

type LoginAsCompanyUseCaseInterface interface {
	Execute(command *loginAsCompany) (*entity.Company, error)
}