package rest

import (
	"context"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	"github.com/maicongiehl/nuvora-api/configs/env"
	postgresdb_config "github.com/maicongiehl/nuvora-api/internal/infra/dataprovider/sql/pg"
	logrus_config "github.com/maicongiehl/nuvora-api/internal/infra/log/logrus"
)

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
	appRouter := NewAppRouter(app, env.TokenAuth, logger)


	expiresIn, _ := strconv.Atoi(env.JWTExpiresIn)
	appRouter.JWTClaims.ExpiresIn = expiresIn

	logger.Infof("Server.StartServer: Starting server in %s", env.WebServerPort)
	err := http.ListenAndServe(port, appRouter.Route())
	if err != nil {
		logger.Fatalf("Rest.StartServer: Unable to start server, %s", err.Error())
	}
}
