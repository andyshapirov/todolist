package server

import (
	"net/http"

	"github.com/andyshapirov/todolist/internal/handlers"
)

func setupRoutes(mux *http.ServeMux, h *handlers.TaskHandler) {
	mux.Handle("/", http.FileServer(http.Dir("./web")))

	mux.HandleFunc("/api/nextdate", h.TaskDate)

	mux.HandleFunc("/api/signin", h.SignIn)
	mux.HandleFunc("/api/task", handlers.Auth(h.CreateUpdateRemoveTask))
	mux.HandleFunc("/api/tasks", handlers.Auth(h.UpcommingTasks))
	mux.HandleFunc("/api/task/done", handlers.Auth(h.DoneTask))
}
