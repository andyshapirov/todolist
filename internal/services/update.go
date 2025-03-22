package services

import (
	"database/sql"
	"errors"

	"github.com/andyshapirov/todolist/internal/database"
)

func (s TaskService) UpdateOne(t *database.Task) error {
	if t == nil {
		return errors.New("internal server error")
	}

	q := `
	UPDATE scheduler
	SET  date = :date, title = :title, comment = :comment, repeat = :repeat
	WHERE id = :id;`
	res, err := s.Database.Exec(q,
		sql.Named("date", t.Date),
		sql.Named("title", t.Title),
		sql.Named("comment", t.Comment),
		sql.Named("repeat", t.Repeat),
		sql.Named("id", t.ID),
	)
	if err != nil {
		return err
	}

	updRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if updRows == 0 {
		return errors.New("task not exist")
	}

	return nil
}

func (s TaskService) UpdateDateOne(id int, date string) error {
	q := `
	UPDATE scheduler
	SET  date = :date
	WHERE id = :id;`
	res, err := s.Database.Exec(q,
		sql.Named("date", date),
		sql.Named("id", id),
	)
	if err != nil {
		return err
	}

	updRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if updRows == 0 {
		return errors.New("task not exist")
	}

	return nil
}
