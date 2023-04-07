package postgresdb

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
)

func ConnectWithConnector(dbHost, dbPort, dbUser, dbPassword, dbName string) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
	  panic(err)
	}

	fmt.Print("Connected")
	err = db.Ping()
	if err != nil {
	  panic(err)
	}

	return db
}
