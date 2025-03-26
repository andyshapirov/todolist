package main

import (
	"github.com/andyshapirov/todolist/internal/config"
	"github.com/andyshapirov/todolist/internal/database"
	"github.com/andyshapirov/todolist/internal/server"
	"github.com/andyshapirov/todolist/internal/storage"
	_ "modernc.org/sqlite"
)

func main() {
	cfg := config.LoadConfig()

	db := database.InitDatabase(cfg.DBFile)
	defer db.Close()

	s := storage.NewTaskService(db)

	srv := server.NewServer(cfg.Port, cfg.Password, cfg.Secret, s)
	srv.Run()
}
