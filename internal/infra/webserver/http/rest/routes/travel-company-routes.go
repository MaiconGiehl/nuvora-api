package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/handlers"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/middlewares"
)

type TravelCompanyRouter struct {
	tokenAuth *jwtauth.JWTAuth
	logger    logger.Logger
	app       *di.App
}

func NewTravelCompanyRouter(
	tokenAuth *jwtauth.JWTAuth,
	logger logger.Logger,
	app *di.App,
) *TravelCompanyRouter {
	return &TravelCompanyRouter{
		tokenAuth: tokenAuth,
		logger:    logger,
		app:       app,
	}
}

func (routes *TravelCompanyRouter) TravelCompanyRoutes(r chi.Router) {
	travelCompanyHandler := handlers.NewTravelCompanyHandler(routes.logger, routes.app)

	r.Use(jwtauth.Verifier(routes.tokenAuth))
	r.Use(jwtauth.Authenticator)
	r.Use(middlewares.TravelCompanyMiddleware)
	r.Get("/{id}/bus", travelCompanyHandler.GetAllBus)
	r.Post("/{id}/travels", travelCompanyHandler.CreateTravel)
	r.Post("/{id}/travel/{travelId}", travelCompanyHandler.DeleteTravel)
}
