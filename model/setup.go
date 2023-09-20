package model

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func SetupModels() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(strings.Join([]string{"Can't connect to database", err.Error()}, " -> "))
	}

	return db
}
