package web

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"bakku.org/sherlock/web/routes"
	"github.com/go-chi/chi"
)

type Server struct {
	port        string
	templateDir string
	router      *chi.Mux
	templates   map[string]*template.Template
}

func NewServer(port, templateDir string) *Server {
	s := &Server{
		port:        port,
		templateDir: templateDir,
		router:      chi.NewRouter(),
	}

	s.parseTemplates()
	s.setupRoutes()

	return s
}

func (s *Server) parseTemplates() {
	s.templates = make(map[string]*template.Template)

	layouts, err := filepath.Glob(s.templateDir + "shared/*.html")
	if err != nil {
		log.Fatalf("server: could not get layouts: %v\n", err)
	}

	templates, err := filepath.Glob(s.templateDir + "*.html")
	if err != nil {
		log.Fatalf("server: could not get templates: %v\n", err)
	}

	for _, tpl := range templates {
		files := append(layouts, tpl)
		s.templates[filepath.Base(tpl)] = template.Must(template.ParseFiles(files...))
	}
}

func (s *Server) setupRoutes() {
	home := routes.Home{Template: s.templates["home.html"]}

	s.router.Get("/", home.ServeHTTP)
}

func (s *Server) Start() {
	log.Println("Sherlock started listening on " + s.port)
	http.ListenAndServe(":"+s.port, s.router)
}
