package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/MaiconGiehl/API/config"
	_ "github.com/MaiconGiehl/API/docs"
	"github.com/MaiconGiehl/API/internal/infra/web/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           TechTur Service
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   TechTur
// @contact.url    https://techtur.com.br
// @contact.email  atendimento@techtur.com.br

// @license.name   Acerta Promotora License
// @license.url    https://techtur.com.br

// @host      localhost:8080
// @BasePath  /
func main() {

	dbconfig := config.LoadDBConfig()
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.User, dbconfig.Password, dbconfig.Name)
	
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
	
	http.ListenAndServe(":8080", createRouter(db))
}

func createRouter(db *sql.DB) *chi.Mux {
	var ctx context.Context
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	busHandler := handlers.NewBusHandler(ctx, db)
	
	r.Route("/bus", func(r chi.Router) {
		r.Get("/", busHandler.GetAll)
		r.Post("/", busHandler.CreateBus)
		r.Get("/{id}", busHandler.GetBus)
		r.Delete("/{id}", busHandler.DeleteBus)
		r.Patch("/{id}", busHandler.UpdateBus)
	})
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	
	return r
}