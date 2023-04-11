package rest

import (
	"context"
	"net/http"

	_ "github.com/lib/pq"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/configs/env"
	_ "github.com/maicongiehl/nuvora-api/docs"
	postgresdb_config "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg"
	logrus_config "github.com/maicongiehl/nuvora-api/internal/infra/log/logrus"
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
	logger := logrus_config.NewLogrusLogger()

	env := env.LoadConfig(".env", logger)

	db := postgresdb_config.ConnectWithDB(
		logger,
		env.DBHost, 
		env.DBPort, 
		env.DBUser, 
		env.DBPassword, 
		env.DBName,
	)

	app := di.SetupDIConfig(ctx, db, logger)

	logger.Infof("Server.StartServer: Starting server in %s", port)
	http.ListenAndServe(port, Router(app, logger))
}