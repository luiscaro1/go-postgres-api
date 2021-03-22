package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	. "github.com/luiscaro1/go-postgres-api/server/error_handlers"
)

// Establish the config for the postgres instance
const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "sample-db"
)

func OpenConnection() *sql.DB {

	// convert the config to a connection string
	psql_conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open the connect to the database
	db, err := sql.Open("postgres", psql_conn)
	CheckError(err)

	// verify that the connection is in fact open
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// return the db instance
	return db

}
