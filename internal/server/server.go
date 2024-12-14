package server

import "github.com/go-chi/chi"

type Server struct {
	Router *chi.Mux
}

// responsible for all the top-level HTTP matters that relate to all endpoints
// like CORS, auth middleware, and logging
func NewServer() *Server {
	router := chi.NewRouter()
	server := &Server{
		Router: router,
	}

	return server
}
