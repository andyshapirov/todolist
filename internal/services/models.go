package services

import "database/sql"

type TaskService struct {
	Database *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{Database: db}
}
