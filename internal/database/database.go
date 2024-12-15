package database

import (
	"database/sql"
	"fmt"
	"log"

	// this package is needed for the PostgreSQL drivers to work
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

type Repositories interface {
	CounterUp() error
	CounterDown() error
}

const (
	host     = "localhost"
	port     = 5432
	user     = "dmitch"
	password = "password123"
	dbname   = "apollodb"
)

func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("could not connect to db %s, with error: %s", dbname, err)
		panic(err)
	} else {
		log.Printf("connected to db %s", dbname)
	}

	return db
}
