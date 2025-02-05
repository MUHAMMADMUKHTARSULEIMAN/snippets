package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	id        int
	title     string
	content   string
	createdAt time.Time
	expiresAt time.Time
}

type SnippetModel struct {
	db *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
