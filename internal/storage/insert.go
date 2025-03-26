package storage

import (
	"database/sql"

	"github.com/andyshapirov/todolist/internal/database"
)

func (s TaskService) InsertOne(t *database.Task) (int, error) {
	q := `
	INSERT INTO scheduler (date, title, comment, repeat)
	VALUES (:date, :title, :comment, :repeat);`
	res, err := s.Database.Exec(q,
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
