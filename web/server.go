package web

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	port   string
	router *chi.Mux
}

func NewServer(port string) *server {
	s := &server{
		port:   port,
		router: chi.NewRouter(),
	}

	s.setupRoutes()

	return s
}

func (s *server) setupRoutes() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World"))
	})
}

func (s *server) Start() {
	log.Println("Sherlock started listening on " + s.port)
	http.ListenAndServe(":"+s.port, s.router)
}
