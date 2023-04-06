package setup_di_config

import (
	get_last_purchases "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-last-purchases"
	get_possible_travels_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"

	login_as_company_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/company/login"
	login__as_customer_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
)

type App struct {
	LoginAsCustomerUseCase login__as_customer_usecase.LoginAsCustomerUseCaseInterface
	LoginAsCompanyUseCase login_as_company_usecase.LoginAsCompanyUseCaseInterface
	GetPossibleTravelsUseCase  get_possible_travels_usecase.GetPossibleTravelsUseCaseInterface
	GetLastPurchasesUseCase get_last_purchases.GetLastPurchasesUseCaseInterface
}

