package server

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/andyshapirov/todolist/internal/handlers"
	"github.com/andyshapirov/todolist/internal/services"
	"github.com/andyshapirov/todolist/tests"
)

type Server struct {
	Mux  *http.ServeMux
	Port string
}

func NewServer(s *services.TaskService) *Server {
	mux := http.NewServeMux()
	h := handlers.NewTaskHandler(s)
	setupRoutes(mux, h)

	p := strconv.Itoa(tests.Port)
	if v, ok := os.LookupEnv("TODO_PORT"); len(v) > 0 && ok {
		p = v
	}

	return &Server{Mux: mux, Port: p}
}

func (s *Server) Run() {
	log.Printf("Starting server on localhost:%s\n", s.Port)
	if err := http.ListenAndServe(":"+s.Port, s.Mux); err != nil {
		log.Fatal(err)
	}
}
