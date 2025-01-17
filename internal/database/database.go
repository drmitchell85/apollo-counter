package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// this package is needed for the PostgreSQL drivers to work
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	// check connection
	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		log.Printf("could not insert test string into redis, with error: %s", err)
		panic(err)
	}

	_, err = client.Get(ctx, "foo").Result()
	if err != nil {
		log.Printf("could not connect to redis, with error: %s", err)
		panic(err)
	} else {
		log.Print("connected to redis")
	}

	return client
}
