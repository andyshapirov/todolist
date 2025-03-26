package server

import (
	"net/http"

	"github.com/andyshapirov/todolist/internal/handlers"
	"github.com/go-chi/chi"
)

func setupRoutes(r *chi.Mux, h *handlers.TaskHandler) {
	r.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("./web"))))

	r.Post("/api/signin", h.SignIn)
	r.Get("/api/task", handlers.Auth(h.Password, h.Secret, h.GetTask))
	r.Post("/api/task", handlers.Auth(h.Password, h.Secret, h.CreateTask))
	r.Put("/api/task", handlers.Auth(h.Password, h.Secret, h.UpdateTask))
	r.Get("/api/tasks", handlers.Auth(h.Password, h.Secret, h.GetUpcommingTasks))
	r.Post("/api/task/done", handlers.Auth(h.Password, h.Secret, h.GetDoneTask))
	r.Delete("/api/task", handlers.Auth(h.Password, h.Secret, h.RemoveTask))
}
