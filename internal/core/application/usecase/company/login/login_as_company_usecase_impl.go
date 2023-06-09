package usecase

import (
	"context"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	city_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/city"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
)

type LoginAsCompanyUseCase struct {
	ctx context.Context
	logger logger.Logger
	cityPGSQLRepository *city_entity.CityPGSQLRepository
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository
	personPGSQLRepository  *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
}

func NewLoginAsCompanyUseCase(
	ctx context.Context,
	logger logger.Logger,
	cityPGSQLRepository *city_entity.CityPGSQLRepository,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
) *LoginAsCompanyUseCase {
	return &LoginAsCompanyUseCase{
		ctx: ctx,
		logger: logger,
		cityPGSQLRepository: cityPGSQLRepository,
		companyPGSQLRepository: companyPGSQLRepository,
		personPGSQLRepository:  personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
	}
}

func (u *LoginAsCompanyUseCase) Execute(
	command *loginAsCompany) (*dto.CompanyAccountOutputDTO, error) {
	var output *dto.CompanyAccountOutputDTO

	companyAccount, err := u.accountPGSQLRepository.Login(command.Email, command.Password)
	if err != nil {
		u.logger.Errorf("LoginAsCompanyUseCase.Execute: Unable to login in account, %s", err.Error())
		return output, err
	}

	companyPerson, err := u.personPGSQLRepository.FindPersonByID(companyAccount.PersonID)
	if err != nil {
		u.logger.Errorf("LoginAsCompanyUseCase.Execute: Unable to get person, %s", err.Error())
		return output, err
	}

	if companyPerson.CustomerID.Valid {
		err = errors.New("invalid credentials")
		u.logger.Errorf("LoginAsCustomerUseCase.Execute: Unable to login, %s", err.Error())
		return output, err
	}

	company, err := u.companyPGSQLRepository.FindCompanyByID(int(companyPerson.CompanyID.Int64))
	if err != nil {
		u.logger.Errorf("LoginAsCompanyUseCase.Execute: Unable to get company, %s", err.Error())
		return output, err
	}

	city, err := u.cityPGSQLRepository.FindCityByID(companyPerson.CityID)
	if err != nil {
		u.logger.Errorf("LoginAsCompanyUseCase.Execute: Unable to get city, %s", err.Error())
		return output, err
	}

	output = dto.NewCompanyOutputDTO(
		companyAccount.ID,
		companyAccount.Email,
		companyPerson.PermissionLevel,
		city.Name,
		company.Cnpj,
		company.SocialReason,
		company.FantasyName.String,
		int(company.Phone.Int64),
		company.CompanyTypeID,
	)

	return output, err
}
