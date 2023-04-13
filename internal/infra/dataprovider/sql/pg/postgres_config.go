package postgresdb

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"

	"fmt"
)

func ConnectWithDB(
	logger logger.Logger,
	dbHost,
	dbPort,
	dbUser,
	dbPassword,
	dbName string,
) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatalf("PostgresConfig.ConnectWithConnector: Unable to establish connection with database, %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalf("PostgresConfig.ConnectWithConnector: Something went wrong while pinging to database, %s", err.Error())
	}

	logger.Infof("PostgresConfig.ConnectWithDB: Database connection successfully established")

	return db
}
