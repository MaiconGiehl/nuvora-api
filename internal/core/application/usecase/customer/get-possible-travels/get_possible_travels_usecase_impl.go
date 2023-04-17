package usecase

import (
	"context"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	bus_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/bus"
	city_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/city"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type GetPossibleTravelsUseCase struct {
	ctx context.Context
	busPGSQLRepository *bus_entity.BusPGSQLRepository
	cityPGSQLRepository *city_entity.CityPGSQLRepository
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository
	companyPGSQLRepository  *company_entity.CompanyPGSQLRepository
	personPGSQLRepository   *person_entity.PersonPGSQLRepository
	accountPGSQLRepository  *account_entity.AccountPGSQLRepository
	travelPGSQLRepository   *travel_entity.TravelPGSQLRepository
}

func NewGetPossibleTravelsUseCase(
	ctx context.Context,
	busPGSQLRepository *bus_entity.BusPGSQLRepository,
	cityPGSQLRepository *city_entity.CityPGSQLRepository,
	customerPGSQLRepository *customer_entity.CustomerPGSQLRepository,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository,
) *GetPossibleTravelsUseCase {
	return &GetPossibleTravelsUseCase{
		ctx: ctx,
		busPGSQLRepository: busPGSQLRepository,
		cityPGSQLRepository: cityPGSQLRepository,
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

	company, err := u.companyPGSQLRepository.FindCompanyByID(customer.CompanyID)
	if err != nil {
		return &output, err
	}

	companyPerson, _ := u.personPGSQLRepository.FindPersonByCompanyID(company.ID)

	possibleTravels, err := u.travelPGSQLRepository.FindTravelsByCities(customerPerson.CityID, companyPerson.CityID)
	if err != nil {
		return &output, err
	}

	if len(*possibleTravels) < 1 {
		return &output, errors.New("no travel avaiable")
	}

	allBus, _ := u.busPGSQLRepository.FindAll()
	getBus := func(id int) *bus_entity.Bus {
		for _, bus := range allBus {
			if bus.ID == id {
				return bus
			}
		}
		return &bus_entity.Bus{}
	}


	allCities, _ := u.cityPGSQLRepository.FindAll()
	getCities := func(id int) city_entity.City {
		for _, city := range allCities {
			if city.ID == id {
				return *city
			}
		}
		return city_entity.City{}
	}

	for _, travel := range *possibleTravels {
		bus := getBus(travel.BusID)
		departureCity := getCities(travel.Departure.CityID)
		arrivalCity := getCities(travel.Arrival.CityID)
		output = append(output, *dto.NewTravelOutputDTO(
			travel.ID,
			travel.Price,
			bus.ID,
			bus.Number,
			bus.MaxPassengers,
			travel.AccountID,
			travel.Departure.Time,
			departureCity.Name,
			travel.Arrival.Time,
			arrivalCity.Name,
		))
	}

	return &output, nil
}
