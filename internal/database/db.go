package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func InitDatabase() *sql.DB {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	if v, ok := os.LookupEnv("TODO_DBFILE"); len(v) > 0 && ok {
		dbFile = v
	}
	_, err = os.Stat(dbFile)
	install := err != nil

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	if install {
		createSchema(db)
	}

	return db
}

func createSchema(db *sql.DB) {
	q := `
	CREATE TABLE scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT NOT NULL,
		title TEXT COLLATE NOCASE NOT NULL,
		comment TEXT COLLATE NOCASE NOT NULL,
		repeat TEXT NOT NULL,
		UNIQUE (date, title)
	);
	CREATE INDEX index_scheduler_date ON scheduler(date);`
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}
}
