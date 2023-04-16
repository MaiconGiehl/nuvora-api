package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/configs/env"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	"github.com/maicongiehl/nuvora-api/internal/infra/webserver/http/rest/handlers"
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
		customerHandler := handlers.NewCustomerHandler(ar.logger, ar.app) 
		r.Route("/customer", func (r chi.Router) {
			r.Post("/",  customerHandler.Login)
	
			// Protected routes
			r.Route("/", func(r chi.Router) {
				r.Use(jwtauth.Authenticator)
				r.Use(customerMiddleware)
				r.Get("/{id}/tickets", customerHandler.Purchases)
				r.Post("/{id}/tickets/{travelId}", customerHandler.BuyTicket)
				r.Get("/{id}/travels", customerHandler.PossibleTravels)
			})
		})
		
		companyHandler := handlers.NewCompanyHandler(ar.logger, ar.app)
		r.Route("/company", func (r chi.Router) {
			r.Post("/", companyHandler.Login)
	
			// Protected routes
			r.Route("/", func(r chi.Router) {
				r.Use(jwtauth.Verifier(env.LoadConfig(ar.logger).TokenAuth))
				r.Use(jwtauth.Authenticator)
				r.Use(companyMiddleware)
				r.Get("/{id}/employees", companyHandler.GetEmployees)
				r.Get("/{id}/employees/tickets", companyHandler.GetEmployeesTickets)
				r.Patch("/{id}/employees/tickets", companyHandler.PayAllTickets)
				r.Delete("/{id}/employee/{employeeId}", companyHandler.DeleteEmployee)
			})
		})
		
		travelCompanyHandler := handlers.NewTravelCompanyHandler(ar.logger, ar.app)
		// Protected routes
		r.Route("/travel-company", func (r chi.Router) {
			r.Use(jwtauth.Authenticator)
			r.Use(travelCompanyMiddleware)
			r.Post("/{id}/travels", travelCompanyHandler.CreateTravel)
			r.Post("/{id}/travel/{travelId}", travelCompanyHandler.DeleteTravel)
		})
	}) 

	return r
}

func customerMiddleware(h http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		permissionLevel, err := getHeaderValues(r)
		if err != nil {
			json.NewEncoder(w).Encode("invalid token")
			return 
		}
		permissionLevelNeeded := 3
		err = validateToken(w, r, permissionLevelNeeded, permissionLevel)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return

		}
		h.ServeHTTP(w, r)
	})
}

func companyMiddleware(h http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		permissionLevel, err := getHeaderValues(r)
		if err != nil {
			json.NewEncoder(w).Encode("invalid token")
			return 
		}
		permissionLevelNeeded := 2
		err = validateToken(w, r, permissionLevelNeeded, permissionLevel)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		h.ServeHTTP(w, r)
	})
}

func travelCompanyMiddleware(h http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		permissionLevel, err := getHeaderValues(r)
		if err != nil {
			json.NewEncoder(w).Encode("invalid token")
			return 
		}
		
		permissionLevelNeeded := 1
		err = validateToken(w, r, permissionLevelNeeded, permissionLevel)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		h.ServeHTTP(w, r)
	})
}

func getHeaderValues(r *http.Request) (int, error){
	token := strings.Split(r.Header.Values("Authorization")[0], "Bearer ")[1]
	permissionLevel, err := extractPermissionLevel(token)
	return permissionLevel, err
}

func validateToken(w http.ResponseWriter, r *http.Request, permissionLevelNeeded, permissionLevel int) error {
	if permissionLevel != permissionLevelNeeded {
		return errors.New("unauthorized")
	}
	return nil
}

func extractPermissionLevel(tokenString string) (int, error) {
	var pm int
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
			return -1, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		pm, err = strconv.Atoi(fmt.Sprint(claims["permission_level"]))
		if err != nil {
			return -1, errors.New("invalid token")
		}
	}

	return pm, nil
}