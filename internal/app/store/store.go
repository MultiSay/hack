package store

import (
	"database/sql"
	"errors"
	"fmt"
	"hack/internal/app/config"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Store struct {
	db *sql.DB
}

func New(config config.Config) (*Store, error) {

	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse DATABASE_URL '%s'", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to create connection pool '%s'", err)
	}

	// Устанавливаем параметры
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))

	err = initMigrations(config.URL)
	if err != nil && err != migrate.ErrNoChange && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("unable to create database '%s'", err)
	}

	return &Store{
		db: db,
	}, nil
}

func initMigrations(databaseDSN string) error {
	m, err := migrate.New(
		"file://migrations",
		databaseDSN)
	if err != nil {
		if err == os.ErrNotExist {
			return nil
		}
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	log.Printf("[INIT] Migrations UP")
	return nil
}
