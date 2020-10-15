package models

import (
	"time"

	"github.com/google/uuid"
)

// Housing represent the database housing table
type Housing struct {
	ID             uuid.UUID `json:"id,omitempty" pg:"id" swaggerignore:"true"`
	TypeID         uuid.UUID `json:"type_id" pg:"type_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	SurfaceArea    float32   `json:"surface_area" example:"15.5"`
	RentPrice      float32   `json:"rent_price" example:"60.95"`
	RentalCharges  float32   `json:"rental_charges" example:"60.9"`
	Country        string    `json:"country" example:"FR"`
	State          string    `json:"state" example:"Haut-de-France"`
	City           string    `json:"city" example:"Lille"`
	Street         string    `json:"street" example:"78 Rue Solférino"`
	Zip            string    `json:"zip" example:"59000"`
	IsFurnished    bool      `json:"is_furnished,omitempty" example:"false"`
	HasElectricity bool      `json:"has_electricity" example:"true"`
	HasGas         bool      `json:"has_gas" example:"false"`
	StatusID       uuid.UUID `json:"status_id" pg:"status_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	OwnerID        uuid.UUID `json:"owner_id" pg:"owner_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	LastTenantID   uuid.UUID `json:"last_tenant_id,omitempty" pg:"last_tenant_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	CreatedAt      time.Time `json:"created_at,omitempty" swaggerignore:"true"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" swaggerignore:"true"`
}
