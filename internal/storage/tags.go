package storage

import (
	"database/sql"
	"fmt"
	"tags/internal/entity"
)

func (s *PostgreStore) Init() error {
	return s.CreateTableTags()
}

func (s *PostgreStore) GetAll() ([]*entity.Tag, error) {
	rows, err := s.db.Query(`SELECT * FROM tags`)
	if err != nil {
		return nil, err
	}
	tags := []*entity.Tag{}

	for rows.Next() {
		tag, err := scanSQLTag(rows)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (s *PostgreStore) GetOne(id string) (*entity.Tag, error) {
	rows, err := s.db.Query(`SELECT * FROM tags WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanSQLTag(rows)
	}

	return nil, fmt.Errorf("tag %s not found", id)
}

func (s *PostgreStore) CreateTableTags() error {
	query := `CREATE TABLE IF NOT EXISTS tags (id SERIAL PRIMARY KEY UNIQUE, name VARCHAR(255) NOT NULL UNIQUE);`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgreStore) Create(tag *entity.Tag) error {
	query := `INSERT INTO tags (id, name) VALUES ($1, $2);`
	_, err := s.db.Query(query, tag.ID, tag.Name)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgreStore) Delete(id string) error {
	_, err := s.db.Query(`DELETE FROM tags WHERE id = $1`, id)
	return err
}

func (s *PostgreStore) Update(tag *entity.Tag) error {
	query := `UPDATE tags SET name=$1 WHERE id = $2`
	_, err := s.db.Query(query, tag.Name, tag.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanSQLTag(rows *sql.Rows) (*entity.Tag, error) {
	tag := new(entity.Tag)
	err := rows.Scan(&tag.ID, &tag.Name)
	return tag, err
}
