package storage

import (
	"database/sql"
	"errors"

	"github.com/andyshapirov/todolist/internal/database"
)

func (s TaskService) GetTasks(search string, limit int) (*map[string][]database.Task, error) {
	searchLike := "%" + search + "%"
	q := `
	SELECT id, date, title, comment, repeat FROM scheduler
	WHERE date = :search OR title LIKE :searchLike OR comment LIKE :searchLike OR :search = ''
	ORDER BY date
	LIMIT :limit;`
	rows, err := s.Database.Query(q,
		sql.Named("search", search),
		sql.Named("searchLike", searchLike),
		sql.Named("limit", limit),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[string][]database.Task)
	var task database.Task
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
			return nil, err
		}
		tasks["tasks"] = append(tasks["tasks"], task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if _, ok := tasks["tasks"]; !ok {
		tasks["tasks"] = []database.Task{}
	}

	return &tasks, nil
}

func (s TaskService) GetTask(id int) (*database.Task, error) {
	q := `
	SELECT id, date, title, comment, repeat FROM scheduler
	WHERE id = :id;`
	row := s.Database.QueryRow(q, sql.Named("id", id))

	var task database.Task
	if err := row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task not exist")
		}

		return nil, err
	}

	return &task, nil
}
