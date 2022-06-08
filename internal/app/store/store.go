package store

import (
	"database/sql"
	"hack/internal/app/config"
	"time"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func New(config config.Config) (*Store, error) {

	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Устанавливаем параметры
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))

	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
