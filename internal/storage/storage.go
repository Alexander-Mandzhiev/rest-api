package storage

import (
	"database/sql"
	"os"
	"tags/internal/entity"

	_ "github.com/lib/pq"
)

type Storage interface {
	Create(*entity.Tag) error
	Delete(string) error
	Update(*entity.Tag) error
	GetAll() ([]*entity.Tag, error)
	GetOne(string) (*entity.Tag, error)
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgreStore, error) {
	database_string := os.Getenv("DATABASE_STRING")
	connStr := database_string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgreStore{db: db}, nil
}
