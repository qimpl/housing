package models

import (
	"time"

	"github.com/google/uuid"
)

// CreateVisitBody is used inside visit creation request
type CreateVisitBody struct {
	Date      time.Time `json:"date" example:"2020-11-05"`
	Hour      string    `json:"hour" example:"18:00"`
	HousingID uuid.UUID `json:"housing_id" example:"cb7bc97f-45b0-4972-8edf-dc7300cc059c"`
	Owner     struct {
		FirstName string `json:"first_name" example:"Jean"`
		Email     string `json:"email" example:"jean@qimpl.io"`
	} `json:"owner"`
}

// AcceptVisit is used as body value inside accept visit request
type AcceptVisit struct {
	UserEmail string `json:"user_email" example:"jean@qimpl.io"`
}

// Visit a single visit inside the database
type Visit struct {
	ID         uuid.UUID `json:"id,omitempty" pg:"id" swaggerignore:"true"`
	Date       time.Time `json:"date" example:"2020-11-05"`
	Hour       string    `json:"hour" example:"18:00"`
	IsAccepted bool      `json:"is_accepted,omitempty" example:"false"`
	HousingID  uuid.UUID `json:"housing_id" example:"cb7bc97f-45b0-4972-8edf-dc7300cc059c"`
	CreatedAt  time.Time `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
}
