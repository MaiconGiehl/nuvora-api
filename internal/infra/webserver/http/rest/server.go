package rest

import (
	"context"
	"net/http"

	_ "github.com/lib/pq"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/configs/env"
	_ "github.com/maicongiehl/nuvora-api/docs"
	postgresdb_config "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg"
)

//	@title			Nuvora API
//	@version		1.0
//	@description	Product API with auhtentication
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Nuvora
//	@contact.url	https://techtur.com.br
//	@contact.email	atendimento@nuvora.com.br

//	@license.name	Nuvora Promotora License
//	@license.url	https://nuvora.com.br

//	@host		localhost:8080
//	@BasePath	/
func StartServer() {
	port := ":8080"

	ctx := context.Background()
	
	env := env.LoadConfig()

	db := postgresdb_config.ConnectWithConnector(env.DBHost, env.DBPort, env.DBUser, env.DBPassword, env.DBName)
	err := db.Ping()

	if err != nil {
		panic(err)
	}
	app := di.SetupDIConfig(ctx, db)

	http.ListenAndServe(port, Router(app))
}