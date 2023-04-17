package usecase

import (
	"context"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	city_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/city"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

type GetEmployeesTicketsUseCase struct {
	ctx context.Context
	logger logger.Logger
	cityPGSQLRepository *city_entity.CityPGSQLRepository
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository
	personPGSQLRepository  *person_entity.PersonPGSQLRepository
	accountPGSQLRepository *account_entity.AccountPGSQLRepository
	ticketPGSQLRepository  *ticket_entity.TicketPGSQLRepository
}

func NewGetEmployeesTicketsUseCase(
	ctx context.Context,
	logger logger.Logger,
	cityPGSQLRepository *city_entity.CityPGSQLRepository,
	travelPGSQLRepository *travel_entity.TravelPGSQLRepository,
	companyPGSQLRepository *company_entity.CompanyPGSQLRepository,
	personPGSQLRepository *person_entity.PersonPGSQLRepository,
	accountPGSQLRepository *account_entity.AccountPGSQLRepository,
	ticketPGSQLRepository *ticket_entity.TicketPGSQLRepository,
) *GetEmployeesTicketsUseCase {
	return &GetEmployeesTicketsUseCase{
		ctx: ctx,
		logger: logger,
		cityPGSQLRepository: cityPGSQLRepository,
		travelPGSQLRepository: travelPGSQLRepository,
		companyPGSQLRepository: companyPGSQLRepository,
		personPGSQLRepository:  personPGSQLRepository,
		accountPGSQLRepository: accountPGSQLRepository,
		ticketPGSQLRepository:  ticketPGSQLRepository,
	}
}

func (u *GetEmployeesTicketsUseCase) Execute(
	command *getEmployeesTicketsCommand,
) (*[]dto.EmployeeTicket, error) {
	var output []dto.EmployeeTicket

	companyAccount, err := u.accountPGSQLRepository.FindAccountByID(command.companyId)
	if err != nil {
		return &output, err
	}

	companyPerson, err := u.personPGSQLRepository.FindPersonByID(companyAccount.PersonID)
	if err != nil {
		return &output, err
	}

	company, err := u.companyPGSQLRepository.FindCompanyByID(int(companyPerson.CompanyID.Int64))
	if err != nil {
		return &output, err
	}

	tickets, err := u.ticketPGSQLRepository.GetEmployeesTickets(company.ID)
	if err != nil {
		return &output, err
	}

	
	employees, _ := u.accountPGSQLRepository.FindAccountsByCompanyID(company.ID)
	getEmployee := func(id int) *account_entity.Account {
		for _, employee := range employees {
			if employee.ID == id {
				return employee
			}
		}
		return &account_entity.Account{}
	}

	travels, _ := u.travelPGSQLRepository.FindAll()
	getTravels := func(id int) *travel_entity.Travel {
		for _, travel := range travels {
			if travel.ID == id {
				return travel
			}
		}
		return &travel_entity.Travel{}
	}

	cities, _ := u.cityPGSQLRepository.FindAll()
	getCities := func(id int) *city_entity.City {
		for _, city := range cities {
			if city.ID == id {
				return city
			}
		}
		return &city_entity.City{}
	}

	travelCompaniesAccount, _ := u.companyPGSQLRepository.FindAllTravelCompanies()
	getTravelCompanies := func(id int) *company_entity.Company {
		for _, travelCompany := range travelCompaniesAccount {
			if travelCompany.ID == id {
				return travelCompany
			}
		}
		return &company_entity.Company{}
	}


	for _, ticket := range tickets {
		var status string
		travel := getTravels(ticket.TravelID)
		employee := getEmployee(ticket.AccountID) 
		departureCity := getCities(travel.Departure.CityID)
		arrivalCity := getCities(travel.Arrival.CityID)
		transporter := getTravelCompanies(travel.AccountID)
		if ticket.StatusID == 0 {
			status = "NOT PAID"
		} else {
			status = "PAID"
		}
		output = append(output, *dto.NewEmployeeTicket(
			ticket.ID,
			status,
			employee.Email,
			travel.Price,
			transporter.SocialReason,
			travel.Departure.Time,
			departureCity.Name,
			travel.Arrival.Time,
			arrivalCity.Name,
			ticket.CreatedAt,
		))
	}
	return &output, nil
}
