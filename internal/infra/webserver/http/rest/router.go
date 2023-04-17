package rest

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

type AppRouter struct {
	app *di.App
	tokenAuth *jwtauth.JWTAuth
	JWTClaims struct {
		ExpiresIn int
	}
	logger logger.Logger
}

func NewAppRouter(
	app *di.App,
	tokenAuth *jwtauth.JWTAuth,
	logger logger.Logger,
) *AppRouter {
	return &AppRouter{
		app: app,
		tokenAuth: tokenAuth,
		logger: logger,
	}
}

func (ar *AppRouter) Route() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.WithValue("jwt", ar.tokenAuth))
	r.Use(middleware.WithValue("JwtExpiresIn", ar.JWTClaims.ExpiresIn))

	r.Get("/docs/nuvora/v1*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/nuvora/v1/doc.json")))
	
	r.Route("/nuvora/v1", func (r chi.Router) {
		r.Route("/customer", routes.NewCustomerRouter(ar.logger, ar.app).CustomerRoutes)
		r.Route("/company", routes.NewCompanyRoutes(ar.logger, ar.app).CompanyRoutes)
		r.Route("/travel-company", routes.NewTravelCompanyRouter(ar.logger, ar.app).CompanyRoutes)
	}) 

	return r
}