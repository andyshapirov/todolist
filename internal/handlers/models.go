package handlers

import "github.com/andyshapirov/todolist/internal/storage"

type TaskHandler struct {
	Password string
	Secret   string
	Service  *storage.TaskService
}

func NewTaskHandler(pass, secret string, s *storage.TaskService) *TaskHandler {
	return &TaskHandler{
		Password: pass,
		Secret:   secret,
		Service:  s,
	}
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
