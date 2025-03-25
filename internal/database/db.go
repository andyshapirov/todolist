package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func InitDatabase(envDBFile string) *sql.DB {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")

	if len(envDBFile) > 0 {
		dbFile = envDBFile
	}

	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Fatal(err)
	}

	initSchema(db)

	return db
}

func initSchema(db *sql.DB) {
	q := `
	CREATE TABLE IF NOT EXISTS scheduler (
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
