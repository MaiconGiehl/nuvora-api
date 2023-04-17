package usecase

import (
	"context"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type GetPossibleTravelsUseCase struct {
	ctx                     context.Context
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
	companyPGSQLRepository  *company_entity.CompanyPGSQLRepository
	personPGSQLRepository   *person_entity.PersonPGSQLRepository
	accountPGSQLRepository  *account_entity.AccountPGSQLRepository
	travelPGSQLRepository   *travel_entity.TravelPGSQLRepository
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
		ctx:                     ctx,
		customerPGSQLRepository: customerPGSQLRepository,
		companyPGSQLRepository:  companyPGSQLRepository,
		personPGSQLRepository:   personPGSQLRepository,
		accountPGSQLRepository:  accountPGSQLRepository,
		travelPGSQLRepository:   travelPGSQLRepository,
	}
}

func (u *GetPossibleTravelsUseCase) Execute(
	command *getPossibleTravelsCommand,
) (*[]dto.TravelOutputDTO, error) {
	var output []dto.TravelOutputDTO
	
	customerAccount, err := u.accountPGSQLRepository.FindAccountByID(command.accountID)
	if err != nil {
		return &output, err
	}

	customerPerson, err := u.personPGSQLRepository.FindPersonByID(customerAccount.ID)
	if err != nil {
		return &output, err
	}

	customer, err := u.customerPGSQLRepository.FindCustomerByID(int(customerPerson.CustomerID.Int64))
	if err != nil {
		return &output, err
	}

	companyPerson, err := u.personPGSQLRepository.FindPersonByID(customer.CompanyID)
	if err != nil {
		return &output, err
	}

	possibleTravels, err := u.travelPGSQLRepository.FindTravelsByCities(customerPerson.CityID, companyPerson.CityID)
	if err != nil {
		return &output, err
	}

	if len(*possibleTravels) < 1 {
		return &output, errors.New("no travel avaiable")
	}


	for _, travel := range *possibleTravels {
		output = append(output, *dto.NewTravelOutputDTO(
			travel.ID, 
			travel.Price,
			travel.AccountID,
			travel.Departure.Time,
			travel.Departure.CityID,
			travel.Arrival.Time,
			travel.Arrival.CityID,
		))
	}

	return &output, nil
}
