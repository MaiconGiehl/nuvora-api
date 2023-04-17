package setup_di_config

import (
	"context"
	"database/sql"

	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
	bus_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/bus"
	city_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/city"
	company_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/company"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	person_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/person"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	create_employee_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/create-employee"
	get_employees_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees"
	get_employees_tickets_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees-tickets"
	login_as_company_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/login"
	pay_tickets_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/pay-tickets"
	buy_ticket_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/buy-ticket"
	get_possible_travels_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"
	get_purchases_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-purchases"
	login_as_customer_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
	create_travel_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/travel-company/create-travel"
	get_all_bus_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/travel-company/get-all-bus"
)

func SetupDIConfig(
	ctx context.Context,
	db *sql.DB,
	logrus logger.Logger,
) *App {

	newCustomerPGSQLRepository := customer_entity.NewCustomerPGSQLRepository(ctx, db, logrus)
	newCompanyPGSQLRepository := company_entity.NewCompanyPGSQLRepository(ctx, db, logrus)
	newPersonPGSQLRepository := person_entity.NewPersonPGSQLRepository(ctx, db, logrus)
	newAccountPGSQLRepository := account_entity.NewAccountPGSQLRepository(ctx, db, logrus)

	newTicketPGSQLRepository := ticket_entity.NewTicketPGSQLRepository(ctx, db, logrus)
	newTravelPGSQLRepository := travel_entity.NewTravelPGSQLRepository(ctx, db, logrus)

	newBusPGSQLRepository := bus_entity.NewBusPGSQLRepository(ctx, db, logrus)
	newCityPGSQLRepository := city_entity.NewCityPGSQLRepository(ctx, db, logrus)

	// CompanyUseCases
	newCreateEmployeeUseCase := create_employee_usecase.NewCreateEmployeeUseCase(ctx,
		newCustomerPGSQLRepository,
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
	)

	newGetEmployeesTicketsUseCase := get_employees_tickets_usecase.NewGetEmployeesTicketsUseCase(
		ctx,
		logrus,
		newCompanyPGSQLRepository,
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
		newTicketPGSQLRepository,
	)

	newGetEmployeesUseCase := get_employees_usecase.NewGetEmployeesUseCase(
		ctx,
		logrus,
		newCustomerPGSQLRepository,
		newCompanyPGSQLRepository,
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
	)

	newLoginAsCompanyUseCase := login_as_company_usecase.NewLoginAsCompanyUseCase(
		ctx,
		logrus,
		newCompanyPGSQLRepository,
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
	)

	newPayTicketsUseCase := pay_tickets_usecase.NewPayTicketsUseCase(
		ctx,
		logrus,
		newCompanyPGSQLRepository,
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
		newTicketPGSQLRepository,
	)

	// CustomerUseCases
	newLoginAsCustomerUseCase := login_as_customer_usecase.NewLoginAsCustomerUseCase(
		ctx,
		logrus,
		newCustomerPGSQLRepository,
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
	)

	newGetPurchasesUseCase := get_purchases_usecase.NewGetPurchasesUsecase(
		ctx, 
		newTicketPGSQLRepository,
	)

	newGetPossibleTravelsUseCase := get_possible_travels_usecase.NewGetPossibleTravelsUseCase(
		ctx,
		newCustomerPGSQLRepository,
		newCompanyPGSQLRepository,
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
		newTravelPGSQLRepository,
	)

	newBuyTicketUseCase := buy_ticket_usecase.NewBuyTicketUseCase(
		ctx,
		logrus,
		newAccountPGSQLRepository,
		newTicketPGSQLRepository,
		newTravelPGSQLRepository,
	)

	// TravelCompanyUseCases
	newCreateTravelUseCase := create_travel_usecase.NewCreateTravelUseCase(
		ctx,
		logrus,
		newBusPGSQLRepository,
		newCityPGSQLRepository,
		newCompanyPGSQLRepository,
		newTravelPGSQLRepository,
	)

	newGetAllBusUseCase := get_all_bus_usecase.NewGetAllBusUseCase(
		ctx,
		logrus,
		newBusPGSQLRepository,
	)


	return &App{
		BuyTicketUseCase: newBuyTicketUseCase,
		CreateTravelUseCase: newCreateTravelUseCase,
		CreateEmployeeUseCase: newCreateEmployeeUseCase,
		GetAllBusUseCase: *newGetAllBusUseCase,
		GetEmployeesTicketsUsecase: *newGetEmployeesTicketsUseCase,
		GetEmployeesUseCase: newGetEmployeesUseCase,
		GetPurchasesUseCase: newGetPurchasesUseCase,
		GetPossibleTravelsUseCase: newGetPossibleTravelsUseCase,
		LoginAsCompanyUseCase: newLoginAsCompanyUseCase,
		LoginAsCustomerUseCase: newLoginAsCustomerUseCase,
		PayTickets: newPayTicketsUseCase,
	}
}
