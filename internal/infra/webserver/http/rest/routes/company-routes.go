package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/handlers"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/middlewares"
)

type CompanyRouter struct {
	tokenAuth *jwtauth.JWTAuth
	logger logger.Logger
	app    *di.App
}

func NewCompanyRoutes(
	tokenAuth *jwtauth.JWTAuth,
	logger logger.Logger,
	app *di.App,
) *CompanyRouter {
	return &CompanyRouter{
		logger: logger,
		app:    app,
		tokenAuth: tokenAuth,
	}
}

func (router *CompanyRouter) CompanyRoutes(r chi.Router) {

	companyHandler := handlers.NewCompanyHandler(router.logger, router.app)

	r.Post("/", companyHandler.Login)

	// Protected routes
	r.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(router.tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(middlewares.CompanyMiddleware)
		r.Get("/{id}/employees", companyHandler.GetEmployees)
		r.Get("/{id}/employees/tickets", companyHandler.GetEmployeesTickets)
		r.Patch("/{id}/employees/tickets", companyHandler.PayAllTickets)
		r.Delete("/{id}/employee/{employeeId}", companyHandler.DeleteEmployee)
	})
}
