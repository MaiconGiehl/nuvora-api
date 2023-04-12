package usecase

import (
	"context"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
)

type GetEmployeesUseCase struct {
	ctx context.Context
	logger logger.Logger
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository
	personPGSQLRepository *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
}

func NewGetEmployeesUseCase(
	ctx context.Context,
	logger logger.Logger,
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
) *GetEmployeesUseCase {
	return &GetEmployeesUseCase{
		ctx: ctx,
		logger: logger,
		customerPGSQLRepository: customerPGSQLRepository,
		companyPGSQLRepository: companyPGSQLRepository,
		personPGSQLRepository: personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
	}	
}

func (u *GetEmployeesUseCase) Execute(
	command *getEmployeesCommand,
) (*[]dto.EmployeeOutputDTO, error) {
	var output []dto.EmployeeOutputDTO

	companyAccount, err := u.accountPGSQLRepository.FindAccountByID(command.companyId)
	if err != nil {
		u.logger.Errorf("GetEmployeesUseCase.Execute: Unable to find company account, %s", err.Error())
		return &output, err
	}

	companyPerson, err := u.personPGSQLRepository.FindPersonByID(companyAccount.PersonID)
	if err != nil {
		u.logger.Errorf("GetEmployeesUseCase.Execute: Unable to find a person associated with the company account, %s", err.Error())
		return &output, err
	}

	if !companyPerson.CompanyID.Valid && companyPerson.CustomerID.Valid {
		err = errors.New("invalid account")
		u.logger.Warnf("GetEmployeesUseCase.Execute: There was a try to use a customer account to use a company usecase , %s.", err.Error())
		return &output, err
	} 

	company, err := u.companyPGSQLRepository.FindCompanyByID(int(companyPerson.CompanyID.Int64))
	if err != nil {
		u.logger.Errorf("GetEmployeesUseCase.Execute: Unable to find a company associated with the person, %s", err.Error())
		return &output, err
	}

	employees, err := u.customerPGSQLRepository.GetCustomersByCompanyID(company.ID)
	if err != nil {
		u.logger.Errorf("GetEmployeesUseCase.Execute: Unable to get customers associated with the company, %s", err.Error())
		return &output, err
	}

	for _, employee := range *employees {
		output = append(output, dto.EmployeeOutputDTO{
			Name: employee.Name,
		})
	}

	return &output, nil
}