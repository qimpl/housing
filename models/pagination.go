package models

import "time"

// Cursor yo
type Cursor struct {
	ID        string
	CreatedAt time.Time
}

// PaginationHeaders a
type PaginationHeaders struct {
	Link   []string
	Cursor string
}
