package setup_di_config

import (
	create_employee_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/create-employee"
	get_employees_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees"
	get_employees_tickets_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/get-employees-tickets"
	login_as_company_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/login"
	get_last_purchases "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-last-purchases"
	get_possible_travels_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"
	login_as_customer_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
)

type App struct {
	CreateEmployeeUseCase create_employee_usecase.CreateEmployeeUseCaseInterface
	GetEmployeesTicketsUsecase get_employees_tickets_usecase.GetEmployeesTicketsUseCase  
	GetEmployeesUseCase get_employees_usecase.GetEmployeesUseCaseInterface
	GetLastPurchasesUseCase get_last_purchases.GetLastPurchasesUseCaseInterface
	GetPossibleTravelsUseCase  get_possible_travels_usecase.GetPossibleTravelsUseCaseInterface
	LoginAsCompanyUseCase login_as_company_usecase.LoginAsCompanyUseCaseInterface
	LoginAsCustomerUseCase login_as_customer_usecase.LoginAsCustomerUseCaseInterface
}

