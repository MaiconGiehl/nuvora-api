package usecase

import (
	"context"

	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
)

type LoginAsCompanyUseCase struct {
	ctx context.Context
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository
	personPGSQLRepository *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
}

func NewLoginAsCompanyUseCase(
	ctx context.Context,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
) *LoginAsCompanyUseCase {
	return &LoginAsCompanyUseCase{
		ctx:          					 ctx,
		companyPGSQLRepository: companyPGSQLRepository,
		personPGSQLRepository: personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
	}
}

func (u *LoginAsCompanyUseCase) Execute(command *loginAsCompany) (*company_entity.Company, error) {
	var output *company_entity.Company

	output, err := u.companyPGSQLRepository.Login(command.Email, command.Password)
	if err != nil {
		return output, err
	}

	return output, err
}
