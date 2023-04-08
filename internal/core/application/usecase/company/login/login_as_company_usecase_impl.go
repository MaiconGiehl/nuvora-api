package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
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

func (u *LoginAsCompanyUseCase) Execute(
	command *loginAsCompany) (*dto.CompanyAccountOutputDTO, error) {
	var output *dto.CompanyAccountOutputDTO

	companyAccount, err := u.accountPGSQLRepository.LoginAsCompany(command.Email, command.Password)
	if err != nil {
		return output, err
	}

	companyPerson, err := u.personPGSQLRepository.GetPersonByID(companyAccount.PersonID)
	if err != nil {
		return output, err
	}

	company, err := u.companyPGSQLRepository.GetCompanyByID(companyPerson.CompanyID)
	if err != nil {
		return output, err
	}

	output = dto.NewCompanyOutputDTO(
		companyAccount.ID,
		company.FantasyName,
		companyPerson.PermissionLevel,
	)

	return output, err
}
