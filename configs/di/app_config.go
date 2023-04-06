package setup_di_config

import (
	"context"
	"database/sql"

	get_last_purchases_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-last-purchases"
	get_possible_travels_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"
	login_usecase "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
	customer_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/customer"
	ticket_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/ticket"
	travel_entity "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg/travel"
)

func SetupDIConfig(
	ctx context.Context,
	db *sql.DB,
) *App {

	newCustomerPGSQLRepository := customer_entity.NewCustomerPGSQLRepository(ctx, db)
	newLoginAsCustomerUseCase := login_usecase.NewLoginAsCustomerUseCase(ctx, newCustomerPGSQLRepository)

	newTravelPGSQLRepository := travel_entity.NewTravelPGSQLRepository(ctx, db)
	newGetPossibleTravelsUseCase := get_possible_travels_usecase.NewGetPossibleTravelsUseCase(ctx, newTravelPGSQLRepository)

	newTicketPGSQLRepository := ticket_entity.NewTicketPGSQLRepository(ctx, db)
	newGetLastPurchasesUseCase := get_last_purchases_usecase.NewGetLastPurchases(ctx, newTicketPGSQLRepository)

	return &App{
		LoginAsCustomerUseCase: newLoginAsCustomerUseCase,
		GetPossibleTravelsUseCase: newGetPossibleTravelsUseCase,
		GetLastPurchasesUseCase: newGetLastPurchasesUseCase,
	}
}