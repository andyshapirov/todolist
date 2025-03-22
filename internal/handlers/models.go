package handlers

import "github.com/andyshapirov/todolist/internal/services"

type TaskHandler struct {
	Service *services.TaskService
}

func NewTaskHandler(s *services.TaskService) *TaskHandler {
	return &TaskHandler{Service: s}
}

type IDResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type PasswordRequest struct {
	Password string `json:"password"`
}

type TokenRequest struct {
	Token string `json:"token"`
}
