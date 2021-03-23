package services

import (
	"database/sql"
	"fmt"

	// import postgres for drivers
	_ "github.com/lib/pq"
	"github.com/luiscaro1/go-postgres-api/server/error_handlers"
)

// Establish the config for the postgres instance
const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "sample-db"
)

// OpenConnection opens the connection with the postgres db
func OpenConnection() *sql.DB {

	// convert the config to a connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open the connect to the database
	db, err := sql.Open("postgres", psqlConn)
	error_handlers.CheckError(err, nil)

	// verify that the connection is in fact open
	err = db.Ping()
	error_handlers.CheckError(err, nil)

	// return the db instance
	return db

}
