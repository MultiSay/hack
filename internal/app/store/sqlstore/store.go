package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"hack/internal/app/config"
	"hack/internal/app/store"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Store struct {
	db             *sql.DB
	fileRepository *FileRepository
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

	if err != nil && err != migrate.ErrNoChange && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("unable to create database '%s'", err)
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) File() store.FileRepository {
	if s.fileRepository != nil {
		return s.fileRepository
	}

	s.fileRepository = &FileRepository{
		store: s,
	}

	return s.fileRepository
}
