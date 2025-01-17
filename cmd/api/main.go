package main

import (
	"apollo-counter/internal/controllers"
	"apollo-counter/internal/database"
	"apollo-counter/internal/handlers"
	"apollo-counter/internal/repository"
	"apollo-counter/internal/server"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	// get all our config information
	// start our DB
	// start logger

	// get the NewServer handler
	srv := server.NewServer()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: srv.Router,
	}

	db := database.InitDB()
	rdb := database.InitRedis()

	ur := repository.NewUserRepository(db)
	er := repository.NewEventRepository(db, rdb)

	uc := controllers.NewUserController(ur)
	ec := controllers.NewEventController(er)

	uh := handlers.NewUserHandler(uc)
	eh := handlers.NewEventHandler(ec)

	server.AddRoutes(srv, uh, eh)

	log.Printf("listening on %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		return err
	}

	return nil
}
