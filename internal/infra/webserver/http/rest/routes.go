package rest

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router(app *di.App) http.Handler {
	r := chi.NewRouter()
	
	customerHandler := handlers.NewCustomerHandler(app) 

	r.Use(middleware.Logger)
	r.Get("/docs/nuvora/v1*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/nuvora/v1/doc.json")))
	r.Route("/customer", func (r chi.Router) {
		r.Post("/",  customerHandler.Login)
	})

	return r
}