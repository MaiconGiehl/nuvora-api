package setup_di_config

import (
	create_employee_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/create-employee"
	get_employees_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees"
	get_employees_tickets_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees-tickets"
	login_as_company_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/login"
	pay_tickets_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/pay-tickets"
	buy_ticket_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/buy-ticket"
	get_possible_travels_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"
	get_purchases "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-purchases"
	login_as_customer_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
	create_travel_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/travel-company/create-travel"
	get_all_bus_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/travel-company/get-all-bus"
)

type App struct {
	BuyTicketUseCase buy_ticket_usecase.BuyTicketUseCaseInterface
	CreateEmployeeUseCase create_employee_usecase.CreateEmployeeUseCaseInterface
	CreateTravelUseCase create_travel_usecase.CreateTravelUseCaseInterface
	GetAllBusUseCase  get_all_bus_usecase.GetAllBusUseCase
	GetEmployeesTicketsUsecase get_employees_tickets_usecase.GetEmployeesTicketsUseCase  
	GetEmployeesUseCase get_employees_usecase.GetEmployeesUseCaseInterface
	GetPurchasesUseCase get_purchases.GetPurchasesUseCaseInterface
	GetPossibleTravelsUseCase  get_possible_travels_usecase.GetPossibleTravelsUseCaseInterface
	LoginAsCompanyUseCase login_as_company_usecase.LoginAsCompanyUseCaseInterface
	LoginAsCustomerUseCase login_as_customer_usecase.LoginAsCustomerUseCaseInterface
	PayTickets  pay_tickets_usecase.PayTicketsUseCaseInterface
}

