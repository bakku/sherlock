package web

import (
	"log"
	"net/http"

	"bakku.org/sherlock/web/routes"
	"github.com/go-chi/chi"
)

type Server struct {
	port   string
	router *chi.Mux
}

func NewServer(port string) *Server {
	s := &Server{
		port:   port,
		router: chi.NewRouter(),
	}

	s.setupRoutes()

	return s
}

func (s *Server) setupRoutes() {
	home := routes.Home{}

	s.router.Get("/", home.ServeHTTP)
}

func (s *Server) Start() {
	log.Println("Sherlock started listening on " + s.port)
	http.ListenAndServe(":"+s.port, s.router)
}
