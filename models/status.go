package models

import (
	"time"

	"github.com/google/uuid"
)

// Status represent the database statuses table structure
// It's also used in POST status creation method
type Status struct {
	ID        uuid.UUID `json:"id" pg:"id" swaggerignore:"true" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	Name      string    `json:"name" example:"sold"`
	CreatedAt time.Time `json:"created_at,omitempty" swaggerignore:"true" example:"1977-04-22T06:00:00Z"`
	UpdatedAt time.Time `json:"updated_at,omitempty" swaggerignore:"true" example:"1977-04-22T06:00:00Z"`
}
