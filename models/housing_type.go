package models

import (
	"time"

	"github.com/google/uuid"
)

// HousingType represent the database housing_types table
// It's also used in POST creation methods
type HousingType struct {
	ID        uuid.UUID `json:"id" pg:"id" swaggerignore:"true" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	Name      string    `json:"name" example:"garage"`
	CreatedAt time.Time `json:"created_at,omitempty" swaggerignore:"true" example:"1977-04-22T06:00:00Z"`
	UpdatedAt time.Time `json:"updated_at,omitempty" swaggerignore:"true" example:"1977-04-22T06:00:00Z"`
}
