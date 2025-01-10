package server

import (
	"apollo-counter/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func AddRoutes(server *Server, uh *handlers.UserHandler, eh *handlers.EventHandler) {
	r := server.Router

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ping!"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/email/{email}", uh.GetUserByEmail)
		r.Post("/newUser", uh.NewUser)
	})

	// TODO add events route
	r.Route("/events", func(r chi.Router) {
		// r.Get("/", eh.GetEvents)
		// r.Get("/id/{id}", eh.GetEventByID)
		// r.Get("/popular", eh.GetPopularEvents)
		r.Post("/newEvent", eh.NewEvent)
	})

}
