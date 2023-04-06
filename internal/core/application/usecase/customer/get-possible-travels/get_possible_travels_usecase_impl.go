package usecase

import (
	"context"

	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type GetPossibleTravelsUseCase struct {
	ctx context.Context
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository
	personPGSQLRepository *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository
}

func NewGetPossibleTravelsUseCase(
	ctx context.Context,
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository,
) *GetPossibleTravelsUseCase {
	return &GetPossibleTravelsUseCase{
		ctx: ctx,
		customerPGSQLRepository: customerPGSQLRepository,
		companyPGSQLRepository: companyPGSQLRepository,
		personPGSQLRepository: personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
		travelPGSQLRepository: travelPGSQLRepository,
	}	
}

func (u *GetPossibleTravelsUseCase) Execute(
	command *getPossibleTravelsCommand,
) (*[]travel_entity.Travel, error) {
	var output *[]travel_entity.Travel
	
	customerAccount, err := u.accountPGSQLRepository.GetAccount(command.accountID)
	if err != nil {
		return output, err
	}

	customerPerson, err := u.personPGSQLRepository.GetPersonByAccount(customerAccount.ID)
	if err != nil {
		return output, err
	}

	customer, err := u.customerPGSQLRepository.GetCustomerByPerson(customerPerson.ID)
	if err != nil {
		return output, err
	}

	companyPerson, err := u.personPGSQLRepository.GetPersonByCompany(customer.CompanyID)
	if err != nil {
		return output, err
	}

	output, err = u.travelPGSQLRepository.GetTravelsByCities(customerPerson.CityID, companyPerson.CityID)
	if err != nil {
		return output, err
	}

	return output, nil
}