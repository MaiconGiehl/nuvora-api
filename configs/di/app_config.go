package setup_di_config

import (
	"context"
	"database/sql"

	account_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/account"
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
	buy_ticket_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/buy-ticket"
	get_last_purchases_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-last-purchases"
	get_possible_travels_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"
	login_as_customer_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
	create_travel_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/travel-company/create-travel"
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

	newTicketPGSQLRepository := ticket_entity.NewTicketPGSQLRepository(ctx, db)
	newTravelPGSQLRepository := travel_entity.NewTravelPGSQLRepository(ctx, db)

	// CompanyUseCases
	newCreateEmployeeUseCase := create_employee_usecase.NewCreateEmployeeUseCase(ctx, 
		newCustomerPGSQLRepository, 
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
	)

	newGetEmployeesTicketsUseCase := get_employees_tickets_usecase.NewGetEmployeesTicketsUseCase(
		ctx,
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

	// CustomerUseCases
	newLoginAsCustomerUseCase := login_as_customer_usecase.NewLoginAsCustomerUseCase(
		ctx,
		logrus,
		newCustomerPGSQLRepository, 
		newPersonPGSQLRepository,
		newAccountPGSQLRepository,
	)

	newGetLastPurchasesUseCase := get_last_purchases_usecase.NewGetLastPurchasesUsecase(
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

	newBuyTicketUseCase := buy_ticket_usecase.NewBuyTicketUsecase(
		ctx,
		newTicketPGSQLRepository,
	)

	// TravelCompanyUseCases
	newCreateTravelUseCase := create_travel_usecase.NewCreateTravelUseCase(
		ctx,
		newTravelPGSQLRepository,
	)


	return &App{
		BuyTicketUseCase: newBuyTicketUseCase,
		CreateTravelUseCase: newCreateTravelUseCase,
		CreateEmployeeUseCase: newCreateEmployeeUseCase,
		GetEmployeesTicketsUsecase: *newGetEmployeesTicketsUseCase,
		GetEmployeesUseCase: newGetEmployeesUseCase,
		GetLastPurchasesUseCase: newGetLastPurchasesUseCase,
		GetPossibleTravelsUseCase: newGetPossibleTravelsUseCase,
		LoginAsCompanyUseCase: newLoginAsCompanyUseCase,
		LoginAsCustomerUseCase: newLoginAsCustomerUseCase,
	}
}