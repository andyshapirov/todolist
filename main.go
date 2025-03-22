package main

import (
	"github.com/andyshapirov/todolist/internal/database"
	"github.com/andyshapirov/todolist/internal/server"
	"github.com/andyshapirov/todolist/internal/services"
	_ "modernc.org/sqlite"
)

func main() {
	db := database.InitDatabase()
	defer db.Close()

	s := services.NewTaskService(db)

	srv := server.NewServer(s)
	srv.Run()
}
