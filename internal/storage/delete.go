package storage

import (
	"database/sql"
	"errors"
)

func (s TaskService) DeleteOne(id int) error {
	q := `
	DELETE FROM scheduler
	WHERE id = :id;`
	res, err := s.Database.Exec(q, sql.Named("id", id))
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
