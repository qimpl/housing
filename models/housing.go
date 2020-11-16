package models

import (
	"time"

	"github.com/google/uuid"
)

// Housing represent the database housing table
type Housing struct {
	ID                    uuid.UUID `json:"id" pg:"id" swaggerignore:"true"`
	TypeID                uuid.UUID `json:"type_id" pg:"type_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	SurfaceArea           float32   `json:"surface_area" example:"15.5"`
	RentPrice             float32   `json:"rent_price" example:"60.95"`
	RentalCharges         float32   `json:"rental_charges" example:"60.9"`
	Description           string    `json:"description" example:"Appartement située aux abords d'un parc arboré"`
	Country               string    `json:"country" example:"FR"`
	State                 string    `json:"state" example:"Haut-de-France"`
	City                  string    `json:"city" example:"Lille"`
	Street                string    `json:"street" example:"78 Rue Solférino"`
	Zip                   string    `json:"zip" example:"59000"`
	IsFurnished           bool      `json:"is_furnished" pg:",use_zero" example:"false"`
	HasElectricity        bool      `json:"has_electricity" pg:",use_zero" example:"true"`
	HasGas                bool      `json:"has_gas" pg:",use_zero" example:"false"`
	IsPublished           bool      `json:"is_published" pg:",use_zero" example:"true"`
	StatusID              uuid.UUID `json:"status_id" pg:"status_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	OwnerID               uuid.UUID `json:"owner_id" pg:"owner_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	LastTenantID          uuid.UUID `json:"last_tenant_id" pg:"last_tenant_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	StripeCustomerID      string    `json:"stripe_customer_id" pg:"stripe_customer_id" example:"cus_IOwdRp9gIlOjTD"`
	StripePaymentMethodID string    `json:"stripe_payment_method_id" pg:"stripe_payment_method_id" example:"pm_1Ho8k8CMhQMU3AqAKJwPYAXj"`
	CreatedAt             time.Time `json:"created_at" swaggerignore:"true"`
	UpdatedAt             time.Time `json:"updated_at" swaggerignore:"true"`
}

// HousingBody is used as POST/PUT request body
type HousingBody struct {
	TypeID                uuid.UUID `json:"type_id,omitempty" pg:"type_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	SurfaceArea           float32   `json:"surface_area,omitempty" example:"15.5"`
	RentPrice             float32   `json:"rent_price,omitempty" example:"60.95"`
	RentalCharges         float32   `json:"rental_charges,omitempty" example:"60.9"`
	Description           string    `json:"description" example:"Appartement située aux abords d'un parc arboré"`
	Country               string    `json:"country,omitempty" example:"FR"`
	State                 string    `json:"state,omitempty" example:"Haut-de-France"`
	City                  string    `json:"city,omitempty" example:"Lille"`
	Street                string    `json:"street,omitempty" example:"78 Rue Solférino"`
	Zip                   string    `json:"zip,omitempty" example:"59000"`
	IsFurnished           bool      `json:"is_furnished,omitempty" pg:",use_zero" example:"false"`
	HasElectricity        bool      `json:"has_electricity,omitempty" pg:",use_zero" example:"true"`
	HasGas                bool      `json:"has_gas,omitempty" pg:",use_zero" example:"false"`
	IsPublished           bool      `json:"is_published,omitempty" pg:",use_zero" example:"true"`
	StatusID              uuid.UUID `json:"status_id,omitempty" pg:"status_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	OwnerID               uuid.UUID `json:"owner_id,omitempty" pg:"owner_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	LastTenantID          uuid.UUID `json:"last_tenant_id,omitempty" pg:"last_tenant_id" example:"e185deb2-91d5-4ab7-87b3-daaffac00e3d"`
	StripeCustomerID      string    `json:"stripe_customer_id,omitempty" pg:"stripe_customer_id" example:"cus_IOwdRp9gIlOjTD"`
	StripePaymentMethodID string    `json:"stripe_payment_method_id,omitempty" pg:"stripe_payment_method_id" example:"pm_1Ho8k8CMhQMU3AqAKJwPYAXj"`
}

// UpdatePublicationStatus is used inside publication status PATCH request body
type UpdatePublicationStatus struct {
	IsPublished bool `json:"is_published" example:"true"`
}
