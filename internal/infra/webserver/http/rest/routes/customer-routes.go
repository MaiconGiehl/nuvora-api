package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/handlers"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/middlewares"
)

type CustomerRouter struct {
	tokenAuth *jwtauth.JWTAuth
	logger logger.Logger
	app    *di.App
}

func NewCustomerRouter(
	tokenAuth *jwtauth.JWTAuth,
	logger logger.Logger,
	app *di.App,
) *CustomerRouter {
	return &CustomerRouter{
		tokenAuth: tokenAuth,
		logger: logger,
		app:    app,
	}
}

func (router *CustomerRouter) CustomerRoutes(r chi.Router) {
	customerHandler := handlers.NewCustomerHandler(router.logger, router.app)

	r.Post("/", customerHandler.Login)

	// Protected routes
	r.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(router.tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(middlewares.CustomerMiddleware)
		r.Get("/{id}/tickets", customerHandler.Purchases)
		r.Post("/{id}/tickets/{travelId}", customerHandler.BuyTicket)
		r.Get("/{id}/travels", customerHandler.PossibleTravels)
	})
}
