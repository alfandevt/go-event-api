package database

import (
	"database/sql"
	"log"
	"path"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var dbPath = path.Join("internal", "database")
var migrationPath = path.Join(dbPath, "migrations")

const dbFile = "event-api.db"

func InitDB(migrate bool) {
	var err error

	DB, err = sql.Open("sqlite3", filepath.Join(dbPath, dbFile))
	if err != nil {
		panic("could not connect to database")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if migrate {
		log.Println("initializing migrations on:")
		log.Println(filepath.Join(dbPath, dbFile))
		err = runInitMigration()
		if err != nil {
			log.Println(err)
			panic("could not run migrations")
		}
		log.Print("initializing migrations: ok")
	}
}

func runInitMigration() error {

	dbSourceURL := "sqlite3://" + filepath.Join(dbPath, dbFile)
	migrationSourcesURL := "file://" + migrationPath
	log.Println(dbSourceURL)
	log.Println(migrationSourcesURL)

	log.Println("running migration...")
	m, err := migrate.New(migrationSourcesURL, dbSourceURL)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Println("unable to proceed migration")
		return err
	}

	return nil
}
