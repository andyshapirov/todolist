package server

import (
	"log"
	"net/http"

	"github.com/andyshapirov/todolist/internal/handlers"
	"github.com/andyshapirov/todolist/internal/storage"
	"github.com/go-chi/chi"
)

type Server struct {
	Router *chi.Mux
	Port   string
}

func NewServer(port, pass, secret string, s *storage.TaskService) *Server {
	r := chi.NewRouter()
	h := handlers.NewTaskHandler(pass, secret, s)
	setupRoutes(r, h)

	return &Server{Router: r, Port: port}
}

func (s *Server) Run() {
	log.Println("todolist running...")
	if err := http.ListenAndServe(":"+s.Port, s.Router); err != nil {
		log.Fatal(err)
	}
}
