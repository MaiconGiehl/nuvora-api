package setup_di_config

import (
	get_last_purchases "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-last-purchases"
	get_possible_travels_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"

	login_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
)

type App struct {
	LoginAsCustomerUseCase login_usecase.LoginAsCustomerUseCaseInterface
	GetPossibleTravelsUseCase  get_possible_travels_usecase.GetPossibleTravelsUseCaseInterface
	GetLastPurchasesUseCase get_last_purchases.GetLastPurchasesInterface
}

