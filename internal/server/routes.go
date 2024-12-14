package server

import (
	"apollo-counter/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func AddRoutes(server *Server, uh *handlers.UserHandler) {
	r := server.Router

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ping!"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/newUser", uh.NewUser)
	})

}
