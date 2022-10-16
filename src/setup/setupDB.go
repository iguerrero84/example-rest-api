package setup

import (
	"database/sql"
	"fmt"

	"example-rest-api/src/handler"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345678"
	DB_NAME     = "movies"
)

// DB set up
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	handler.CheckErr(err)
	return db
}
