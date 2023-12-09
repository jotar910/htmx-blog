package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("invalid command: expected up/down argument")
	}
	operation := os.Args[1]
	if operation != "up" && operation != "down" {
		log.Fatalf("invalid operation argument: expected up/down argument but go %s", operation)
	}
	db, err := sql.Open("sqlite3", filepath.Join(".", "articles.db"))
	if err != nil {
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			return
		}
	}()
	driver, err := WithInstance(db, &Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"ql", driver)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("running migration %s", operation)
	switch operation {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	}
	if err != nil {
		log.Fatal(errors.Wrap(err, "running migration"))
	}
}
